package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	mpcProtocol "github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

type Network struct {
	host     host.Host
	ctx      context.Context
	protocol string
	peerID   map[string]peer.ID
	selfID   peer.ID
	inChan   chan *mpcProtocol.Message
	outChan  chan *mpcProtocol.Message
}

func InitNetwork(ctx context.Context, h host.Host, protocolName string, parties int, timeout int) (*Network, []peer.AddrInfo, error) {
	n := &Network{
		host:     h,
		ctx:      ctx,
		protocol: protocolName,
		peerID:   make(map[string]peer.ID),
		selfID:   h.ID(),
		inChan:   make(chan *mpcProtocol.Message, 20),
		outChan:  make(chan *mpcProtocol.Message, 20),
	}

	h.SetStreamHandler(protocol.ID(p2pProtocol[protocolName]), n.handleStream)

	peerids := waitingForParticipants(ctx, h, parties-1, timeout, protocolName)
	if peerids == nil {
		log.Infoln("Waiting for participants timeout!")
		return nil, nil, fmt.Errorf("timeout")
	}

	for _, p := range peerids {
		n.peerID[p.ID.String()] = p.ID
	}

	return n, peerids, nil
}

func (n *Network) SendMessage(id string, msg *mpcProtocol.Message) {
	// fmt.Println("send msg to: ", id)

	data, err := msg.MarshalBinary()
	// fmt.Println("Marshal data: ", data)
	if err != nil {
		log.Panicln("Marshal failed! ", err)
	}

	s, err := n.host.NewStream(n.ctx, n.peerID[id], protocol.ID(n.protocol))
	if err != nil {
		log.Panicln("Stream open failed! ", err)
	}
	w := bufio.NewWriter(s)
	_, err = w.Write(data)
	if err != nil {
		log.Panicln("Write data failed! ", err)
	}
	if err = w.Flush(); err != nil {
		log.Panicln("Flush failed! ", err)
	}
	s.Close()
}

func (n *Network) handleStream(s network.Stream) {
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	go readData(rw, n.inChan)
}

func (n *Network) Quit() {
	n.host.Close()
}

func readData(rw *bufio.ReadWriter, ch chan *mpcProtocol.Message) {
	buf, err := io.ReadAll(rw)
	if err != nil {
		log.Panicln("Error reading from buffer. ", err)
	}
	log.Debugf("New stream! (%d Bytes)", len(buf))
	// fmt.Println("received data: ", buf)
	msg := new(mpcProtocol.Message)
	err = msg.UnmarshalBinary(buf)
	if err != nil {
		log.Panicln("Unmarshal failed.", err)
	}
	ch <- msg
}

func waitingForParticipants(ctx context.Context, h host.Host, num int, timeout int, channel string) []peer.AddrInfo {
	pich := make(chan peer.AddrInfo, 20)
	peers := []peer.AddrInfo{}
	// setup local mDNS discovery
	if err := setupDiscovery(ctx, h, channel, pich, num); err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Second * 20)

	for {
		select {
		case <-ticker.C:
			log.Println("Timeout!")
			ticker.Stop()
			return nil
		case pi, ok := <-pich:
			if !ok {
				return peers
			}
			log.Debug("append peers.")
			peers = append(peers, pi)
		}
	}
}

// // shortID returns the last 8 chars of a base58-encoded peer id.
// func shortID(p peer.ID) string {
// 	pretty := p.Pretty()
// 	return pretty[len(pretty)-8:]
// }

type discoveryNotifee struct {
	ctx        context.Context
	h          host.Host
	pich       chan peer.AddrInfo
	threshold  int
	peersCount int
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	log.Infoln("Discovered new peer %s\n", pi.ID.String())
	err := n.h.Connect(n.ctx, pi)
	if err != nil {
		log.Infof("error connecting to peer %s: %s\n", pi.ID.String(), err)
	} else {
		log.Infof("Connected to peer %s\n", pi.ID.String())
		n.pich <- pi
		n.peersCount += 1
		if n.peersCount >= n.threshold {
			log.Infoln("Found", n.threshold, " peers, stop scan.")
			close(n.pich)
		}
	}
}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupDiscovery(ctx context.Context, h host.Host, group string, ch chan peer.AddrInfo, t int) error {
	log.Infoln("Register mDNS.")
	// setup mDNS discovery to find local peers
	s := mdns.NewMdnsService(h, group, &discoveryNotifee{ctx: ctx, h: h, pich: ch, threshold: t, peersCount: 0})
	return s.Start()
}
