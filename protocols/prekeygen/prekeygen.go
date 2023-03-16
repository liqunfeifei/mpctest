package prekeygen

import (
	"helloworld.com/okx_mpc/common"
)

var protocol = "Prekeygen"
var totalRound = 1

// var maxMsgCount = 10

func StartPrekeygeS(n *common.Network, info *common.SetupInfo, helper *common.Helper) *round1S {

	helper.Protocol = protocol
	helper.MachineId = 0
	helper.Msgs = common.NewMsgQueue(totalRound)

	info.RoundNumber = 1
	info.DeviceNumber = 1

	r1 := &round1S{
		Helper: helper,
		Info:   info,
	}
	return r1
}

func StartPrekeygeC(n *common.Network, info *common.SetupInfo, helper *common.Helper) *round1C {

	helper.Protocol = protocol
	helper.Msgs = common.NewMsgQueue(totalRound)

	info.RoundNumber = 1

	r1 := &round1C{
		Helper: helper,
		Info:   info,
	}
	return r1
}
