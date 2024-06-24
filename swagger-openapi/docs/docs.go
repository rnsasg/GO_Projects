// Package classification awesome.
//
// Documentation of our awesome API.
//
//	 Schemes: http
//	 BasePath: /
//	 Version: 1.0.0
//	 Host: some-url.com
//
//	 Consumes:
//	 - application/json
//
//	 Produces:
//	 - application/json
//
//	 Security:
//	 - basic
//
//	SecurityDefinitions:
//	basic:
//	  type: basic
//
// swagger:meta
package docs

import "github.com/pdrum/swagger-automation/api"

// swagger:route POST /foobar foobar-tag idOfFoobarEndpoint
// Foobar does some amazing stuff.
// responses:
//   200: foobarResponse

// This text will appear as description of your response body.
// swagger:response foobarResponse
type foobarResponseWrapper struct {
	// in:body
	Body api.FooBarResponse
}

// swagger:parameters idOfFoobarEndpoint
type foobarParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body api.FooBarRequest
}
