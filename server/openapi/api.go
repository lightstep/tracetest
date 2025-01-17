/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)

// ApiApiRouter defines the required methods for binding the api requests to a responses for the ApiApi
// The ApiApiRouter implementation should parse necessary information from the http request,
// pass the data to a ApiApiServicer to perform the required actions, then write the service results to the http response.
type ApiApiRouter interface {
	CreateAssertion(http.ResponseWriter, *http.Request)
	CreateTest(http.ResponseWriter, *http.Request)
	DeleteAssertion(http.ResponseWriter, *http.Request)
	DeleteTest(http.ResponseWriter, *http.Request)
	GetAssertions(http.ResponseWriter, *http.Request)
	GetTest(http.ResponseWriter, *http.Request)
	GetTestResult(http.ResponseWriter, *http.Request)
	GetTestResults(http.ResponseWriter, *http.Request)
	GetTests(http.ResponseWriter, *http.Request)
	RunTest(http.ResponseWriter, *http.Request)
	UpdateAssertion(http.ResponseWriter, *http.Request)
	UpdateTest(http.ResponseWriter, *http.Request)
	UpdateTestResult(http.ResponseWriter, *http.Request)
}

// ApiApiServicer defines the api actions for the ApiApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ApiApiServicer interface {
	CreateAssertion(context.Context, string, Assertion) (ImplResponse, error)
	CreateTest(context.Context, Test) (ImplResponse, error)
	DeleteAssertion(context.Context, string, string) (ImplResponse, error)
	DeleteTest(context.Context, string) (ImplResponse, error)
	GetAssertions(context.Context, string) (ImplResponse, error)
	GetTest(context.Context, string) (ImplResponse, error)
	GetTestResult(context.Context, string, string) (ImplResponse, error)
	GetTestResults(context.Context, string) (ImplResponse, error)
	GetTests(context.Context) (ImplResponse, error)
	RunTest(context.Context, string) (ImplResponse, error)
	UpdateAssertion(context.Context, string, string, Assertion) (ImplResponse, error)
	UpdateTest(context.Context, string, Test) (ImplResponse, error)
	UpdateTestResult(context.Context, string, string, TestAssertionResult) (ImplResponse, error)
}
