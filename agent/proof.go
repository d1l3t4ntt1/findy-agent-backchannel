package agent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
)

type ProofStatus = agency.ProtocolStatus_PresentProofStatus
type ProofAttribute = agency.Protocol_Proof_Attribute

type Proof struct {
	ID        string `json:"referent"`
	CredDefID string `json:"proof_def_id"`
	SchemaID  string `json:"schema_id"`
}

type ProofStore struct {
	agent       *AgencyClient
	readyProofs map[string]*ProofStatus
	sentProofs  map[string]string
	requests    map[string]string
	sync.RWMutex
}

func InitProofs(a *AgencyClient) *ProofStore {
	return &ProofStore{
		agent:       a,
		readyProofs: make(map[string]*ProofStatus),
		sentProofs:  make(map[string]string),
		requests:    make(map[string]string),
	}
}

func (s *ProofStore) HandleProofNotification(notification *agency.Notification) (err error) {
	defer err2.Return(&err)

	// Cred issued
	if notification.GetTypeID() == agency.Notification_STATUS_UPDATE {
		if notification.GetProtocolType() == agency.Protocol_PRESENT_PROOF {
			protocolID := &agency.ProtocolID{
				ID:     notification.ProtocolID,
				TypeID: notification.ProtocolType,
			}

			var status *agency.ProtocolStatus
			status, err = s.agent.ProtocolClient.Status(context.TODO(), protocolID)
			err2.Check(err)

			if status.State.State == agency.ProtocolState_OK {
				proof := status.GetPresentProof()
				log.Printf("New proof %v\n", proof)
				_, err = s.AddProof(protocolID.ID, proof)
				err2.Check(err)
			}
		}

		// Cred request received
	} else if notification.GetProtocolType() == agency.Protocol_PRESENT_PROOF &&
		notification.GetTypeID() == agency.Notification_PROTOCOL_PAUSED {
		_, err = s.AddProofRequest(notification.ProtocolID)
		err2.Check(err)
	}
	return nil
}

func (s *ProofStore) SendProofPresentation(id string) (threadID string, err error) {
	defer err2.Return(&err)

	threadID, err = s.GetProofRequest(id)
	err2.Check(err)

	state := &agency.ProtocolState{
		ProtocolID: &agency.ProtocolID{
			TypeID: agency.Protocol_PRESENT_PROOF,
			Role:   agency.Protocol_RESUMER,
			ID:     id,
		},
		State: agency.ProtocolState_ACK,
	}

	_, err = s.agent.ProtocolClient.Resume(
		context.TODO(),
		state,
	)
	err2.Check(err)

	_, err = s.AddSentProof(id)
	err2.Check(err)

	return threadID, nil
}

func (s *CredentialStore) ProposeProof(connectionID string, attributes []*ProofAttribute) (threadID string, err error) {
	defer err2.Return(&err)

	log.Printf("Propose proof, conn id: %s, attrs: %v", connectionID, attributes)

	protocol := &agency.Protocol{
		ConnectionID: connectionID,
		TypeID:       agency.Protocol_PRESENT_PROOF,
		Role:         agency.Protocol_ADDRESSEE,
		StartMsg: &agency.Protocol_PresentProof{
			PresentProof: &agency.Protocol_PresentProofMsg{
				AttrFmt: &agency.Protocol_PresentProofMsg_Attributes{
					Attributes: &agency.Protocol_Proof{
						Attributes: attributes,
					},
				},
			},
		},
	}
	res, err := s.agent.Conn.DoStart(context.TODO(), protocol)
	err2.Check(err)

	return res.ID, nil
}

func (s *ProofStore) AddProof(id string, c *ProofStatus) (*Proof, error) {
	s.Lock()
	defer s.Unlock()
	if c != nil {
		s.readyProofs[id] = c
		res := &Proof{
			ID: id,
		}
		return res, nil
	}
	return nil, fmt.Errorf("cannot add non-existent proof with id %s", id)
}

func (s *ProofStore) GetProof(id string) (*Proof, error) {
	s.RLock()
	defer s.RUnlock()
	if _, ok := s.readyProofs[id]; ok {
		res := &Proof{
			ID: id,
		}
		return res, nil
	}
	return nil, fmt.Errorf("proof by the id %s not found", id)
}

func (s *ProofStore) AddSentProof(id string) (string, error) {
	s.Lock()
	defer s.Unlock()
	if id != "" {
		s.sentProofs[id] = id
		return id, nil
	}
	return "", fmt.Errorf("cannot add sent proof with id %s", id)
}

func (s *ProofStore) GetSentProof(id string) (string, error) {
	s.RLock()
	defer s.RUnlock()
	if _, ok := s.sentProofs[id]; ok {
		return id, nil
	}
	return "", fmt.Errorf("sent proof by the id %s not found", id)
}

func (s *ProofStore) AddProofRequest(id string) (string, error) {
	s.Lock()
	defer s.Unlock()
	if id != "" {
		s.requests[id] = id
		return id, nil
	}
	return "", errors.New("cannot add non-existent proof request")
}

func (s *ProofStore) GetProofRequest(id string) (string, error) {
	s.RLock()
	defer s.RUnlock()
	if request, ok := s.requests[id]; ok {
		return request, nil
	}
	return "", fmt.Errorf("proof request by the id %s not found", id)
}