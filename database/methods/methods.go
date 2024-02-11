package methods

import (
	"github.com/udonetsm/investing/database/funcs"
	"github.com/udonetsm/investing/models"
)

type Database models.Database

// Делаю тип Database Saver-ом
func (d *Database) Save(transaction *models.Transaction) {
	if d.Err != nil {
		return
	}
	funcs.SaveTransaction(transaction)
}
