package general

import (
	"github.com/udonetsm/investing/interfaces"
)

func SaveSomething(saver interfaces.Saver, something any) error {
	return saver.Save(something)
}

func UpdateBalance(updater interfaces.BalanceUpdater, something any) error {
	return updater.BalanceUpdate(something)
}
