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
	"context"
	"net/http"
	"errors"
)

// PresentProofApiService is a service that implents the logic for the PresentProofApiServicer
// This service should implement the business logic for every endpoint for the PresentProofApi API. 
// Include any external packages or services that will be required by this service.
type PresentProofApiService struct {
}

// NewPresentProofApiService creates a default api service
func NewPresentProofApiService() PresentProofApiServicer {
	return &PresentProofApiService{}
}

// PresentProofGetByThreadId - Get presentation exchange record by thread id
func (s *PresentProofApiService) PresentProofGetByThreadId(ctx context.Context, presentationExchangeThreadId string) (ImplResponse, error) {
	// TODO - update PresentProofGetByThreadId with the required logic for this service method.
	// Add api_present_proof_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofOperationResponse{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofOperationResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofGetByThreadId method not implemented")
}

// PresentProofSendPresentation - Send presentation
func (s *PresentProofApiService) PresentProofSendPresentation(ctx context.Context, inlineObject9 InlineObject9) (ImplResponse, error) {
	// TODO - update PresentProofSendPresentation with the required logic for this service method.
	// Add api_present_proof_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofOperationResponse{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofOperationResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofSendPresentation method not implemented")
}

// PresentProofSendProposal - Send presentation proposal
func (s *PresentProofApiService) PresentProofSendProposal(ctx context.Context, inlineObject7 InlineObject7) (ImplResponse, error) {
	// TODO - update PresentProofSendProposal with the required logic for this service method.
	// Add api_present_proof_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofOperationResponse{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofOperationResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofSendProposal method not implemented")
}

// PresentProofSendRequest - Send presentation request
func (s *PresentProofApiService) PresentProofSendRequest(ctx context.Context, inlineObject8 InlineObject8) (ImplResponse, error) {
	// TODO - update PresentProofSendRequest with the required logic for this service method.
	// Add api_present_proof_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofOperationResponse{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofOperationResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofSendRequest method not implemented")
}

// PresentProofVerifyPresentation - Verify presentation
func (s *PresentProofApiService) PresentProofVerifyPresentation(ctx context.Context, inlineObject10 InlineObject10) (ImplResponse, error) {
	// TODO - update PresentProofVerifyPresentation with the required logic for this service method.
	// Add api_present_proof_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PresentProofOperationResponse{}) or use other options such as http.Ok ...
	//return Response(200, PresentProofOperationResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PresentProofVerifyPresentation method not implemented")
}

