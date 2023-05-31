package sign

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"

	"github.com/fxamacker/cbor/v2"
	"github.com/liqunfeifei/mpctest/common"
	"github.com/okx/threshold-lib/tss/ecdsa/sign"
	log "github.com/sirupsen/logrus"
)

type round2C struct {
	*common.Helper
}

type MessagePoint struct {
	X *big.Int
	Y *big.Int
}

type MessageProof struct {
	R MessagePoint
	S *big.Int
}
type Message2 struct {
	Proof MessageProof
	R2    MessagePoint
}

func (r *round2C) Finalize() common.Round {
	hash := sha256.New()
	message := hash.Sum([]byte("hello"))

	var commit **big.Int
	msg := r.Msgs[r.Protocol][r.Number()][0]
	cbor.Unmarshal(msg.Data, &commit)

	x2 := r.Tsskey.ShareI()
	p2 := sign.NewP2(x2, r.P2SaveData.E_x1, r.Pubkey, r.P2SaveData.PaiPubKey, hex.EncodeToString(message))
	bobProof, R2, _ := p2.Step1(commit)

	message_s := &Message2{
		Proof: MessageProof{
			R: MessagePoint{
				X: bobProof.R.X,
				Y: bobProof.R.Y,
			},
			S: bobProof.S,
		},
		R2: MessagePoint{
			X: R2.X,
			Y: R2.Y,
		},
	}

	data, err := cbor.Marshal(message_s)
	if err != nil {
		log.Panicln("failed to marshal round message: ", err)
	}
	msg_s := &common.Message{
		From:        r.MachineId,
		To:          1,
		Protocol:    r.Protocol,
		RoundNumber: r.Number() + 1,
		Data:        data,
	}
	r.SendMessage(msg_s, 1)

	round3 := &round3C{
		Helper: r.Helper,
		p2:     p2,
	}
	return round3
}
func (r *round2C) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round2C) Proto() string { return r.Protocol }
func (r *round2C) Number() int   { return 2 }
func (r *round2C) ReceivedAll() bool {
	return len(r.Msgs[r.Protocol][r.Number()]) == 1
}
