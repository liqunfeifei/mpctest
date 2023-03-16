package keygen

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/okx/threshold-lib/tss/key/dkg"
	"helloworld.com/okx_mpc/common"
)

var (
	curve      = btcec.S256()
	protocol   = "Keygen"
	totalRound = 1
)

// var maxMsgCount = 10

func StartKeygeS(n *common.Network, helper *common.Helper) common.Round {
	helper.Protocol = protocol
	helper.Msgs = common.NewMsgQueue(totalRound)

	r1 := &round1S{
		Helper: helper,
		Info:   dkg.NewSetUp(helper.MachineId, 3, curve),
	}
	return r1
}

func StartKeygeC(n *common.Network, helper *common.Helper) common.Round {
	helper.Protocol = protocol
	helper.Msgs = common.NewMsgQueue(totalRound)

	r1 := &round1C{
		Helper: helper,
		Info:   dkg.NewSetUp(helper.MachineId, 3, curve),
	}
	return r1
}
