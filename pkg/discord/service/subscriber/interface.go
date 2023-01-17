package subscriber

import "github.com/dwarvesf/fortress-discord/pkg/model"

type SubscriberServicer interface {
	GetList() ([]*model.Subscriber, error)
}
