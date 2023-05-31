package sign

import (
	"github.com/liqunfeifei/mpctest/common"
)

type round4C struct {
	*common.Helper
}

func (r *round4C) Finalize() common.Round {

	return nil
}
func (r *round4C) StoreMessage(msg *common.Message) error {
	common.DumpMsg(msg)
	r.SaveMessage(msg)
	return nil
}
func (r *round4C) Proto() string     { return r.Protocol }
func (r *round4C) Number() int       { return 1 }
func (r *round4C) ReceivedAll() bool { return true }
