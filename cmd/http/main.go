package main

import (
	"fmt"
	"github.com/sarulabs/di"
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
	support.Container = support.Store{App: createApp()}
	defer readMessages()
	service := createsvc.CreateSvc{}
	err := service.Create(time.Now().String())
	if err != nil {
		fmt.Print("error!!!")
		return
	}
}

func readMessages() {
	support.Container.Logger().Info("Main Reading...")
	support.Container.Messenger().Read(func (Msg messenger.Message) {
		support.Container.Logger().Info(fmt.Sprintf("new msg: %s with %s", Msg.Key, Msg.Value))
	}, func (err error) {
		support.Container.Logger().Error(err)
	})
}

func createApp() di.Container {
	builder, _ := di.NewBuilder()
	builder.Add([]di.Def{
		{
			Name: "logger",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return logger.NewZeroLogger(), nil
			},
		},
		{
			Name:  "eventManager",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return pkgEvent.NewGoKitManager("create-svc", []event.Subscriber{
					&createsvc.Subscriber{},
				}), nil
			},
		},
		{
			Name:  "messenger",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return pkgMessenger.NewKafkaMessenger("kafka:9092", "uuids2", "isbSvc"), nil
			},
		},
	}...)
	return builder.Build()
}
