package createsvc

import (
	"worker/core/domain/event"
	serviceEvent "worker/core/service/createsvc/event"
)

type Subscriber struct {
}

func (s *Subscriber) Events() map[string]event.Listener {
	return map[string]event.Listener {
		event.EvtPackCreated: &serviceEvent.CreatedListener{},
	}
}
