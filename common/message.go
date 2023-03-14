package common

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
)

type networkMessage struct {
	From    int
	To      int
	Content []byte
}

type Message struct {
	// From is the party.ID of the sender
	From int
	// To is the intended recipient for this message. If To == "", then the message should be sent to all.
	To int
	// Protocol identifies the protocol this message belongs to
	Protocol string
	// RoundNumber is the index of the round this message belongs to
	RoundNumber int
	// Data is the actual content consumed by the round.
	Data []byte
	// Broadcast indicates whether this message should be reliably broadcast to all participants.
}

// String implements fmt.Stringer.
func (m Message) String() string {
	return fmt.Sprintf("message: round %d, from: %s, to %v, protocol: %s", m.RoundNumber, m.From, m.To, m.Protocol)
}

// IsFor returns true if the message is intended for the designated party.
func (m Message) IsFor(id int) bool {
	if m.From == id {
		return false
	}
	return m.To == 0 || m.To == id
}

// // Hash returns a 64 byte hash of the message content, including the headers.
// // Can be used to produce a signature for the message.
// func (m *Message) Hash() []byte {
// 	var broadcast byte
// 	if m.Broadcast {
// 		broadcast = 1
// 	}
// 	h := hash.New(
// 		hash.BytesWithDomain{TheDomain: "SSID", Bytes: m.SSID},
// 		m.From,
// 		m.To,
// 		hash.BytesWithDomain{TheDomain: "Protocol", Bytes: []byte(m.Protocol)},
// 		m.RoundNumber,
// 		hash.BytesWithDomain{TheDomain: "Content", Bytes: m.Data},
// 		hash.BytesWithDomain{TheDomain: "Broadcast", Bytes: []byte{broadcast}},
// 		hash.BytesWithDomain{TheDomain: "BroadcastVerification", Bytes: m.BroadcastVerification},
// 	)
// 	return h.Sum()
// }

// marshallableMessage is a copy of message for the purpose of cbor marshalling.
//
// This is a workaround to use cbor's default marshalling for Message, all while providing
// a MarshalBinary method
type marshallableMessage struct {
	From        int
	To          int
	Protocol    string
	RoundNumber int
	Data        []byte
}

func (m *Message) toMarshallable() *marshallableMessage {
	return &marshallableMessage{
		From:        m.From,
		To:          m.To,
		Protocol:    m.Protocol,
		RoundNumber: m.RoundNumber,
		Data:        m.Data,
	}
}

func (m *Message) MarshalBinary() ([]byte, error) {
	return cbor.Marshal(m.toMarshallable())
}

func (m *Message) UnmarshalBinary(data []byte) error {
	deserialized := m.toMarshallable()
	if err := cbor.Unmarshal(data, deserialized); err != nil {
		return nil
	}
	m.From = deserialized.From
	m.To = deserialized.To
	m.Protocol = deserialized.Protocol
	m.RoundNumber = deserialized.RoundNumber
	m.Data = deserialized.Data
	return nil
}
