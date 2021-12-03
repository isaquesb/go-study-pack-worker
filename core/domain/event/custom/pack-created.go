package custom

import "worker/core/domain/event"

type PackCreated struct {
	ID string
}

func (PackCreated) Name() string {
	return event.EvtPackCreated
}

func (c PackCreated) Data() map[string]interface{} {
	payload := make(map[string]interface{})
	payload["id"] = c.ID
	return payload
}
