package keygen

import (
	"github.com/okx/threshold-lib/tss"
	"github.com/okx/threshold-lib/tss/key/dkg"
	"helloworld.com/okx_mpc/common"
)

type round2C struct {
	*common.Helper
	Info *dkg.SetupInfo
}

func (r *round2C) get_tmsgs() []*tss.Message {
	var msg_list []*tss.Message

	for _, msg := range r.Msgs[r.Protocol][r.Number()] {
		msg_list = append(msg_list, r.Msg2Tssmsg(msg))
	}
	return msg_list
}

func (r *round2C) Finalize() common.Round {
	msgsn_2_in := r.get_tmsgs()
	msgsn_2, _ := r.Info.DKGStep2(msgsn_2_in)

	for i := 1; i <= r.Info.Total; i++ {
		if i == r.MachineId {
			continue
		}
		r.SendTssMessage(msgsn_2[i], i, r.Number()+1)
	}

	round3 := &round3C{
		Helper: r.Helper,
		Info:   r.Info,
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
	return len(r.Msgs[r.Protocol][r.Number()]) == 2
}
