package methods

import (
	"github.com/udonetsm/investing/database/funcs"
	"github.com/udonetsm/investing/models"
)

type Database models.Database

// Делаю тип Database Saver-ом
func (d *Database) Save(object any) error {
	funcs.SaveTransaction(object.(models.Transaction))
	return d.Err
}

func (d *Database) BalanceUpdate(object any) error {
	return funcs.UpdateBalances(object.(models.Bills))
}
