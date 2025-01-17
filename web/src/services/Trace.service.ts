import AssertionService from './Assertion.service';
import {
  IAssertionResult,
  TAssertionResultList,
  ISpanAssertionResult2,
  ITestAssertionResult,
} from '../types/Assertion.types';
import {ITest} from '../types/Test.types';
import {ITrace} from '../types/Trace.types';

const TraceService = () => ({
  runTest(trace: ITrace, {assertions = []}: ITest) {
    const resultList = assertions?.map(assertion => AssertionService.runByTrace(trace, assertion));

    return resultList;
  },
  parseTestResultToAssertionResultList(
    assertionResult: TAssertionResultList,
    {assertions}: ITest,
    trace: ITrace
  ): IAssertionResult[] {
    const assertionResultList = assertionResult.map(({assertionId, spanAssertionResults = []}) => {
      const assertion = assertions.find(({assertionId: id}) => id === assertionId);

      return {
        assertion: assertion!,
        spanListAssertionResult: spanAssertionResults.map(({spanId, passed, observedValue, spanAssertionId}) => {
          const span = trace.spans.find(({spanId: id}) => id === spanId);
          const spanAssertion = assertion?.spanAssertions?.find(({spanAssertionId: id}) => id === spanAssertionId);

          return {
            span: span!,
            resultList: spanAssertion
              ? [{...spanAssertion, spanId, hasPassed: passed, actualValue: observedValue}]
              : [],
          };
        }),
      };
    });

    return assertionResultList;
  },
  parseAssertionResultListToTestResult(assertionResultList: IAssertionResult[] = []): ITestAssertionResult {
    const {totalFailedCount} = this.getTestResultCount(assertionResultList);

    return {
      assertionResultState: !totalFailedCount,
      assertionResult: assertionResultList.map(({assertion, spanListAssertionResult}) => ({
        assertionId: assertion.assertionId,
        spanAssertionResults: spanListAssertionResult.reduce<ISpanAssertionResult2[]>(
          (accList, {resultList}) =>
            accList.concat(
              resultList.map(({spanId, hasPassed, actualValue, spanAssertionId = ''}) => ({
                spanAssertionId,
                spanId,
                passed: hasPassed,
                observedValue: actualValue,
              }))
            ),
          []
        ),
      })),
    };
  },
  getTestResultCount(assertionResultList: IAssertionResult[]) {
    const [totalPassedCount, totalFailedCount] = assertionResultList.reduce<[number, number]>(
      ([innerTotalPassedCount, innerTotalFailedCount], {spanListAssertionResult}) => {
        const [passed, failed] = spanListAssertionResult.reduce<[number, number]>(
          ([passedResultCount, failedResultCount], {resultList}) => {
            const passedCount = resultList.filter(({hasPassed}) => hasPassed).length;
            const failedCount = resultList.filter(({hasPassed}) => !hasPassed).length;

            return [passedResultCount + passedCount, failedResultCount + failedCount];
          },
          [0, 0]
        );

        return [innerTotalPassedCount + passed, innerTotalFailedCount + failed];
      },
      [0, 0]
    );

    return {totalFailedCount, totalPassedCount};
  },
});

export default TraceService();
