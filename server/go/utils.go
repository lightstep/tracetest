package openapi

import (
	"fmt"
	"net/http"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	v1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

func EncodeJSONPBResponse(i interface{}, status *int, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	trace := i.(proto.Message)
	m := jsonpb.Marshaler{}
	return m.Marshal(w, trace)
}

type Trace struct {
	Parent *v1.Span `json:"parent"`
	Child  []*Trace `json:"child"`
}

func TransformTrace(tr *v1.TracesData, traceID, parentSpanID string) *Trace {
	spanIdChild := make(map[string][]*v1.Span)
	spans := make(map[string]*v1.Span)
	for _, rs := range tr.ResourceSpans {
		for _, ils := range rs.InstrumentationLibrarySpans {
			for _, sp := range ils.Spans {
				spans[string(sp.SpanId)] = sp
				if sp.ParentSpanId == nil {
					continue
				}
				sps := spanIdChild[string(sp.ParentSpanId)]
				spanIdChild[string(sp.ParentSpanId)] = append(sps, sp)
			}
		}
	}

	// Fix parent id
	for _, sp := range spans {
		if sp.ParentSpanId == nil {
			continue
		}
		_, ok := spans[string(sp.ParentSpanId)]
		if !ok {
			fmt.Printf("setting parent span for %v\n", sp)
			sp.ParentSpanId = []byte(parentSpanID)
			// Fix up the child map
			sps := spanIdChild[string(sp.ParentSpanId)]
			spanIdChild[string(sp.ParentSpanId)] = append(sps, sp)
		}
		fmt.Printf("span : %v\n", sp)
	}

	res := &Trace{
		Parent: &v1.Span{
			TraceId:      []byte(traceID),
			SpanId:       []byte(parentSpanID),
			ParentSpanId: nil,
			Name:         "tracetest",
			Kind:         v1.Span_SPAN_KIND_CLIENT,
		},
		Child: nil,
	}

	fmt.Printf("initial parent: %v\n", res.Parent)
	return MakeTrace(spanIdChild, res)
}

func MakeTrace(spans map[string][]*v1.Span, trace *Trace) *Trace {
	childSpans, ok := spans[string(trace.Parent.SpanId)]
	if !ok {
		return trace
	}

	for _, childSpan := range childSpans {
		childTrace := MakeTrace(spans, &Trace{Parent: childSpan, Child: nil})
		trace.Child = append(trace.Child, childTrace)
	}
	return trace
}
