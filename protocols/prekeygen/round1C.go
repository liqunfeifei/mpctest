package prekeygen

import (
	"log"

	"github.com/fxamacker/cbor/v2"
	"helloworld.com/okx_mpc/common"
)

type round1C struct {
	*common.Helper
}

func (r *round1C) Finalize() common.Round {
	var data message1
	msg := r.Msgs[r.Protocol][r.Number()][0]
	cbor.Unmarshal(msg.Data, &data)

	for id, mid := range data.Parties {
		for _, addr := range r.Peers() {
			if id == r.Net.SelfID.String() {
				log.Println("myid:", id, "mid:", mid)
				r.MachineId = mid
				break
			} else if id == addr.ID.String() {
				log.Println("id:", id, "mid:", mid)
				r.PeerId[mid] = addr.ID
				r.Net.PeerID[mid] = addr.ID
				break
			}
		}
	}
	return nil
}
func (r *round1C) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round1C) Proto() string { return r.Protocol }
func (r *round1C) Number() int   { return 1 }
func (r *round1C) ReceivedAll() bool {
	return len(r.Msgs[r.Protocol][r.Number()]) == 1
}
