/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type OutOfBandCreateInvitationRequest struct {

	Data OutOfBandCreateInvitationRequestData `json:"data"`
}

// AssertOutOfBandCreateInvitationRequestRequired checks if the required fields are not zero-ed
func AssertOutOfBandCreateInvitationRequestRequired(obj OutOfBandCreateInvitationRequest) error {
	elements := map[string]interface{}{
		"data": obj.Data,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertOutOfBandCreateInvitationRequestDataRequired(obj.Data); err != nil {
		return err
	}
	return nil
}

// AssertRecurseOutOfBandCreateInvitationRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of OutOfBandCreateInvitationRequest (e.g. [][]OutOfBandCreateInvitationRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseOutOfBandCreateInvitationRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aOutOfBandCreateInvitationRequest, ok := obj.(OutOfBandCreateInvitationRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertOutOfBandCreateInvitationRequestRequired(aOutOfBandCreateInvitationRequest)
	})
}