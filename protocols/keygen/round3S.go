package keygen

import (
	"fmt"

	"github.com/okx/threshold-lib/tss"
	"github.com/okx/threshold-lib/tss/key/dkg"
	"helloworld.com/okx_mpc/common"
)

type round3S struct {
	*common.Helper
	Info *dkg.SetupInfo
}

func (r *round3S) get_tmsgs() []*tss.Message {
	var tmsg_list []*tss.Message

	for _, msg := range r.Msgs[r.Protocol][r.Number()] {
		tmsg_list = append(tmsg_list, r.Msg2Tssmsg(msg))
	}
	return tmsg_list
}

func (r *round3S) Finalize() common.Round {
	msgs1_3_in := r.get_tmsgs()
	p1SaveData, _ := r.Info.DKGStep3(msgs1_3_in)
	r.KeyInfo = p1SaveData
	fmt.Println("setUp1", p1SaveData, p1SaveData.PublicKey)
	round4 := &round4S{
		Helper: r.Helper,
		Info:   r.Info,
	}
	return round4
}
func (r *round3S) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round3S) Number() int { return 3 }
func (r *round3S) ReceivedAll() bool {
	return len(r.Msgs[r.Protocol][r.Number()]) == 2
}
