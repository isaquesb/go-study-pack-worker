package gookit

import domainEvt "worker/core/domain/event"

type goKitSubscriber struct {
	domainSub domainEvt.Subscriber
}

func (sub goKitSubscriber) SubscribedEvents() map[string]interface{} {
	from := sub.domainSub.Events()
	to := make(map[string]interface{})
	for index, lsn := range from {
		to[index] = &goKitListener{
			domainListener: lsn,
		}
	}
	return to
}
