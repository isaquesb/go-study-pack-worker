package main

import (
	"fmt"
	"time"
	"worker/core/domain/event"
	"worker/core/domain/messenger"
	"worker/core/service/createsvc"
	"worker/core/support"
	pkgEvent "worker/pkg/event/gookit"
	"worker/pkg/logger"
	pkgMessenger "worker/pkg/messenger"
)

func main () {
	support.Container = support.Store{
		Logger: logger.NewZeroLogger(),
		EventManager: pkgEvent.NewGoKitManager("create-svc", []event.Subscriber{
			&createsvc.Subscriber{},
		}),
		//Messenger: pkgMessenger.NewKafkaMessenger(os.Getenv("kafkaURL"), os.Getenv("topic"), "createSvc"),
		Messenger: pkgMessenger.NewKafkaMessenger("kafka:9092", "uuids2", "isbSvc"),
	}
	defer readMessages()
	service := createsvc.CreateSvc{}
	err := service.Create(time.Now().String())
	if err != nil {
		fmt.Print("error!!!")
		return
	}
}

func readMessages() {
	support.Container.Logger.Info("Main Reading...")
	support.Container.Messenger.Read(func (Msg messenger.Message) {
		support.Container.Logger.Info(fmt.Sprintf("new msg: %s with %s", Msg.Key, Msg.Value))
	}, func (err error) {
		support.Container.Logger.Error(err)
	})
}
