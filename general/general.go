package general

import (
	"github.com/udonetsm/investing/actors/stack"
	"github.com/udonetsm/investing/interfaces"
	"github.com/udonetsm/investing/models"
)

// Saves not transactions in any sql-databases
func Save(saver interfaces.SomethingSaveGetUpdater, something any) *models.Reply {
	return saver.Save(something)
}

func Get(geter interfaces.SomethingSaveGetUpdater, something any) *models.Reply {
	return geter.Get(something)
}

func Update(updater interfaces.SomethingSaveGetUpdater, targetKeyStr, targetValueStr string, targetModel, newEntry any) *models.Reply {
	return updater.Update(targetKeyStr, targetValueStr, targetModel, newEntry)
}

func GetLastSeen(setgeter interfaces.LastSeenSetGeter, uid string) int64 {
	return setgeter.GetLastSeen(uid)
}

func SetLastSeen(setgeter interfaces.LastSeenSetGeter, uid string, timestamp int64) {
	setgeter.SetLastSeen(uid, timestamp)
}

func Subcribe(subscriber interfaces.SubscribePublisher, uid string) *stack.Stack {
	return subscriber.Subscribe(uid)
}

func Publish(publisher interfaces.SubscribePublisher, targetUID string) {
	publisher.Publish(targetUID)
}
