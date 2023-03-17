package keygen

import (
	"crypto/ecdsa"

	"helloworld.com/okx_mpc/common"
)

type round1C struct {
	*common.Helper
}

type message1 struct {
	PubKey *ecdsa.PublicKey
}

func (r *round1C) Finalize() common.Round {
	// //send pubkey
	// data, err := cbor.Marshal(r.Pubkey)
	// if err != nil {
	// 	log.Panicf("failed to marshal round message: %w", err)
	// }
	// msg := &common.Message{
	// 	From:        r.MachineId,
	// 	To:          1,
	// 	Protocol:    r.Protocol,
	// 	RoundNumber: 1,
	// 	Data:        data,
	// }
	// r.SendMessage(msg, 1)

	// x2 := r.Tsskey.ShareI()
	// log.Infoln("=========2/2 sign==========")
	// hash := sha256.New()
	// message := hash.Sum([]byte("hello"))

	// p2 := sign.NewP2(x2, r.P2SaveData.E_x1, r.Pubkey, r.P2SaveData.PaiPubKey, hex.EncodeToString(message))

	// bobProof, R2, _ := p2.Step1(commit)

	return nil
}
func (r *round1C) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round1C) Number() int       { return 1 }
func (r *round1C) ReceivedAll() bool { return true }
