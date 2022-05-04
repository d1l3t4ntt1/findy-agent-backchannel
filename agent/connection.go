package agent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/findy-network/findy-common-go/agency/client/async"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
)

type ConnectionStatus = agency.ProtocolStatus_DIDExchangeStatus

type Connection struct {
	ID string `json:"id"`
}

type ConnectionStore struct {
	agent       *AgencyClient
	conns       map[string]*ConnectionStatus
	invitations map[string]string
	User        string
	mtx         sync.RWMutex
}

func InitConnections(a *AgencyClient, userName string) *ConnectionStore {
	return &ConnectionStore{
		agent:       a,
		conns:       make(map[string]*ConnectionStatus),
		invitations: make(map[string]string),
		User:        userName,
	}
}

func (s *ConnectionStore) HandleConnectionNotification(notification *agency.Notification) (err error) {
	defer err2.Return(&err)

	// Conn established
	if notification.GetTypeID() == agency.Notification_STATUS_UPDATE {
		if notification.GetProtocolType() == agency.Protocol_DIDEXCHANGE {
			protocolID := &agency.ProtocolID{
				ID:     notification.ProtocolID,
				TypeID: notification.ProtocolType,
			}
			status, err := s.agent.ProtocolClient.Status(context.TODO(), protocolID)
			err2.Check(err)

			log.Printf("Connection status %v\n", status)
			if status.State.State == agency.ProtocolState_OK {
				log.Printf("New connection %v\n", status.GetDIDExchange())
				_, err = s.AddConnection(status.GetDIDExchange().ID, status.GetDIDExchange())
				err2.Check(err)
			} else {
				log.Printf("Connection status NOK %v\n", status)
			}
		}
	}
	return nil
}

func (s *ConnectionStore) CreateInvitation() (invitation string, err error) {
	defer err2.Return(&err)

	var res *agency.Invitation
	res, err = s.agent.AgentClient.CreateInvitation(
		context.TODO(),
		&agency.InvitationBase{Label: s.User},
	)
	err2.Check(err)

	invitation = res.JSON
	log.Printf("Created invitation\n %s\n", invitation)

	return invitation, nil
}

func (s *ConnectionStore) RequestConnection(id string) (invitationID string, err error) {
	defer err2.Return(&err)

	var invitationJSON string
	invitationJSON, err = s.GetConnectionInvitation(id)
	err2.Check(err)

	invitationID = id

	pw := async.NewPairwise(s.agent.Conn, "")
	pw.Label = authnCmd.UserName

	_, err = pw.Connection(context.TODO(), invitationJSON)
	err2.Check(err)

	return invitationID, nil
}

func (s *ConnectionStore) TrustPing(connectionID string) (res string, err error) {
	defer err2.Return(&err)

	_, err = s.GetConnection(connectionID)

	var totalWaitTime time.Duration
	// TODO: use waitgroups or such
	for err != nil && totalWaitTime < MaxWaitTime {
		totalWaitTime += WaitTime
		log.Println("Connection not found, waiting for to receive the connection", connectionID)
		time.Sleep(WaitTime)
		_, err = s.GetConnection(connectionID)
	}
	err2.Check(err)

	pw := async.NewPairwise(s.agent.Conn, connectionID)
	_, err = pw.Ping(context.TODO())
	err2.Check(err)

	return connectionID, nil
}

func (s *ConnectionStore) AddConnection(id string, c *ConnectionStatus) (*Connection, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if c != nil {
		s.conns[id] = c
		res := &Connection{
			ID: id,
		}
		return res, nil
	}
	return nil, fmt.Errorf("cannot add non-existent connection with id %s", id)
}

func (s *ConnectionStore) GetConnection(id string) (*Connection, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if _, ok := s.conns[id]; ok {
		res := &Connection{
			ID: id,
		}
		return res, nil
	}
	return nil, fmt.Errorf("connection by the id %s not found", id)
}

func (s *ConnectionStore) AddConnectionInvitation(id, invitationJSON string) (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if id != "" && invitationJSON != "" {
		s.invitations[id] = invitationJSON
		return id, nil
	}
	return "", errors.New("cannot add non-existent connection invitation")
}

func (s *ConnectionStore) GetConnectionInvitation(id string) (string, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if invitation, ok := s.invitations[id]; ok {
		return invitation, nil
	}
	return "", fmt.Errorf("connection invitation by the id %s not found", id)
}
