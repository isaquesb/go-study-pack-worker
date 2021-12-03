package createsvc

import (
	"worker/core/domain/event/custom"
	"worker/core/support"
)

type CreateSvc struct {
}

func (r CreateSvc) Create(id string) error {
	e := custom.PackCreated{ID: id}
	return support.Container.EventManager().FireEvent(e)
}
