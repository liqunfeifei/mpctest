package common

import (
	"github.com/libp2p/go-libp2p/core/peer"
)

type Helper struct {
	machineId int
	network   Network
	peerId    map[int]peer.AddrInfo
}

func (h *Helper) SendMessage(msg *Message, content []byte, to int) error {
	h.network.SendMessage(to, msg)
	return nil
}

func (h *Helper) Peers() []peer.AddrInfo {
	return h.network.peers
}
