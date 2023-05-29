package keygen

import (
	"crypto/ecdsa"

	"github.com/okx/threshold-lib/crypto/curves"
	"github.com/okx/threshold-lib/tss"
	"github.com/okx/threshold-lib/tss/ecdsa/keygen"
	"github.com/okx/threshold-lib/tss/key/bip32"
	"github.com/okx/threshold-lib/tss/key/dkg"
	log "github.com/sirupsen/logrus"
	"helloworld.com/okx_mpc/common"
)

type round4C struct {
	*common.Helper
	Info *dkg.SetupInfo
}

func (r *round4C) get_tmsgs() []*tss.Message {
	var tmsg_list []*tss.Message

	for _, msg := range r.Msgs[r.Protocol][r.Number()] {
		tmsg_list = append(tmsg_list, r.Msg2Tssmsg(msg))
	}
	return tmsg_list
}

func (r *round4C) Finalize() common.Round {
	p1Data := r.get_tmsgs()[0]
	publicKey, _ := curves.NewECPoint(curve, r.KeyInfo.PublicKey.X, r.KeyInfo.PublicKey.Y)
	pnData, error := keygen.P2(r.KeyInfo.ShareI, publicKey, p1Data, 1, r.MachineId)
	if error != nil {
		log.Errorln(error)
	}
	// log.Infoln("p", r.MachineId, "Data", pnData)

	r.P2SaveData = pnData
	// log.Infoln("P", r.MachineId, pnData)

	log.Infoln("=========bip32==========")
	tssKey, err := bip32.NewTssKey(pnData.X2, r.KeyInfo.PublicKey, r.KeyInfo.ChainCode)
	if err != nil {
		log.Errorln(err)
	}
	tssKey, err = tssKey.NewChildKey(996)
	if err != nil {
		log.Errorln(err)
	}
	r.Tsskey = tssKey
	// x2 := tssKey.ShareI()
	pubKey := &ecdsa.PublicKey{Curve: curve, X: tssKey.PublicKey().X, Y: tssKey.PublicKey().Y}
	r.Pubkey = pubKey

	return nil
}

func (r *round4C) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round4C) Proto() string { return r.Protocol }
func (r *round4C) Number() int   { return 4 }
func (r *round4C) ReceivedAll() bool {
	return len(r.Msgs[r.Protocol][r.Number()]) == 1
}
