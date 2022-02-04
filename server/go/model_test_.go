/*
 * Project X
 *
 * OpenAPI definition for project X endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Test struct {

	// ID
	Id int64 `gorm:"primaryKey,autoIncrement" json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	ServiceUnderTest TestServiceUnderTest `gorm:"embedded" json:"serviceUnderTest,omitempty"`

	// Definition of assertions that are going to be made
	Assertions []Assertion `gorm:"many2many:tests_assertions;" json:"assertions,omitempty"`

	Repeats int32 `json:"repeats,omitempty"`
}

// AssertTestRequired checks if the required fields are not zero-ed
func AssertTestRequired(obj Test) error {
	if err := AssertTestServiceUnderTestRequired(obj.ServiceUnderTest); err != nil {
		return err
	}
	for _, el := range obj.Assertions {
		if err := AssertAssertionRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseTestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Test (e.g. [][]Test), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTest, ok := obj.(Test)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestRequired(aTest)
	})
}
