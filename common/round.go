package common

type Round interface {
	Finalize() error
	StoreMessage(msg Message) error
	Number() uint16
	ReceivedAll() bool
}
