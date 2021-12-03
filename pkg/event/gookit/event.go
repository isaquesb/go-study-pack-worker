package gookit

import "github.com/gookit/event"

type domEvent struct {
	goKitEvent event.Event
}

func (e *domEvent) Name() string {
	return e.goKitEvent.Name()
}

func (e *domEvent) Data() map[string]interface{} {
	return e.goKitEvent.Data()
}
