package sign

import (
	"github.com/fxamacker/cbor/v2"
	"github.com/okx/threshold-lib/crypto/curves"
	"github.com/okx/threshold-lib/crypto/schnorr"
	"github.com/okx/threshold-lib/tss/ecdsa/sign"
	log "github.com/sirupsen/logrus"
	"helloworld.com/okx_mpc/common"
)

type round3C struct {
	*common.Helper
	p2 *sign.P2Context
}

func (r *round3C) Finalize() common.Round {
	var data Message3
	msg := r.Msgs[r.Protocol][r.Number()][0]
	cbor.Unmarshal(msg.Data, &data)

	p := &curves.ECPoint{
		Curve: curve,
		X:     data.Proof.R.X,
		Y:     data.Proof.R.Y,
	}
	proof := &schnorr.Proof{
		R: p,
		S: data.Proof.S,
	}
	cmtD := data.CmtD

	// log.Debugln("proof.R.X:", data.Proof.R.X)
	// log.Debugln("proof.R.Y:", data.Proof.R.Y)
	// log.Debugln("proof.S:", data.Proof.S)
	// log.Debugln("cmtD:", data.CmtD)

	E_k2_h_xr, _ := r.p2.Step2(&cmtD, proof)

	b, err := cbor.Marshal(E_k2_h_xr)
	if err != nil {
		log.Panicln("failed to marshal round message: ", err)
	}
	msg_s := &common.Message{
		From:        r.MachineId,
		To:          1,
		Protocol:    r.Protocol,
		RoundNumber: r.Number() + 1,
		Data:        b,
	}
	r.SendMessage(msg_s, 1)

	return nil
}
func (r *round3C) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round3C) Proto() string { return r.Protocol }
func (r *round3C) Number() int   { return 3 }
func (r *round3C) ReceivedAll() bool {
	return len(r.Msgs[r.Protocol][r.Number()]) == 1
}
