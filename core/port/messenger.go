package port

import "worker/core/domain/messenger"

type Messenger interface {
	Write(Message messenger.Message) error
	Read(Callback func (message messenger.Message), ErrCallback func (err error))
}
