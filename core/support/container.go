package support

import (
	"github.com/sarulabs/di"
	"worker/core/port"
	"worker/core/port/event"
)

var Container = Application{}

type Application struct {
	Container di.Container
}

func (app Application) Logger() port.Logger {
	return app.Container.Get("logger").(port.Logger)
}

func (app Application) EventManager() event.Manager {
	return app.Container.Get("eventManager").(event.Manager)
}

func (app Application) Messenger() port.Messenger {
	return app.Container.Get("messenger").(port.Messenger)
}
