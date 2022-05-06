package assertions

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/kubeshop/tracetest/executor"
	"github.com/kubeshop/tracetest/openapi"
	"github.com/kubeshop/tracetest/testdb"
	"github.com/kubeshop/tracetest/traces"
)

type RunAssertionsMessage struct {
	Test   openapi.Test
	Result openapi.TestRunResult
}

type AssertionFinishCallback func(openapi.Test, openapi.TestRunResult)

type Executor struct {
	resultDB     testdb.ResultRepository
	inputChannel chan RunAssertionsMessage
	exitChannel  chan bool
}

var _ executor.WorkerPool = &Executor{}

func NewExecutor(resultRepository testdb.ResultRepository) *Executor {
	return &Executor{
		resultDB:     resultRepository,
		inputChannel: make(chan RunAssertionsMessage, 1),
	}
}

func (e *Executor) Start(workers int) {
	e.exitChannel = make(chan bool, workers)

	for i := 0; i < workers; i++ {
		ctx := context.Background()
		go e.startWorker(ctx)
	}
}

func (e *Executor) Stop() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			e.exitChannel <- true
			return
		}
	}
}

func (e *Executor) startWorker(ctx context.Context) {
	for {
		select {
		case <-e.exitChannel:
			fmt.Println("Exiting assertion executor worker")
			return
		case request := <-e.inputChannel:
			response, err := e.executeAssertions(request)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = e.resultDB.UpdateResult(ctx, &response.Result)
			if err != nil {
				fmt.Println(fmt.Errorf("could not save result on database: %w", err).Error())
			}
		}
	}
}

func (e *Executor) executeAssertions(request RunAssertionsMessage) (*RunAssertionsMessage, error) {
	trace, err := traces.FromOtel(request.Result.Trace)
	if err != nil {
		return nil, err
	}

	testDefinition := convertAssertionsIntoTestDefinition(request.Test.Assertions)

	result := Assert(trace, testDefinition)

	response := e.setResults(request, result)

	return response, nil
}

func (e *Executor) setResults(request RunAssertionsMessage, testResult TestResult) *RunAssertionsMessage {
	response := request
	response.Result.State = executor.TestRunStateFinished
	response.Result.CompletedAt = time.Now()
	assertionResultArray := make([]openapi.AssertionResult, 0)
	allTestsPassed := true

	for _, assertionResult := range testResult {
		spanAssertions := make([]openapi.SpanAssertionResult, 0)
		for _, spanAssertionResult := range assertionResult.AssertionSpanResults {
			spanID := hex.EncodeToString(spanAssertionResult.Span.ID[:])
			testPassed := spanAssertionResult.CompareErr == nil
			if !testPassed {
				allTestsPassed = false
			}

			spanAssertions = append(spanAssertions, openapi.SpanAssertionResult{
				Passed:        testPassed,
				SpanId:        spanID,
				ObservedValue: spanAssertionResult.ActualValue,
			})
		}

		result := openapi.AssertionResult{
			AssertionId:          assertionResult.Assertion.ID,
			SpanAssertionResults: spanAssertions,
		}

		assertionResultArray = append(assertionResultArray, result)
	}

	response.Result.AssertionResult = assertionResultArray
	response.Result.AssertionResultState = allTestsPassed

	return &response
}

func (e *Executor) RunAssertions(test openapi.Test, result openapi.TestRunResult) {
	message := RunAssertionsMessage{
		test,
		result,
	}

	e.inputChannel <- message
}
