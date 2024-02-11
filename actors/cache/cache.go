package cache

import (
	"sync"

	"github.com/udonetsm/investing/models"
)

type TransactionCache models.Cache

var TransactionsCache TransactionCache

func init() {
	TransactionsCache = NewTransactionCache()
}

func NewTransactionCache() TransactionCache {
	return TransactionCache{
		Mux:     new(sync.Mutex),
		Storage: NewStorage().(map[string]any),
	}
}

func NewStorage() any {
	return make(map[string]any)
}

// Private method
func insertTransaction(c *TransactionCache, transaction *models.Transaction) error {
	if c.Err != nil {
		return c.Err
	}
	c.Storage[transaction.Transaction_id] = transaction
	return nil
}

// Делаю TransactionCache Saver-ом
func (transactionCache *TransactionCache) Save(transaction *models.Transaction) {
	if transactionCache.Err != nil {
		return
	}
	insertTransaction(transactionCache, transaction)
}
