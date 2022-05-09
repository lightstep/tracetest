/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Assertion struct {
	Id string `json:"id,omitempty"`

	Attr string `json:"attr,omitempty"`

	Comparator string `json:"comparator,omitempty"`

	Expected string `json:"expected,omitempty"`
}

// AssertAssertionRequired checks if the required fields are not zero-ed
func AssertAssertionRequired(obj Assertion) error {
	return nil
}

// AssertRecurseAssertionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Assertion (e.g. [][]Assertion), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAssertionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAssertion, ok := obj.(Assertion)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAssertionRequired(aAssertion)
	})
}
