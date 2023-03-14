package main

import (
	"context"
	"flag"
	"os"

	"github.com/libp2p/go-libp2p"
	log "github.com/sirupsen/logrus"
)

var p2pProtocol = map[string]string{
	"cmpkeygen": "/cmp/keygen/1.0.0",
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&MyFormatter{})
}

func main() {
	protocolFlag := flag.String("protocol", "defaultprotocol", "sepecify a protocol")
	partyNumFlag := flag.Int("n", 3, "total parties")
	thresholdFlag := flag.Int("t", 1, "threshold")
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

	network, peers, err := InitNetwork(ctx, h, p2pProtocol[*protocolFlag], *partyNumFlag, *timeoutFlag)
	if err != nil {
		panic(err)
	}

}