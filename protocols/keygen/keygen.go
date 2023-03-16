package keygen

import (
	"crypto/elliptic"

	"github.com/libp2p/go-libp2p/core/peer"
	"helloworld.com/okx_mpc/common"
)

var protocol = "Keygen"
var totalRound = 1

// var maxMsgCount = 10

func StartKeygeS(n *common.Network, curve elliptic.Curve) *round1S {
	info := common.SetupInfo{
		DeviceNumber: 0,
		Threshold:    2,
		Total:        3,
		RoundNumber:  1,
		curve:        curve,
	}
	helper := common.Helper{
		Protocol:  protocol,
		Net:       n,
		PeerId:    make(map[int]peer.ID),
		MachineId: 0,
		Msgs:      newMsgQueue(totalRound),
	}

	r1 := &round1S{
		Helper: &helper,
		Info:   &info,
	}
	return r1
}

func StartKeygeC(n *common.Network, curve elliptic.Curve) *round1C {
	info := SetupInfo{
		Threshold:   2,
		Total:       3,
		RoundNumber: 1,
		curve:       curve,
	}
	helper := common.Helper{
		Protocol: protocol,
		Net:      n,
		PeerId:   make(map[int]peer.ID),
		Msgs:     newMsgQueue(totalRound),
	}

	r1 := &round1C{
		Helper: &helper,
		Info:   &info,
	}
	return r1
}

func newMsgQueue(rounds int) map[int][]*common.Message {

	msgMap := make(map[int][]*common.Message)

	for i := 1; i <= rounds; i++ {
		msgMap[i] = make([]*common.Message, 0)
	}
	return msgMap
}
