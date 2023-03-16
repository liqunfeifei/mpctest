package keygen

import (
	"helloworld.com/okx_mpc/common"
)

type round1S struct {
	*common.Helper
	Info *common.SetupInfo
}

func (r *round1S) Finalize() (*round1S, error) {
	return nil, nil
}
func (r *round1S) StoreMessage(common.Message) error { return nil }
func (r *round1S) Number() int                       { return 2 }
func (r *round1S) ReceivedAll() bool                 { return true }
