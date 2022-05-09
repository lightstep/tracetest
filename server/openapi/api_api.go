/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ApiApiController binds http requests to an api service and writes the service results to the http response
type ApiApiController struct {
	service      ApiApiServicer
	errorHandler ErrorHandler
}

// ApiApiOption for how the controller is set up.
type ApiApiOption func(*ApiApiController)

// WithApiApiErrorHandler inject ErrorHandler into controller
func WithApiApiErrorHandler(h ErrorHandler) ApiApiOption {
	return func(c *ApiApiController) {
		c.errorHandler = h
	}
}

// NewApiApiController creates a default api controller
func NewApiApiController(s ApiApiServicer, opts ...ApiApiOption) Router {
	controller := &ApiApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the ApiApiController
func (c *ApiApiController) Routes() Routes {
	return Routes{
		{
			"CreateAssertion",
			strings.ToUpper("Post"),
			"/api/tests/{testId}/assertions",
			c.CreateAssertion,
		},
		{
			"CreateTest",
			strings.ToUpper("Post"),
			"/api/tests",
			c.CreateTest,
		},
		{
			"DeleteAssertion",
			strings.ToUpper("Delete"),
			"/api/tests/{testId}/assertions/{assertionId}",
			c.DeleteAssertion,
		},
		{
			"DeleteTest",
			strings.ToUpper("Delete"),
			"/api/tests/{testId}",
			c.DeleteTest,
		},
		{
			"GetAssertions",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/assertions",
			c.GetAssertions,
		},
		{
			"GetTest",
			strings.ToUpper("Get"),
			"/api/tests/{testId}",
			c.GetTest,
		},
		{
			"GetTestResult",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/results/{resultId}",
			c.GetTestResult,
		},
		{
			"GetTestResults",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/results",
			c.GetTestResults,
		},
		{
			"GetTests",
			strings.ToUpper("Get"),
			"/api/tests",
			c.GetTests,
		},
		{
			"RunTest",
			strings.ToUpper("Post"),
			"/api/tests/{testId}/run",
			c.RunTest,
		},
		{
			"UpdateAssertion",
			strings.ToUpper("Put"),
			"/api/tests/{testId}/assertions/{assertionId}",
			c.UpdateAssertion,
		},
		{
			"UpdateTest",
			strings.ToUpper("Put"),
			"/api/tests/{testId}",
			c.UpdateTest,
		},
		{
			"UpdateTestResult",
			strings.ToUpper("Put"),
			"/api/tests/{testId}/results/{resultId}",
			c.UpdateTestResult,
		},
	}
}

// CreateAssertion - Create an assertion for a test
func (c *ApiApiController) CreateAssertion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	assertionParam := Assertion{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&assertionParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAssertionRequired(assertionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateAssertion(r.Context(), testIdParam, assertionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateTest - Create new test
func (c *ApiApiController) CreateTest(w http.ResponseWriter, r *http.Request) {
	testParam := Test{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&testParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTestRequired(testParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateTest(r.Context(), testParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteAssertion - delete an assertion
func (c *ApiApiController) DeleteAssertion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	assertionIdParam := params["assertionId"]

	result, err := c.service.DeleteAssertion(r.Context(), testIdParam, assertionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteTest - delete a test
func (c *ApiApiController) DeleteTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	result, err := c.service.DeleteTest(r.Context(), testIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetAssertions - Get assertions for a test
func (c *ApiApiController) GetAssertions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	result, err := c.service.GetAssertions(r.Context(), testIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTest - get test
func (c *ApiApiController) GetTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	result, err := c.service.GetTest(r.Context(), testIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTestResult - get test result
func (c *ApiApiController) GetTestResult(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	resultIdParam := params["resultId"]

	result, err := c.service.GetTestResult(r.Context(), testIdParam, resultIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTestResults - get the results for a test
func (c *ApiApiController) GetTestResults(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	result, err := c.service.GetTestResults(r.Context(), testIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTests - Get tests
func (c *ApiApiController) GetTests(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetTests(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RunTest - run test
func (c *ApiApiController) RunTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	result, err := c.service.RunTest(r.Context(), testIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateAssertion - update an assertion
func (c *ApiApiController) UpdateAssertion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	assertionIdParam := params["assertionId"]

	assertionParam := Assertion{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&assertionParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAssertionRequired(assertionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateAssertion(r.Context(), testIdParam, assertionIdParam, assertionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateTest - update test
func (c *ApiApiController) UpdateTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	testParam := Test{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&testParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTestRequired(testParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateTest(r.Context(), testIdParam, testParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateTestResult - update test result state
func (c *ApiApiController) UpdateTestResult(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	resultIdParam := params["resultId"]

	requestBodyParam := map[string][]Assertion{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&requestBodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.UpdateTestResult(r.Context(), testIdParam, resultIdParam, requestBodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
