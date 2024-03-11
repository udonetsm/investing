// Database models
package models

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

const (
	INVESTING = "0"
	RETURNING = "1"
	WITHDRAW  = "2"
	TOPUP     = "3"
)

type Userentry struct {
	Uid      string `gorm:"uid" json:"uid"`
	Baseuser string `gorm:"baseuser" json:"daseuser"`
}
type Transactionentry struct {
	Tid         string `gorm:"tid" json:"tid"`
	Transaction string `gorm:"transaction" json:"transaction"`
}

type User struct {
	// ИНН в качестве uid устанваливается если пользователь прикрепил соответствующие документы и прошел проверку
	// До тех пор в качестве uid используется random uuid
	Uid     string `json:"uid"`
	Raiting int    `json:"raiting,omitempty"`
	Uname   string `json:"uname,omitempty"`
	// Member of contains startup ids where user invested
	Memberof []string `json:"memberof,omitempty"`
	// Founderof contains startup ids which user founded
	Founderof []string `json:"founder,omitempty"`
	Balance   int      `json:"balance,omitempty"`
}

type Startups struct {
	Startup_id string `json:"startupid"`
	Founder_id string `json:"founderid,omitempty"`
	// Members contains ids of all members who is a part of startup_id as investor or founder
	Members []string `json:"members,omitempty"`
	Total   int      `json:"total,omitempty"`
	Raiting int      `json:"raiting,omitempty"`
	Desc    string   `json:"desc,omitempty"`
}

type Transaction struct {
	Tid           string `json:"tid"`
	Type          string `json:"ttype,omitempty"`
	Requested_sum int    `json:"sum,omitempty"`
	DenyReason    int    `json:"denyreason,omitempty"`
	Payer         User   `json:"payer,omitempty"`
	Reciever      User   `json:"reciever,omitempty"`
	Err           error  `json:"-"`
	Success       bool   `json:"success,omitempty"`
	Accepted      bool   `json:"accepted,omitempty"`
}

type Database struct {
	User    string
	Pass    string
	Host    string
	Port    int
	Dbname  string
	SSLMode string
	Err     error
}

type Cache struct {
	Mux     *sync.Mutex
	Storage map[any]any
	Err     error
}

type Redis struct {
	Client *redis.Client
	Key    string
	Value  string
	Error  error
}

type Stack struct {
	Mess  []*Messages
	Count int
}

type Messages struct {
	Message   any   `json:"message,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
	Counter   int   `json:"counter,omitempty"`
	Error     any   `json:"error,omitempty"`
}

type Reply struct {
	Content any `json:"content,omitempty"`
	Error   any `json:"error,omitempty"`
}

type Stub struct{}
