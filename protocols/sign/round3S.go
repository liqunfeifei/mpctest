package sign

import (
	"math/big"

	"github.com/fxamacker/cbor/v2"
	"github.com/okx/threshold-lib/crypto/curves"
	"github.com/okx/threshold-lib/crypto/schnorr"
	"github.com/okx/threshold-lib/tss/ecdsa/sign"
	log "github.com/sirupsen/logrus"
	"helloworld.com/okx_mpc/common"
)

type round3S struct {
	*common.Helper
	p1 *sign.P1Context
}

type Message3 struct {
	Proof MessageProof
	CmtD  []*big.Int
}

func (r *round3S) Finalize() common.Round {
	var data Message2
	msg := r.Msgs[r.Protocol][r.Number()][0]
	cbor.Unmarshal(msg.Data, &data)

	p := &curves.ECPoint{
		Curve: curve,
		X:     data.Proof.R.X,
		Y:     data.Proof.R.Y,
	}
	bobProof := &schnorr.Proof{
		R: p,
		S: data.Proof.S,
	}
	R2 := &curves.ECPoint{
		Curve: curve,
		X:     data.R2.X,
		Y:     data.R2.Y,
	}
	proof, cmtD, _ := r.p1.Step2(bobProof, R2)

	message_s := &Message3{
		Proof: MessageProof{
			R: MessagePoint{
				X: proof.R.X,
				Y: proof.R.Y,
			},
			S: proof.S,
		},
		CmtD: *cmtD,
	}
	// log.Debugln("proof.R.X:", message_s.Proof.R.X)
	// log.Debugln("proof.R.Y:", message_s.Proof.R.Y)
	// log.Debugln("proof.S:", message_s.Proof.S)
	// log.Debugln("cmtD:", message_s.CmtD)

	b, err := cbor.Marshal(message_s)
	if err != nil {
		log.Panicln("failed to marshal round message: ", err)
	}
	// log.Debugln("message_s: ", b)
	msg_s := &common.Message{
		From:        r.MachineId,
		To:          r.Signer,
		Protocol:    r.Protocol,
		RoundNumber: r.Number(),
		Data:        b,
	}
	r.SendMessage(msg_s, r.Signer)

	round4 := &round4S{
		Helper: r.Helper,
		p1:     r.p1,
	}
	return round4
}
func (r *round3S) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round3S) Proto() string { return r.Protocol }
func (r *round3S) Number() int   { return 3 }
func (r *round3S) ReceivedAll() bool {
	return len(r.Msgs[r.Protocol][r.Number()]) == 1
}
