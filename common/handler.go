package common

import (
	log "github.com/sirupsen/logrus"
)

func HandlerLoop(firstround Round, n *Network) {
	r := firstround
	for {
		if r.ReceivedAll() {
			r := r.Finalize()
			if r == nil {
				log.Info("Round finish.")
				break
			}
		}
		msgIn := <-n.inChan
		r.StoreMessage(msgIn)
	}
}
