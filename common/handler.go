package common

import (
	log "github.com/sirupsen/logrus"
)

func HandlerLoop(firstround Round, n *Network) {
	r := firstround
	for {
		if !r.ReceivedAll() {
			msgIn := <-n.inChan
			r.StoreMessage(msgIn)
		}
		log.Debugf("Round%d Finalize start", r.Number())
		r = r.Finalize()
		if r == nil {
			log.Info("Round finish.")
			break
		}
	}
}
