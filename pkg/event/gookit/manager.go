package gookit

import (
	"github.com/gookit/event"
	domainEvt "worker/core/domain/event"
	portEvt "worker/core/port/event"
)

type goKitManager struct {
	Manager *event.Manager
}

func NewGoKitManager(Name string, Subscribers []domainEvt.Subscriber) portEvt.Manager {
	manager := event.NewManager(Name)
	for _, sub := range Subscribers {
		manager.AddSubscriber(goKitSubscriber{
			domainSub: sub,
		})
	}
	return goKitManager{
		Manager: manager,
	}
}

func (manager goKitManager) FireEvent(e domainEvt.Event) error {
	err, _ := manager.Manager.Fire(e.Name(), e.Data())
	return err
}
func (manager goKitManager) Fire(Name string, Data map[string]interface{}) error {
	err, _ := manager.Manager.Fire(Name, Data)
	return err
}

func (manager goKitManager) AddSubscriber(domainSub domainEvt.Subscriber) error {
	sub := goKitSubscriber{
		domainSub: domainSub,
	}
	manager.Manager.AddSubscriber(sub)
	return nil
}
