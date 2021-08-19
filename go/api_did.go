/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"strings"

)

// A DIDApiController binds http requests to an api service and writes the service results to the http response
type DIDApiController struct {
	service DIDApiServicer
}

// NewDIDApiController creates a default api controller
func NewDIDApiController(s DIDApiServicer) Router {
	return &DIDApiController{service: s}
}

// Routes returns all of the api route for the DIDApiController
func (c *DIDApiController) Routes() Routes {
	return Routes{ 
		{
			"DIDGetPublic",
			strings.ToUpper("Get"),
			"/agent/command/did",
			c.DIDGetPublic,
		},
	}
}

// DIDGetPublic - Get public DID
func (c *DIDApiController) DIDGetPublic(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.DIDGetPublic(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}