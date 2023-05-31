package sign

import (
	"crypto/sha256"
	"encoding/hex"

	log "github.com/sirupsen/logrus"

	"github.com/fxamacker/cbor/v2"
	"github.com/liqunfeifei/mpctest/common"
	"github.com/okx/threshold-lib/tss/ecdsa/sign"
)

type round2S struct {
	*common.Helper
}

func (r *round2S) Finalize() common.Round {
	log.Println("=========2/2 sign==========")
	hash := sha256.New()
	message := hash.Sum([]byte("hello"))

	p1 := sign.NewP1(r.Pubkey, hex.EncodeToString(message), r.PaiPrivate)
	commit, _ := p1.Step1()

	data, err := cbor.Marshal(commit)
	if err != nil {
		log.Panicln("failed to marshal round message: ", err)
	}
	msg := &common.Message{
		From:        r.MachineId,
		To:          r.Signer,
		Protocol:    r.Protocol,
		RoundNumber: r.Number(),
		Data:        data,
	}
	r.SendMessage(msg, r.Signer)

	round3 := &round3S{
		Helper: r.Helper,
		p1:     p1,
	}
	return round3
}

func (r *round2S) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round2S) Proto() string     { return r.Protocol }
func (r *round2S) Number() int       { return 2 }
func (r *round2S) ReceivedAll() bool { return true }
