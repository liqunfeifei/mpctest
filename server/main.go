package main

import (
	"context"
	"flag"
	"os"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/peer"
	log "github.com/sirupsen/logrus"
	"helloworld.com/okx_mpc/common"
)

var p2pProtocol = map[string]string{
	"cmpkeygen": "/cmp/keygen/1.0.0",
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&common.MyFormatter{})
}

func main() {
	protocolFlag := flag.String("protocol", "defaultprotocol", "sepecify a protocol")
	partyNumFlag := flag.Int("n", 3, "total parties")
	thresholdFlag := flag.Int("t", 2, "threshold")
	machineIDFlag := flag.Int("i", 0, "machine id")
	timeoutFlag := flag.Int("timeout", 20, "timeout(seconds)")

	flag.Parse()

	log.Debug("Hello world")

	ctx := context.Background()

	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}

	log.Infoln("My ID: ", h.ID().String())

	network, peers, err := common.InitNetwork(ctx, h, p2pProtocol[*protocolFlag], *partyNumFlag, *timeoutFlag)
	if err != nil {
		panic(err)
	}

	startAll(*thresholdFlag, network, []byte("hello"), peers, *machineIDFlag)
}

func startAll(threshold int, n *common.Network, message []byte, peers []peer.AddrInfo, mId int) error {
	return nil
}
