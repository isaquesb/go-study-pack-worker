package event

import (
	"fmt"
	"worker/core/domain/event"
	"worker/core/domain/messenger"
	"worker/core/support"
)

type CreatedListener struct {
}

func (l *CreatedListener) Handle(e event.Event) error {
	data := e.Data()
	support.Container.Logger().Info(fmt.Sprintf("pack created with id: %s", data["id"]))
	msg := messenger.Message{
		//Topic: os.Getenv("topic"),
		Topic: "uuids2",
		Key: []byte("pack.created"),
		Value: []byte(fmt.Sprintf("%s", data["id"])),
	}
	support.Container.Messenger().Write(msg)
	return nil
}
