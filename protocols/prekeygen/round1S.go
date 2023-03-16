package prekeygen

import (
	"fmt"
	"log"

	"github.com/fxamacker/cbor/v2"
	"helloworld.com/okx_mpc/common"
)

type round1S struct {
	*common.Helper
}

type message1 struct {
	Parties map[string]int
}

func (r *round1S) Finalize() common.Round {
	msg1 := message1{
		Parties: make(map[string]int),
	}
	// r.PeerId[0] = r.Net.SelfID
	// peersList[0] = r.PeerId[0].String()
	msg1.Parties[r.Net.SelfID.String()] = r.MachineId
	i := 2
	for _, p := range r.Peers() {
		r.PeerId[i] = p.ID
		r.Net.PeerID[i] = p.ID
		msg1.Parties[r.PeerId[i].String()] = i
		log.Println("peer", i, " ", r.PeerId[i].String())
		i += 1
	}

	data, err := cbor.Marshal(msg1)
	if err != nil {
		panic(fmt.Errorf("failed to marshal round message: %w", err))
	}
	log.Println("Cbor data length:", len(data))

	for id := range r.PeerId {
		msg := &common.Message{
			From:        r.MachineId,
			To:          id,
			Protocol:    r.Protocol,
			RoundNumber: 1,
			Data:        data,
		}
		r.SendMessage(msg, id)
	}

	return nil
}

func (r *round1S) StoreMessage(msg *common.Message) error {
	r.SaveMessage(msg)
	return nil
}
func (r *round1S) Number() int       { return 1 }
func (r *round1S) ReceivedAll() bool { return true }
