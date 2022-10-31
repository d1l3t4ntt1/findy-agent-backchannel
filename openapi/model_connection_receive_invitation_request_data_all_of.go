/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ConnectionReceiveInvitationRequestDataAllOf struct {

	MediatorConnectionId string `json:"mediator_connection_id,omitempty"`
}

// AssertConnectionReceiveInvitationRequestDataAllOfRequired checks if the required fields are not zero-ed
func AssertConnectionReceiveInvitationRequestDataAllOfRequired(obj ConnectionReceiveInvitationRequestDataAllOf) error {
	return nil
}

// AssertRecurseConnectionReceiveInvitationRequestDataAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ConnectionReceiveInvitationRequestDataAllOf (e.g. [][]ConnectionReceiveInvitationRequestDataAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseConnectionReceiveInvitationRequestDataAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aConnectionReceiveInvitationRequestDataAllOf, ok := obj.(ConnectionReceiveInvitationRequestDataAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertConnectionReceiveInvitationRequestDataAllOfRequired(aConnectionReceiveInvitationRequestDataAllOf)
	})
}