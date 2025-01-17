/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ApiV3SpansResponseChunk - Response object with spans.
type ApiV3SpansResponseChunk struct {
	ResourceSpans []V1ResourceSpans `json:"resourceSpans,omitempty"`
}

// AssertApiV3SpansResponseChunkRequired checks if the required fields are not zero-ed
func AssertApiV3SpansResponseChunkRequired(obj ApiV3SpansResponseChunk) error {
	for _, el := range obj.ResourceSpans {
		if err := AssertV1ResourceSpansRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseApiV3SpansResponseChunkRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiV3SpansResponseChunk (e.g. [][]ApiV3SpansResponseChunk), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiV3SpansResponseChunkRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiV3SpansResponseChunk, ok := obj.(ApiV3SpansResponseChunk)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiV3SpansResponseChunkRequired(aApiV3SpansResponseChunk)
	})
}
