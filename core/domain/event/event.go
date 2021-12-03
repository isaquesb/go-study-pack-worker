package event

const EvtPackCreated = "pack.created"

type Event interface {
	Name() string
	Data() map[string]interface{}
}
