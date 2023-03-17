package keygen

import (
	"github.com/okx/threshold-lib/tss/key/dkg"
	"helloworld.com/okx_mpc/common"
)

type round1S struct {
	*common.Helper
	Info *dkg.SetupInfo
}

func (r *round1S) Finalize() common.Round {
	msgs1_1, _ := r.Info.DKGStep1()

	for i := 1; i <= r.Info.Total; i++ {
		if i == r.MachineId {
			continue
		}
		r.SendTssMessage(msgs1_1[i], i, r.Number()+1)
	}

	round2 := &round2S{
		Helper: r.Helper,
		Info:   r.Info,
	}

	return round2
}
func (r *round1S) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round1S) Number() int       { return 1 }
func (r *round1S) ReceivedAll() bool { return true }
