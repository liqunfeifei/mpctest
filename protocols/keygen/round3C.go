package keygen

import (
	"fmt"

	"github.com/okx/threshold-lib/tss"
	"github.com/okx/threshold-lib/tss/key/dkg"
	"helloworld.com/okx_mpc/common"
)

type round3C struct {
	*common.Helper
	Info *dkg.SetupInfo
}

func (r *round3C) get_tmsgs() []*tss.Message {
	var tmsg_list []*tss.Message

	for _, msg := range r.Msgs[r.Protocol][r.Number()] {
		tmsg_list = append(tmsg_list, r.Msg2Tssmsg(msg))
	}
	return tmsg_list
}

func (r *round3C) Finalize() common.Round {
	msgsn_3_in := r.get_tmsgs()
	pnSaveData, _ := r.Info.DKGStep3(msgsn_3_in)
	r.KeyInfo = pnSaveData
	fmt.Println("setUp1", pnSaveData, pnSaveData.PublicKey)
	round4 := &round4C{
		Helper: r.Helper,
		Info:   r.Info,
	}
	return round4
}
func (r *round3C) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round3C) Number() int { return 3 }
func (r *round3C) ReceivedAll() bool {
	return len(r.Msgs[r.Protocol][r.Number()]) == 2
}
