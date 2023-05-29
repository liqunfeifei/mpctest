package sign

import (
	"crypto/ecdsa"

	"github.com/fxamacker/cbor/v2"
	"helloworld.com/okx_mpc/common"
)

type round1S struct {
	*common.Helper
}

func (r *round1S) Finalize() common.Round {
	var data message1
	msg := r.Msgs[r.Protocol][r.Number()][0]
	cbor.Unmarshal(msg.Data, &data)

	r.Pubkey = &ecdsa.PublicKey{
		Curve: curve,
		X:     &data.X,
		Y:     &data.Y,
	}
	round2 := &round2S{
		Helper: r.Helper,
	}
	return round2
}
func (r *round1S) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round1S) Proto() string { return r.Protocol }
func (r *round1S) Number() int   { return 1 }
func (r *round1S) ReceivedAll() bool {
	return len(r.Msgs[r.Protocol][r.Number()]) == 1
}
