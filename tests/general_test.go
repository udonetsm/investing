// Testing of general functions
package tests

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/udonetsm/investing/actions/investor"
	"github.com/udonetsm/investing/actions/startuper"
	"github.com/udonetsm/investing/actions/transactions"
	"github.com/udonetsm/investing/cache"
	"github.com/udonetsm/investing/database/calls"
	"github.com/udonetsm/investing/models"
	"gorm.io/gorm"
)

var Cache cache.TransactionCache

func init() {
	Cache = *cache.NewCache()
}

// Test if Transaction isn't success and database isn't available
func TestInvestorTransfersMoneyToStartuper(t *testing.T) {
	fmt.Println()
	ibill := models.Bills{Bill_id: 123456}
	sbill := models.Bills{Bill_id: 456789}
	payer := investor.Investor{User: models.Users{
		User_name: "RICHMAN",
		User_id:   987654321,
	},
		Bill: ibill,
	}
	getter := investor.Startuper{User: models.Users{
		User_name: "POORMAN",
		User_id:   1234567890,
	},
		Bill: sbill}
	transaction := models.Transaction{
		// Временная генерация transaction_id.
		// В дальнейшем можно брать id из банковской транзакции,
		// а здесь оставлять поле пустым
		Transaction_id: strconv.Itoa(int(time.Now().UnixNano())),
		Sum:            100000,
		Payer:          payer,
		Reciever:       getter,
		Err:            nil,
	}
	transaction = transactions.TransferMoney(transaction)
	transaction.Err = errors.New("ERROR WHILE TRANSACTION AND DATABASE ISN'T AVAILABLE")
	if transaction.Err != nil {
		err := calls.SaveTransactios(transaction)
		err = gorm.ErrInvalidDB
		if err != nil {
			Cache.Insert(transaction)
			log.Println(err, "Saved into the cache")
			return
		}
	} else {
		transaction.Success = true
		err := calls.SaveTransactios(transaction)
		if err != nil {
			Cache.Insert(transaction)
			log.Println(err, "Saved into the cache")
			return
		}
	}
	fmt.Println(transaction)
}

// Test if transaction isn't success and db is available
func TestStartuperTransfersMoneyToInvestor(t *testing.T) {
	fmt.Println()
	ibill := models.Bills{Bill_id: 876591}
	sbill := models.Bills{Bill_id: 123409}
	payer := startuper.Startuper{User: models.Users{
		User_name: "POORMAN",
		User_id:   692683000,
	},
		Bill: sbill,
	}
	getter := startuper.Investor{User: models.Users{
		User_name: "RICHMAN",
		User_id:   859748000,
	},
		Bill: ibill,
	}
	transaction := models.Transaction{
		// Временная генерация transaction_id.
		// В дальнейшем можно брать id из банковской транзакции,
		// а здесь оставлять поле пустым
		Transaction_id: strconv.Itoa(int(time.Now().UnixNano())),
		Payer:          payer,
		Reciever:       getter,
		Sum:            200000,
		Err:            nil,
	}
	transaction = transactions.TransferMoney(transaction)
	transaction.Err = errors.New("ERROR WHILE TRANSACTION AND DATABASE IS AVAILABLE")
	if transaction.Err != nil {
		err := calls.SaveTransactios(transaction)
		if err != nil {
			Cache.Insert(transaction)
			log.Println(err, "Saved into the cache")
			return
		}
	} else {
		transaction.Success = true
		err := calls.SaveTransactios(transaction)
		if err != nil {
			Cache.Insert(transaction)
			log.Println(err, "Saved into the cache")
			return
		}
	}
	fmt.Println(transaction)
}

// Test if transaction is success and database isn't available
func TestDatabaseDatabaseUnavailable(t *testing.T) {
	fmt.Println()
	ibill := models.Bills{Bill_id: 654321}
	sbill := models.Bills{Bill_id: 567890}
	payer := startuper.Startuper{User: models.Users{
		User_name: "POORMAN",
		User_id:   996438856734,
	},
		Bill: sbill,
	}
	getter := startuper.Investor{User: models.Users{
		User_name: "RICHMAN",
		User_id:   995623476283,
	},
		Bill: ibill,
	}
	transaction := models.Transaction{
		// Временная генерация transaction_id.
		// В дальнейшем можно брать id из банковской транзакции,
		// а здесь оставлять поле пустым
		Transaction_id: strconv.Itoa(int(time.Now().UnixNano())),
		Payer:          payer,
		Reciever:       getter,
		Sum:            200000,
		Err:            nil,
	}
	transaction = transactions.TransferMoney(transaction)
	if transaction.Err != nil {
		err := calls.SaveTransactios(transaction)
		if err != nil {
			Cache.Insert(transaction)
			log.Println(err, "Saved into the cache")
			return
		}
		return
	} else {
		transaction.Success = true
		err := calls.SaveTransactios(transaction)
		err = gorm.ErrInvalidDB
		if err != nil {
			Cache.Insert(transaction)
			log.Println(err, "Saved into the cache")
			return
		}
	}
	log.Println(transaction)
}

func TestAllErrorsAreNil(t *testing.T) {
	fmt.Println()
	ibill := models.Bills{Bill_id: 654321}
	sbill := models.Bills{Bill_id: 567890}
	payer := startuper.Startuper{User: models.Users{
		User_name: "POORMAN",
		User_id:   996438856734,
	},
		Bill: sbill,
	}
	getter := startuper.Investor{User: models.Users{
		User_name: "RICHMAN",
		User_id:   995623476283,
	},
		Bill: ibill,
	}
	transaction := models.Transaction{
		// Временная генерация transaction_id.
		// В дальнейшем можно брать id из банковской транзакции,
		// а здесь оставлять поле пустым
		Transaction_id: strconv.Itoa(int(time.Now().UnixNano())),
		Payer:          payer,
		Reciever:       getter,
		Sum:            200000,
		Err:            nil,
	}
	transaction = transactions.TransferMoney(transaction)
	if transaction.Err != nil {
		err := calls.SaveTransactios(transaction)
		if err != nil {
			Cache.Insert(transaction)
			log.Println(err, "Saved into the cache")
			return
		}
		return
	} else {
		transaction.Success = true
		err := calls.SaveTransactios(transaction)
		if err != nil {
			Cache.Insert(transaction)
			log.Println(err, "Saved into the cache")
			return
		}
	}
	log.Println(transaction)
}

func TestCache(t *testing.T) {
	for _, v := range Cache.Storage {
		log.Println(v)
		fmt.Println()
	}
}
