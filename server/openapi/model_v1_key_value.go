/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// V1KeyValue - KeyValue is a key-value pair that is used to store Span attributes, Link attributes, etc.
type V1KeyValue struct {
	Key string `json:"key,omitempty"`

	Value V1AnyValue `json:"value,omitempty"`
}

// AssertV1KeyValueRequired checks if the required fields are not zero-ed
func AssertV1KeyValueRequired(obj V1KeyValue) error {
	if err := AssertV1AnyValueRequired(obj.Value); err != nil {
		return err
	}
	return nil
}

// AssertRecurseV1KeyValueRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of V1KeyValue (e.g. [][]V1KeyValue), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseV1KeyValueRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aV1KeyValue, ok := obj.(V1KeyValue)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertV1KeyValueRequired(aV1KeyValue)
	})
}
