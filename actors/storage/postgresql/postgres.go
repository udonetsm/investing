package postgresql

import (
	"fmt"

	"github.com/udonetsm/investing/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database models.Database

func (d *Database) dsn() string {
	dsn := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Pass, d.Dbname, d.SSLMode)
	return dsn
}

func (d *Database) connect() *gorm.DB {
	dsn := d.dsn()
	con, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		d.Err = err
		return &gorm.DB{}
	}
	return con
}

func (d *Database) Save(something any) *models.Reply {
	db := d.connect()
	if d.Err != nil {
		return &models.Reply{Error: d.Err.Error()}
	}
	tx := db.Create(something)
	if tx.Error != nil {
		return &models.Reply{Error: tx.Error.Error()}
	}
	return &models.Reply{Content: "SAVE PASSED"}
}

func (d *Database) Get(something any) *models.Reply {
	db := d.connect()
	if d.Err != nil {
		return nil
	}
	tx := db.First(something)
	if tx.Error != nil {
		return &models.Reply{Error: tx.Error.Error()}
	}
	return &models.Reply{Content: something}
}

// Example: d.Update("id=?", "999", models.Userentry{}, models.Userentry{User_id: 999, User: models.User{...}})
func (d *Database) Update(whereKye, whereValue string, targetModel, something any) *models.Reply /* add returning new value */ {
	db := d.connect()
	if d.Err != nil {
		return &models.Reply{Error: d.Err.Error()}
	}
	tx := db.Model(targetModel).Where(whereKye, whereValue).Updates(something)
	if tx.Error != nil {
		return &models.Reply{Error: tx.Error.Error()}
	}
	return &models.Reply{Content: "OK"}
}
