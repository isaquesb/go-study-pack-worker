package support

import (
	"worker/core/port"
	"worker/core/port/event"
)

var Container = Store{}

type Store struct {
	EventManager event.Manager
	Logger port.Logger
	Messenger port.Messenger
}
