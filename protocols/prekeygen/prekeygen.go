package prekeygen

import (
	"helloworld.com/okx_mpc/common"
)

// var maxMsgCount = 10

func StartPrekeygeS(n *common.Network, helper *common.Helper) common.Round {

	helper.Protocol = common.ProtocolPrekeygen
	helper.MachineId = 1

	r1 := &round1S{
		Helper: helper,
	}
	return r1
}

func StartPrekeygeC(n *common.Network, helper *common.Helper) common.Round {

	helper.Protocol = common.ProtocolPrekeygen

	r1 := &round1C{
		Helper: helper,
	}
	return r1
}
