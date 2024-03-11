// Interfaces
// Several else are in models. Not here to avoid cycle import
package interfaces

import (
	"github.com/udonetsm/investing/actors/stack"
	"github.com/udonetsm/investing/models"
)

type SomethingSaver interface {
	Save(any) *models.Reply
}

type SomethingGeter interface {
	Get(any) *models.Reply
}

type SomethingUpdater interface {
	Update(string, string, any, any) *models.Reply
}

type SomethingSaveGetUpdater interface {
	SomethingSaver
	SomethingGeter
	SomethingUpdater
}

type LastSeenGeter interface {
	GetLastSeen(string) int64
}

type LastSeenSeter interface {
	SetLastSeen(string, int64)
}

type LastSeenSetGeter interface {
	LastSeenGeter
	LastSeenSeter
}

type Subscriber interface {
	Subscribe(string) *stack.Stack
}

type Publisher interface {
	Publish(string)
}

type SubscribePublisher interface {
	Subscriber
	Publisher
}
