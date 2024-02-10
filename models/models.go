// Database models
package models

import "sync"

type Startupers struct {
	User Users
	Bill Bills
}

type Investors struct {
	User Users
	Bill Bills
}

type Users struct {
	User_id   int
	Raiting   int
	User_name string
	Bill_id   int //Foreign key references Bills(Bill_id)
}

type Startups struct {
	Startup_id   int //Primary key
	Startuper_id int //Foreign key references Startupers.Startuper_id
	Total        int
	Raiting      int
	Startup_name string
	Members      []Investors
}

type Bills struct {
	Bill_id int //Primary key
	Balance float32
}

type Transaction struct {
	Transaction_id string //Primary key
	Payer          Payer
	Reciever       Reciever
	Sum            float32
	Err            error
	Success        bool
}

type Database struct {
	User    string
	Pass    string
	Host    string
	Port    int
	SSLMode bool
	Err     error
}

type Cache struct {
	Mux     *sync.Mutex
	Storage map[string]any
	Err     error
}

type Reciever interface {
	Recieve(Transaction) Transaction
}

type Payer interface {
	DoTransaction(Transaction) Transaction
}
