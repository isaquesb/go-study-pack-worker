package support

import (
	"github.com/sarulabs/di"
	"worker/core/port"
	"worker/core/port/event"
)

var Container = Store{}

type Store struct {
	App di.Container
}

func (c Store) Logger() port.Logger {
	return c.App.Get("logger").(port.Logger)
}

func (c Store) EventManager() event.Manager {
	return c.App.Get("eventManager").(event.Manager)
}

func (c Store) Messenger() port.Messenger {
	return c.App.Get("messenger").(port.Messenger)
}
