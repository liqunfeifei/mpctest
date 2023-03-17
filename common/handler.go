package common

import (
	log "github.com/sirupsen/logrus"
)

func HandlerLoop(firstround Round, n *Network) {
	r := firstround
	for {
		if r.ReceivedAll() {
			log.Debugf("Round%d Finalize start", r.Number())
			r = r.Finalize()
			if r == nil {
				log.Info("Protocol finish.")
				break
			}
		} else {
			msgIn := <-n.inChan
			r.StoreMessage(msgIn)
		}

	}
}
