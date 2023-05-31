package sign

import (
	"math/big"

	"github.com/fxamacker/cbor/v2"
	"github.com/liqunfeifei/mpctest/common"
	"github.com/okx/threshold-lib/tss/ecdsa/sign"
	log "github.com/sirupsen/logrus"
)

type round4S struct {
	*common.Helper
	p1 *sign.P1Context
}

func (r *round4S) Finalize() common.Round {
	var E_k2_h_xr *big.Int
	msg := r.Msgs[r.Protocol][r.Number()][0]
	cbor.Unmarshal(msg.Data, &E_k2_h_xr)

	r1, s, _ := r.p1.Step3(E_k2_h_xr)
	log.Infoln("Sign finished: ")
	log.Infoln("r1: ", r1)
	log.Infoln("s: ", s)
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
	return len(r.Msgs[r.Protocol][r.Number()]) == 1
}
