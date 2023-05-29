package common

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/okx/threshold-lib/crypto/paillier"
	"github.com/okx/threshold-lib/tss"
	"github.com/okx/threshold-lib/tss/ecdsa/keygen"
	"github.com/okx/threshold-lib/tss/key/bip32"
	log "github.com/sirupsen/logrus"
)

const (
	ProtocolPrekeygen = "Prekeygen"
	ProtocolKeygen    = "Keygen"
	ProtocolSign      = "Sign"
	ProtocolBip32     = "Bip32"
	ProtocolRefresh   = "Refresh"
)

type ProtocolMsgs map[int][]*Message

type Helper struct {
	PeerId    map[int]peer.ID
	Protocol  string
	MachineId int
	Net       *Network
	Msgs      map[string]ProtocolMsgs

	Signer     int
	KeyInfo    *tss.KeyStep3Data
	P2SaveData *keygen.P2SaveData
	PaiPrivate *paillier.PrivateKey
	Pubkey     *ecdsa.PublicKey
	Tsskey     *bip32.TssKey
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
	h.Msgs[msg.Protocol][msg.RoundNumber] = append(h.Msgs[msg.Protocol][msg.RoundNumber], msg)
	// h.Msgs[msg.RoundNumber] = append(h.Msgs[msg.RoundNumber], msg)
}

func (h *Helper) Peers() []peer.AddrInfo {
	return h.Net.peers
}

func DumpMsg(msg *Message) {
	log.Infof("\nFrom: %d \t To:%d\nProtocol:%s\t Round:%d\nData:[%d]%s\n",
		msg.From, msg.To, msg.Protocol, msg.RoundNumber, len(msg.Data), msg.Data)
}

func NewMsgQueue() map[string]ProtocolMsgs {
	round := 10

	msgMap := make(map[string]ProtocolMsgs)

	for _, p := range []string{ProtocolPrekeygen, ProtocolKeygen, ProtocolSign, ProtocolBip32, ProtocolRefresh} {
		msgMap[p] = make(ProtocolMsgs)
		for i := 1; i <= round; i++ {
			msgMap[p][i] = make([]*Message, 0)
		}
	}
	return msgMap
}
