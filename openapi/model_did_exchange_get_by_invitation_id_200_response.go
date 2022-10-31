/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DidExchangeGetByInvitationId200Response struct {

	ConnectionId string `json:"connection_id,omitempty"`
}

// AssertDidExchangeGetByInvitationId200ResponseRequired checks if the required fields are not zero-ed
func AssertDidExchangeGetByInvitationId200ResponseRequired(obj DidExchangeGetByInvitationId200Response) error {
	return nil
}

// AssertRecurseDidExchangeGetByInvitationId200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DidExchangeGetByInvitationId200Response (e.g. [][]DidExchangeGetByInvitationId200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDidExchangeGetByInvitationId200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDidExchangeGetByInvitationId200Response, ok := obj.(DidExchangeGetByInvitationId200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDidExchangeGetByInvitationId200ResponseRequired(aDidExchangeGetByInvitationId200Response)
	})
}