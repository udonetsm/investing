package cache

import (
	"sync"

	"github.com/udonetsm/investing/models"
)

type TransactionCache struct {
	Mux     *sync.Mutex
	Storage map[string]models.Transaction
}

func NewCache() *TransactionCache {
	return &TransactionCache{
		Mux:     new(sync.Mutex),
		Storage: NewStorage(),
	}
}

func NewStorage() map[string]models.Transaction {
	return make(map[string]models.Transaction)
}

type Inserter interface {
	Insert(string, models.Transaction)
}

func (t *TransactionCache) Insert(transaction models.Transaction) {
	t.Mux.Lock()
	defer t.Mux.Unlock()
	if _, ok := t.Storage[transaction.Transaction_id]; !ok {
		t.Storage[transaction.Transaction_id] = transaction
	} else {
		return
	}
}
