package event

type Subscriber interface {
	Events() map[string]Listener
}
