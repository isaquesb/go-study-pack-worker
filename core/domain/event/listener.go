package event

type Listener interface {
	Handle(e Event) error
}
