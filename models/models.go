// Database models
package models

import "github.com/udonetsm/investing/interfaces"

type DBEntries struct {
	User Users
	Bill Bills
}

type Users struct {
	User_id   int
	User_name string
	Role      bool //true - investor, false - startuper
}

type Startups struct {
	Startup_id   int //Primary key
	Startuper_id int //Foreign key references Startupers.Startuper_id
	Total        int
	Startup_name string
	Members      DBEntries
}

type Bills struct {
	Bill_id int //Primary key
}

type Transaction struct {
	Transaction_id string //Primary key
	From           interfaces.Payer
	To             interfaces.Getter
}
