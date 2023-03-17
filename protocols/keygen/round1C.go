package keygen

import (
	"github.com/okx/threshold-lib/tss/key/dkg"
	"helloworld.com/okx_mpc/common"
)

type round1C struct {
	*common.Helper
	Info *dkg.SetupInfo
}

func (r *round1C) Finalize() common.Round {
	msgsn_1, _ := r.Info.DKGStep1()

	for i := 1; i <= r.Info.Total; i++ {
		if i == r.MachineId {
			continue
		}
		r.SendTssMessage(msgsn_1[i], i, r.Number()+1)
	}

	round2 := &round2C{
		Helper: r.Helper,
		Info:   r.Info,
	}

	return round2
}
func (r *round1C) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round1C) Number() int { return 1 }
func (r *round1C) ReceivedAll() bool {
	return true
}
