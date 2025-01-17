/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type HttpAuth struct {
	Type string `json:"type,omitempty"`

	ApiKey HttpAuthApiKey `json:"apiKey,omitempty"`

	Basic HttpAuthBasic `json:"basic,omitempty"`

	Bearer HttpAuthBearer `json:"bearer,omitempty"`
}

// AssertHttpAuthRequired checks if the required fields are not zero-ed
func AssertHttpAuthRequired(obj HttpAuth) error {
	if err := AssertHttpAuthApiKeyRequired(obj.ApiKey); err != nil {
		return err
	}
	if err := AssertHttpAuthBasicRequired(obj.Basic); err != nil {
		return err
	}
	if err := AssertHttpAuthBearerRequired(obj.Bearer); err != nil {
		return err
	}
	return nil
}

// AssertRecurseHttpAuthRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of HttpAuth (e.g. [][]HttpAuth), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseHttpAuthRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aHttpAuth, ok := obj.(HttpAuth)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertHttpAuthRequired(aHttpAuth)
	})
}
