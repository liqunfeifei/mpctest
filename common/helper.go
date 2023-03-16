package common

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/okx/threshold-lib/tss"
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
	log.Infoln("Send msg, From:", msg.From, "to:", to)
	h.Net.SendMessage(to, msg)
	return nil
}

func (h *Helper) SendTssMessage(tmsg *tss.Message, to int, round int) error {
	msg := h.Tssmsg2Msg(tmsg, to, round)
	h.SendMessage(msg, to)

	return nil
}

func (h *Helper) Tssmsg2Msg(tmsg *tss.Message, to int, round int) *Message {
	data, err := cbor.Marshal(tmsg)
	if err != nil {
		panic(fmt.Errorf("failed to marshal round message: %w", err))
	}
	log.Println("Cbor data length:", len(data))

	msg := &Message{
		From:        h.MachineId,
		To:          to,
		Protocol:    h.Protocol,
		RoundNumber: round,
		Data:        data,
	}
	return msg
}

func (h *Helper) Msg2Tssmsg(msg *Message) *tss.Message {
	var tssMsg tss.Message
	cbor.Unmarshal(msg.Data, &tssMsg)
	return &tssMsg
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
