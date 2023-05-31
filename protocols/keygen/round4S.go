package keygen

import (
	"time"

	"github.com/liqunfeifei/mpctest/common"
	"github.com/okx/threshold-lib/crypto/paillier"
	"github.com/okx/threshold-lib/tss/ecdsa/keygen"
	"github.com/okx/threshold-lib/tss/key/dkg"
	log "github.com/sirupsen/logrus"
)

type round4S struct {
	*common.Helper
	Info *dkg.SetupInfo
}

func (r *round4S) Finalize() common.Round {
	log.Infoln("Generate key pair...")
	start := time.Now()
	paiPrivate, _, _ := paillier.NewKeyPair(16)
	log.Infoln("Done.(", time.Now().Sub(start))

	r.PaiPrivate = paiPrivate

	for i := 1; i <= r.Info.Total; i++ {
		if i == r.MachineId {
			continue
		}
		p1Dto, _ := keygen.P1(r.KeyInfo.ShareI, paiPrivate, r.Info.DeviceNumber, i)
		// log.Infoln("p1Dto", p1Dto)
		r.SendTssMessage(p1Dto, i, r.Number())
	}

	return nil
}
func (r *round4S) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round4S) Proto() string { return r.Protocol }
func (r *round4S) Number() int   { return 4 }
func (r *round4S) ReceivedAll() bool {
	return true
}
