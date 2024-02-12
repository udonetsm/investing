// Database models
package models

import "sync"

const (
	TRANSFER = "Transfer"
	WITHDRAW = "Withdraw"
	TOPUP    = "Topup"
)

type Startupers struct {
	User       Users
	Bill       Bills
	Founder_of []Startups
}

type Investors struct {
	User      Users
	Bill      Bills
	Member_of []Startups
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
	Members      []BaseUser
}

type Bills struct {
	Bill_id int //Primary key
	Balance int
}

type Transaction struct {
	Transaction_id   string //Primary key
	Transaction_type string
	Payer            BaseUser
	Reciever         BaseUser
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
	Domain string
	Bill   Bills
}

type Reciever interface {
	Recieve(*Transaction)
}

type Payer interface {
	Pay(*Transaction)
}

type Topuper interface {
	Topup(*Transaction)
}

type Withdrawer interface {
	Withdraw(*Transaction)
}

type Requester interface {
	RequestTransaction(*Transaction)
}

type Accepter interface {
	AcceptTransaction(*Transaction)
}

type BaseUser interface {
	Payer
	Reciever
	Topuper
	Withdrawer
	Accepter
	Requester
}
