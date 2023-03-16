package keygen

import (
	"helloworld.com/okx_mpc/common"
)

type round1C struct {
	*common.Helper
	Info *SetupInfo
}

func (r *round1C) Finalize() common.Round {
	return nil
}
func (r *round1C) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round1C) Number() int { return 1 }
func (r *round1C) ReceivedAll() bool {
	if len(r.Msgs[r.Number()]) >= 1 {
		return true
	}
	return false
}
