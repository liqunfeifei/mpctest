package keygen

import (
	"github.com/liqunfeifei/mpctest/common"
	"github.com/okx/threshold-lib/tss"
	"github.com/okx/threshold-lib/tss/key/dkg"
)

type round2S struct {
	*common.Helper
	Info *dkg.SetupInfo
}

func (r *round2S) get_tmsgs() []*tss.Message {
	var msg_list []*tss.Message

	for _, msg := range r.Msgs[r.Protocol][r.Number()] {
		msg_list = append(msg_list, r.Msg2Tssmsg(msg))
	}
	return msg_list
}

func (r *round2S) Finalize() common.Round {
	msgs1_2_in := r.get_tmsgs()
	msgs1_2, _ := r.Info.DKGStep2(msgs1_2_in)

	for i := 1; i <= r.Info.Total; i++ {
		if i == r.MachineId {
			continue
		}
		r.SendTssMessage(msgs1_2[i], i, r.Number()+1)
	}

	round3 := &round3S{
		Helper: r.Helper,
		Info:   r.Info,
	}

	return round3
}
func (r *round2S) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round2S) Proto() string { return r.Protocol }
func (r *round2S) Number() int   { return 2 }
func (r *round2S) ReceivedAll() bool {
	return len(r.Msgs[r.Protocol][r.Number()]) == 2
}
