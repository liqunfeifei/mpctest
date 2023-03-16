package prekeygen

import (
	"testing"

	"helloworld.com/okx_mpc/common"
)

func TestRound1(t *testing.T) {

	helper := common.Helper{
		Protocol: "123",
	}
	println(helper.Protocol)
}
