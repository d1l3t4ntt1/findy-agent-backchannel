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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DIDExchangeApiController binds http requests to an api service and writes the service results to the http response
type DIDExchangeApiController struct {
	service      DIDExchangeApiServicer
	errorHandler ErrorHandler
}

// DIDExchangeApiOption for how the controller is set up.
type DIDExchangeApiOption func(*DIDExchangeApiController)

// WithDIDExchangeApiErrorHandler inject ErrorHandler into controller
func WithDIDExchangeApiErrorHandler(h ErrorHandler) DIDExchangeApiOption {
	return func(c *DIDExchangeApiController) {
		c.errorHandler = h
	}
}

// NewDIDExchangeApiController creates a default api controller
func NewDIDExchangeApiController(s DIDExchangeApiServicer, opts ...DIDExchangeApiOption) Router {
	controller := &DIDExchangeApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DIDExchangeApiController
func (c *DIDExchangeApiController) Routes() Routes {
	return Routes{
		{
			"DidExchangeCreateRequestResolvableDid",
			strings.ToUpper("Post"),
			"/agent/command/did-exchange/create-request-resolvable-did",
			c.DidExchangeCreateRequestResolvableDid,
		},
		{
			"DidExchangeGetById",
			strings.ToUpper("Get"),
			"/agent/command/did-exchange/{connectionId}",
			c.DidExchangeGetById,
		},
		{
			"DidExchangeGetByInvitationId",
			strings.ToUpper("Get"),
			"/agent/response/did-exchange/{invitationId}",
			c.DidExchangeGetByInvitationId,
		},
		{
			"DidExchangeReceiveRequestResolvableDid",
			strings.ToUpper("Post"),
			"/agent/command/did-exchange/receive-request-resolvable-did",
			c.DidExchangeReceiveRequestResolvableDid,
		},
		{
			"DidExchangeSendRequest",
			strings.ToUpper("Post"),
			"/agent/command/did-exchange/{send-request:send-request\\/?}",
			c.DidExchangeSendRequest,
		},
		{
			"DidExchangeSendResponse",
			strings.ToUpper("Post"),
			"/agent/command/did-exchange/send-response",
			c.DidExchangeSendResponse,
		},
	}
}

// DidExchangeCreateRequestResolvableDid - Send a did exchange request to the didcomm service registered in the public did
func (c *DIDExchangeApiController) DidExchangeCreateRequestResolvableDid(w http.ResponseWriter, r *http.Request) {
	didExchangeCreateRequestResolvableDidRequestParam := DidExchangeCreateRequestResolvableDidRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&didExchangeCreateRequestResolvableDidRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDidExchangeCreateRequestResolvableDidRequestRequired(didExchangeCreateRequestResolvableDidRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DidExchangeCreateRequestResolvableDid(r.Context(), didExchangeCreateRequestResolvableDidRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DidExchangeGetById - Get did exchange connection by id
func (c *DIDExchangeApiController) DidExchangeGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	connectionIdParam := params["connectionId"]

	result, err := c.service.DidExchangeGetById(r.Context(), connectionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DidExchangeGetByInvitationId - Get did exchange connection by invitation id. Can be used to determine the connection id based of an invitation id.
func (c *DIDExchangeApiController) DidExchangeGetByInvitationId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	invitationIdParam := params["invitationId"]

	result, err := c.service.DidExchangeGetByInvitationId(r.Context(), invitationIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DidExchangeReceiveRequestResolvableDid - The documentation of this endpoint need to be looked at.
func (c *DIDExchangeApiController) DidExchangeReceiveRequestResolvableDid(w http.ResponseWriter, r *http.Request) {
	didExchangeReceiveRequestResolvableDidRequestParam := DidExchangeReceiveRequestResolvableDidRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&didExchangeReceiveRequestResolvableDidRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDidExchangeReceiveRequestResolvableDidRequestRequired(didExchangeReceiveRequestResolvableDidRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DidExchangeReceiveRequestResolvableDid(r.Context(), didExchangeReceiveRequestResolvableDidRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DidExchangeSendRequest - Send a did exchange request to the connection with specified id.
func (c *DIDExchangeApiController) DidExchangeSendRequest(w http.ResponseWriter, r *http.Request) {
	connectionAcceptInvitationRequestParam := ConnectionAcceptInvitationRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&connectionAcceptInvitationRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertConnectionAcceptInvitationRequestRequired(connectionAcceptInvitationRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DidExchangeSendRequest(r.Context(), connectionAcceptInvitationRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DidExchangeSendResponse - Send a did exchange response to the connection with specified id.
func (c *DIDExchangeApiController) DidExchangeSendResponse(w http.ResponseWriter, r *http.Request) {
	connectionAcceptInvitationRequestParam := ConnectionAcceptInvitationRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&connectionAcceptInvitationRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertConnectionAcceptInvitationRequestRequired(connectionAcceptInvitationRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DidExchangeSendResponse(r.Context(), connectionAcceptInvitationRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
