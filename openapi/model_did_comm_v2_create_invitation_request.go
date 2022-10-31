/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DidCommV2CreateInvitationRequest struct {

	Data DidCommV2CreateInvitationRequestData `json:"data,omitempty"`
}

// AssertDidCommV2CreateInvitationRequestRequired checks if the required fields are not zero-ed
func AssertDidCommV2CreateInvitationRequestRequired(obj DidCommV2CreateInvitationRequest) error {
	if err := AssertDidCommV2CreateInvitationRequestDataRequired(obj.Data); err != nil {
		return err
	}
	return nil
}

// AssertRecurseDidCommV2CreateInvitationRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DidCommV2CreateInvitationRequest (e.g. [][]DidCommV2CreateInvitationRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDidCommV2CreateInvitationRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDidCommV2CreateInvitationRequest, ok := obj.(DidCommV2CreateInvitationRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDidCommV2CreateInvitationRequestRequired(aDidCommV2CreateInvitationRequest)
	})
}