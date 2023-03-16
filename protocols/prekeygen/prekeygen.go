package prekeygen

import (
	"helloworld.com/okx_mpc/common"
)

var protocol = "Prekeygen"
var totalRound = 1

// var maxMsgCount = 10

func StartPrekeygeS(n *common.Network, helper *common.Helper) *round1S {

	helper.Protocol = protocol
	helper.MachineId = 0
	helper.Msgs = common.NewMsgQueue(totalRound)

	r1 := &round1S{
		Helper: helper,
	}
	return r1
}

func StartPrekeygeC(n *common.Network, helper *common.Helper) *round1C {

	helper.Protocol = protocol
	helper.Msgs = common.NewMsgQueue(totalRound)

	r1 := &round1C{
		Helper: helper,
	}
	return r1
}
