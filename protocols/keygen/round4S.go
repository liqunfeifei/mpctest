package keygen

import (
	"github.com/okx/threshold-lib/crypto/paillier"
	"github.com/okx/threshold-lib/tss/ecdsa/keygen"
	"github.com/okx/threshold-lib/tss/key/dkg"
	log "github.com/sirupsen/logrus"
	"helloworld.com/okx_mpc/common"
)

type round4S struct {
	*common.Helper
	Info *dkg.SetupInfo
}

func (r *round4S) Finalize() common.Round {
	log.Infoln("Round4 start")
	p1Dto, _, _ := paillier.NewKeyPair(8)

	for i := 1; i <= r.Info.Total; i++ {
		if i == r.MachineId {
			continue
		}
		p1Data, _ := keygen.P1(r.KeyInfo.ShareI, p1Dto, r.Info.DeviceNumber, i)
		log.Infoln("p1Dto", p1Data)
		r.SendTssMessage(p1Data, i, r.Number())
	}

	return nil
}
func (r *round4S) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round4S) Number() int { return 4 }
func (r *round4S) ReceivedAll() bool {
	return true
}
