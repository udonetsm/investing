// Database models
package models

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
	User_name string
}

type Startups struct {
	Startup_id   int //Primary key
	Startuper_id int //Foreign key references Startupers.Startuper_id
	Total        int
	Startup_name string
	Members      Investors //Foreign key references Investors.Investor_id
}

type Bills struct {
	Bill_id int //Primary key
}

type Transaction struct {
	Transaction_id string //Primary key
	Payer          Payer
	Reciever       Reciever
	Sum            float32
	Err            error
	Success        bool
}

type Reciever interface {
	Recieve(Transaction) Transaction
}

type Payer interface {
	DoTransaction(Transaction) Transaction
}
