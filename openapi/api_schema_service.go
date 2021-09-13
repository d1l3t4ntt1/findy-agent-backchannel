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
	"encoding/json"
	"net/http"

	"github.com/findy-network/findy-agent-backchannel/agent"
)

// SchemaApiService is a service that implents the logic for the SchemaApiServicer
// This service should implement the business logic for every endpoint for the SchemaApi API.
// Include any external packages or services that will be required by this service.
type SchemaApiService struct {
	a *agent.Agent
}

// NewSchemaApiService creates a default api service
func NewSchemaApiService(a *agent.Agent) SchemaApiServicer {
	return &SchemaApiService{
		a: a,
	}
}

// SchemaCreate - Create a new schema
func (s *SchemaApiService) SchemaCreate(ctx context.Context, inlineObject4 InlineObject4) (ImplResponse, error) {
	res, err := s.a.CreateSchema(inlineObject4.Data.SchemaName, inlineObject4.Data.SchemaVersion, inlineObject4.Data.Attributes)
	if err == nil {
		schema := make(map[string]interface{})
		schema["id"] = res
		schema["ver"] = inlineObject4.Data.SchemaVersion
		schema["name"] = inlineObject4.Data.SchemaName
		schema["attrNames"] = inlineObject4.Data.Attributes
		schema["seqNo"] = "" // ?
		return Response(200, InlineResponse2004{SchemaId: res, Schema: schema}), nil
	}

	return Response(http.StatusInternalServerError, nil), err
}

// SchemaGetById - Get schema by id
func (s *SchemaApiService) SchemaGetById(ctx context.Context, schemaId string) (ImplResponse, error) {
	res, err := s.a.GetSchema(schemaId)
	if err == nil {
		var resMap map[string]interface{}
		if err = json.Unmarshal([]byte(res), &resMap); err == nil {
			return Response(200, resMap), nil
		}
	}
	return Response(http.StatusInternalServerError, nil), err
}
