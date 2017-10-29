package server

import (
	"errors"

	"github.com/cashshuffle/cashshuffle/message"
	"github.com/nats-io/nuid"
)

// registerClient registers a new session.
func registerClient(data *signedConn) error {
	if data.message.GetSignature().String() == "" {
		p := data.message.GetPacket()
		if p.From.String() != "" {
			n := nuid.New()

			td := trackerData{
				verificationKey: p.From.String(),
				sessionID:       []byte(n.Next()),
				conn:            data.conn,
				number:          0,
			}
			data.tracker.add(data.conn, &td)

			err := registrationSuccess(data, &td)
			return err
		}
	}

	if err := registrationFailed(data); err != nil {
		return err
	}

	return errors.New("registration failed")
}

// registrationSuccess sends a registration success reply.
func registrationSuccess(data *signedConn, td *trackerData) error {
	m := message.Signed{
		Packet: &message.Packet{
			Session: td.sessionID,
			Number:  td.number,
		},
	}

	err := writeMessage(data.conn, &m)
	return err
}

// registrationFailed sends a registration failed reply.
func registrationFailed(data *signedConn) error {
	m := message.Signed{
		Packet: &message.Packet{
			Message: &message.Message{
				Blame: &message.Blame{
					Reason: message.Reason_INVALIDFORMAT,
				},
			},
		},
	}

	err := writeMessage(data.conn, &m)
	return err
}
