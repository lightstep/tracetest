/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// V1InstrumentationLibrarySpans - A collection of Spans produced by an InstrumentationLibrary.
type V1InstrumentationLibrarySpans struct {
	InstrumentationLibrary V1InstrumentationLibrary `json:"instrumentationLibrary,omitempty"`

	// A list of Spans that originate from an instrumentation library.
	Spans []V1Span `json:"spans,omitempty"`

	// This schema_url applies to all spans and span events in the \"spans\" field.
	SchemaUrl string `json:"schemaUrl,omitempty"`
}

// AssertV1InstrumentationLibrarySpansRequired checks if the required fields are not zero-ed
func AssertV1InstrumentationLibrarySpansRequired(obj V1InstrumentationLibrarySpans) error {
	if err := AssertV1InstrumentationLibraryRequired(obj.InstrumentationLibrary); err != nil {
		return err
	}
	for _, el := range obj.Spans {
		if err := AssertV1SpanRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseV1InstrumentationLibrarySpansRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of V1InstrumentationLibrarySpans (e.g. [][]V1InstrumentationLibrarySpans), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseV1InstrumentationLibrarySpansRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aV1InstrumentationLibrarySpans, ok := obj.(V1InstrumentationLibrarySpans)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertV1InstrumentationLibrarySpansRequired(aV1InstrumentationLibrarySpans)
	})
}
