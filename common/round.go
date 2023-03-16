package common

type Round interface {
	Finalize() Round
	StoreMessage(msg *Message) error
	Number() int
	ReceivedAll() bool
}
