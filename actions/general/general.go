package general

import (
	"github.com/udonetsm/investing/interfaces"
)

// The function Saves any object into any spaces.
// To use it, space object sould implement its method
// Save(any)error
func SaveSomething(saver interfaces.Saver, something any) error {
	return saver.Save(something)
}

func UpdateBalance(updater interfaces.BalanceUpdater, something any) error {
	return updater.BalanceUpdate(something)
}
