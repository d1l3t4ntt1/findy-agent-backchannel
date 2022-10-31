/*
 * Aries Agent Test Harness Backchannel API
 *
 * This page documents the backchannel API the test harness uses to communicate with agents under tests.  For more information checkout the [Aries Interoperability Information](http://aries-interop.info) page. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SendIndyProofRequestDataPresentationRequest struct {

	Format string `json:"format"`

	Comment string `json:"comment,omitempty"`

	Data map[string]interface{} `json:"data"`

	ConnectionId string `json:"connection_id"`
}

// AssertSendIndyProofRequestDataPresentationRequestRequired checks if the required fields are not zero-ed
func AssertSendIndyProofRequestDataPresentationRequestRequired(obj SendIndyProofRequestDataPresentationRequest) error {
	elements := map[string]interface{}{
		"format": obj.Format,
		"data": obj.Data,
		"connection_id": obj.ConnectionId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseSendIndyProofRequestDataPresentationRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SendIndyProofRequestDataPresentationRequest (e.g. [][]SendIndyProofRequestDataPresentationRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSendIndyProofRequestDataPresentationRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSendIndyProofRequestDataPresentationRequest, ok := obj.(SendIndyProofRequestDataPresentationRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSendIndyProofRequestDataPresentationRequestRequired(aSendIndyProofRequestDataPresentationRequest)
	})
}