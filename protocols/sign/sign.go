package sign

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/liqunfeifei/mpctest/common"
)

var (
	curve = btcec.S256()
)

// var maxMsgCount = 10

func StartSignS(n *common.Network, helper *common.Helper) common.Round {
	helper.Protocol = common.ProtocolSign
	r1 := &round1S{
		Helper: helper,
	}
	return r1
}

func StartSignC(n *common.Network, helper *common.Helper) common.Round {
	helper.Protocol = common.ProtocolSign
	r1 := &round1C{
		Helper: helper,
	}
	return r1
}
