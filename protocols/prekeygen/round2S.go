package prekeygen

import "helloworld.com/okx_mpc/common"

type round2 struct {
	*common.Helper
}

func (r *round2) Finalize() (*round1S, error) {
	return nil, nil
}
func (r *round2) StoreMessage(common.Message) error { return nil }
func (r *round2) Number() int                       { return 2 }
func (r *round2) ReceivedAll() bool                 { return true }
