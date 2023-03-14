package prekeygen

import "helloworld.com/okx_mpc/common"

type round1 struct {
	*common.Helper
}

func (r *round1) Finalize() error {
	r.machineId = 1
	for _, p := range r.Peers() {
		r.SendMessage()
	}
	return nil
}
func (r *round1) StoreMessage(common.Message) error { return nil }
func (r *round1) Number() uint16                    { return 1 }
func (r *round1) ReceivedAll() bool                 { return true }
