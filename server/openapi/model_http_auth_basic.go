/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type HttpAuthBasic struct {
	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`
}

// AssertHttpAuthBasicRequired checks if the required fields are not zero-ed
func AssertHttpAuthBasicRequired(obj HttpAuthBasic) error {
	return nil
}

// AssertRecurseHttpAuthBasicRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of HttpAuthBasic (e.g. [][]HttpAuthBasic), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseHttpAuthBasicRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aHttpAuthBasic, ok := obj.(HttpAuthBasic)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertHttpAuthBasicRequired(aHttpAuthBasic)
	})
}
