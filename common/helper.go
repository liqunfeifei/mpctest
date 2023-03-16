package common

import (
	"github.com/libp2p/go-libp2p/core/peer"
	log "github.com/sirupsen/logrus"
)

type Helper struct {
	PeerId    map[int]peer.ID
	Protocol  string
	MachineId int
	Net       *Network
	Msgs      map[int][]*Message
}

func (h *Helper) SendMessage(msg *Message, to int) error {
	h.Net.SendMessage(to, msg)
	return nil
}

func (h *Helper) SaveMessage(msg *Message) {
	h.Msgs[msg.RoundNumber] = append(h.Msgs[msg.RoundNumber], msg)
}

func (h *Helper) Peers() []peer.AddrInfo {
	return h.Net.peers
}

func DumpMsg(msg *Message) {
	log.Infof("\nFrom: %d \t To:%d\nProtocol:%s\t Round:%d\nData:[%d]%s\n",
		msg.From, msg.To, msg.Protocol, msg.RoundNumber, len(msg.Data), msg.Data)
}

func NewMsgQueue(rounds int) map[int][]*Message {
	msgMap := make(map[int][]*Message)

	for i := 1; i <= rounds; i++ {
		msgMap[i] = make([]*Message, 0)
	}
	return msgMap
}
