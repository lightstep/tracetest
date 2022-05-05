package assertions

import (
	"fmt"
	"math"

	"github.com/kubeshop/tracetest/openapi"
	"github.com/kubeshop/tracetest/traces"
	"go.opentelemetry.io/otel/trace"
)

func convertOTelTraceIntoTraceTree(trace openapi.ApiV3SpansResponseChunk) (traces.Trace, error) {
	flattenSpans := make([]openapi.V1Span, 0)
	for _, resource := range trace.ResourceSpans {
		for _, librarySpans := range resource.InstrumentationLibrarySpans {
			for _, span := range librarySpans.Spans {
				flattenSpans = append(flattenSpans, span)
			}
		}
	}

	spansMap := make(map[string]*traces.Span, 0)
	for _, span := range flattenSpans {
		newSpan, err := convertOtelSpanIntoSpan(span)
		if err != nil {
			return traces.Trace{}, err
		}
		spansMap[span.SpanId] = newSpan
	}

	return createTrace(flattenSpans, spansMap), nil
}

func convertOtelSpanIntoSpan(span openapi.V1Span) (*traces.Span, error) {
	attributes := make(traces.Attributes, 0)
	for _, attribute := range span.Attributes {
		attributes[attribute.Key] = getAttributeValue(attribute.Value)
	}

	spanID, err := createSpanID(span.SpanId)
	if err != nil {
		return nil, err
	}

	return &traces.Span{
		ID:         spanID,
		Name:       span.Name,
		Parent:     nil,
		Children:   make([]*traces.Span, 0),
		Attributes: attributes,
	}, nil
}

func getAttributeValue(value openapi.V1AnyValue) string {
	if value.StringValue != "" {
		return value.StringValue
	}

	if value.IntValue != "" {
		return value.IntValue
	}

	if value.DoubleValue != 0.0 {
		isFloatingPoint := math.Abs(value.DoubleValue-math.Abs(value.DoubleValue)) > 0.0
		if isFloatingPoint {
			return fmt.Sprintf("%f", value.DoubleValue)
		}

		return fmt.Sprintf("%.0f", value.DoubleValue)
	}

	return fmt.Sprintf("%t", value.BoolValue)
}

func createSpanID(id string) (trace.SpanID, error) {
	spanId, err := trace.SpanIDFromHex(id)
	if err != nil {
		return trace.SpanID{}, fmt.Errorf("could not convert spanID")
	}

	return spanId, nil
}

func createTrace(spans []openapi.V1Span, spansMap map[string]*traces.Span) traces.Trace {
	var rootSpanID string = ""
	for _, span := range spans {
		if span.ParentSpanId == "" {
			rootSpanID = span.SpanId
		} else {
			parent := spansMap[span.ParentSpanId]
			thisSpan := spansMap[span.SpanId]

			thisSpan.Parent = parent
			parent.Children = append(parent.Children, thisSpan)
		}
	}

	rootSpan := spansMap[rootSpanID]

	return traces.Trace{
		RootSpan: *rootSpan,
	}
}
