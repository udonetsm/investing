// Database models
package models

import "sync"

const (
	TRANSFER = iota
	WITHDRAW
	TOPUP
)

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
	Balance int
}

type Transaction struct {
	Transaction_id   string //Primary key
	Transaction_type int
	Payer            Payer
	Reciever         Reciever
	Transaction_sum  int
	Err              error
	Success          bool
	Accepted         bool
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

type System struct {
	//Элементы системы
}

type Reciever interface {
	Recieve(*Transaction)
}

type Payer interface {
	Pay(*Transaction)
	Topup(*Transaction)
	Withdraw(*Transaction)
}
