package gookit

import (
	"github.com/gookit/event"
	domainEvt "worker/core/domain/event"
)

type goKitListener struct {
	domainListener domainEvt.Listener
}

func (l *goKitListener) Handle(e event.Event) error {
	evt := &domEvent{
		goKitEvent: e,
	}
	return l.domainListener.Handle(evt)
}
