package cache

import (
	"sync"

	"github.com/udonetsm/investing/models"
)

type transactionCache models.Cache

var TransactionsCache transactionCache

func CreateTransactionCache() {
	TransactionsCache = newTransactionCache()
}

func newTransactionCache() transactionCache {
	return transactionCache{
		Mux:     new(sync.Mutex),
		Storage: NewStorage(),
	}
}

// Private method
func insertTransaction(c *transactionCache, transaction *models.Transaction) error {
	if c.Err != nil {
		return c.Err
	}
	c.Storage[transaction.Tid] = transaction
	return nil
}

// Делаю TransactionCache Saver-ом
func (transactionCache *transactionCache) Save(transaction *models.Transaction) {
	if transactionCache.Err != nil {
		return
	}
	insertTransaction(transactionCache, transaction)
}
