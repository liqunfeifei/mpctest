package keygen

import (
	"helloworld.com/okx_mpc/common"
)

type round1S struct {
	*common.Helper
}

func (r *round1S) Finalize() common.Round {

	return nil
}
func (r *round1S) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round1S) Number() int       { return 1 }
func (r *round1S) ReceivedAll() bool { return true }
