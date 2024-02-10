package cache

import (
	"sync"

	"github.com/udonetsm/investing/models"
)

type TransactionCache models.Cache

var TC TransactionCache

func init() {
	TC = NewCache()
}

func NewCache() TransactionCache {
	return TransactionCache{
		Mux:     new(sync.Mutex),
		Storage: NewStorage().(map[string]any),
	}
}

func NewStorage() any {
	return make(map[string]any)
}

// Private method
func insertTransaction(c *TransactionCache, t any) error {
	if c.Err != nil {
		return c.Err
	}
	c.Storage[t.(models.Transaction).Transaction_id] = t.(models.Transaction)
	return nil
}

// Делаю TransactionCache Saver-ом
func (c *TransactionCache) Save(t any) (err error) {
	return insertTransaction(c, t)
}
