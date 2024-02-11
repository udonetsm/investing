package database

import "github.com/udonetsm/investing/database/methods"

var DB methods.Database

func init() {
	DB = NewDatabase()
}

func NewDatabase() methods.Database {
	return methods.Database{
		User:    "Test2",
		Pass:    "testpass",
		Host:    "localhost",
		Port:    5432,
		SSLMode: false,
	}
}
