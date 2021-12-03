package event

import (
	"worker/core/domain/event"
)

type Manager interface {
	Fire(Name string, Data map[string]interface{}) error
	FireEvent(e event.Event) error
}
