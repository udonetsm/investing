// Testing of general functions
package tests

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/udonetsm/investing/actors/investor"
	"github.com/udonetsm/investing/actors/startuper"
	"github.com/udonetsm/investing/actors/storage/cache"
	"github.com/udonetsm/investing/actors/storage/database"
	"github.com/udonetsm/investing/actors/system"
	"github.com/udonetsm/investing/general"
	"github.com/udonetsm/investing/models"
)

func TestStartuperToInvestorWithoutError(t *testing.T) {
	sbill := models.Bills{
		Bill_id: 123,
		Balance: 300000,
	}
	suser := models.Users{
		User_id:   98765,
		User_name: "POORMAN",
		Raiting:   12,
		Bill_id:   sbill.Bill_id,
	}
	payer := &startuper.Startuper{
		User: suser,
		Bill: sbill,
	}
	ibill := models.Bills{
		Bill_id: 456,
		Balance: 850000,
	}
	iuser := models.Users{
		User_id:   54321,
		User_name: "RICHMAN",
		Raiting:   30,
		Bill_id:   ibill.Bill_id,
	}
	reciever := &investor.Investor{
		User: iuser,
		Bill: ibill,
	}

	transaction := &models.Transaction{
		Transaction_id:   strconv.Itoa(int(time.Now().UnixNano())),
		Payer:            payer,
		Reciever:         reciever,
		Transaction_sum:  200000,
		Accepted:         true,
		Success:          false,
		Transaction_type: models.TRANSFER,
	}

	system := &system.System{}
	general.RequestTansaction(reciever, transaction)
	if transaction.Err != nil {
		transaction.Err = errors.Join(transaction.Err, fmt.Errorf("Request transaction error"))
		save(transaction)
		return
	}
	general.AcceptTransaction(payer, transaction)
	if !transaction.Accepted {
		transaction.Err = errors.Join(transaction.Err, fmt.Errorf("Denied by payer"))
		save(transaction)
		return
	}
	general.MakeTransaction(system, transaction)
	if transaction.Err != nil {
		// В процессе транзакции произошла ошибка
		// Сохранить в бд. Если не получилось в бд, сохранить в кэш
		save(transaction)
		return
	}
	save(transaction)
}

func save(transaction *models.Transaction) {
	general.SaveSomething(&database.DB, transaction)
	if database.DB.Err != nil {
		fmt.Print("Can't save into the database...")
		general.SaveSomething(&cache.TransactionsCache, transaction)
		if cache.TransactionsCache.Err != nil {
			fmt.Printf("Can't save into the cache...%v %v %v\n",
				transaction.Success, transaction.Transaction_id, transaction.Err)
			fmt.Print("Cache error...")
			return
		} else {
			fmt.Printf("Saved into the cache...%v %v %v\n",
				transaction.Success, transaction.Transaction_id, transaction.Err)
			return
		}
	} else {
		fmt.Println("Saved into the database...\n",
			transaction.Success, transaction.Transaction_id, transaction.Err)
		return
	}
}

func TestTopup(t *testing.T) {
	pbill := models.Bills{
		Bill_id: 123,
		Balance: 300000,
	}
	puser := models.Users{
		User_id:   990,
		Bill_id:   pbill.Bill_id,
		User_name: "TOPUPER",
		Raiting:   30,
	}
	payer := &investor.Investor{
		User: puser,
		Bill: pbill,
	}
	transaction := &models.Transaction{
		Transaction_id:   strconv.Itoa(int(time.Now().UnixNano())),
		Transaction_type: models.TOPUP,
		Payer:            payer,
		Reciever:         payer,
	}
	system := &system.System{}
	general.MakeTransaction(system, transaction)
	if transaction.Err != nil {
		save(transaction)
		return
	}
	save(transaction)
}

func TestWthdraw(t *testing.T) {
	pbill := models.Bills{
		Bill_id: 123,
		Balance: 300000,
	}
	puser := models.Users{
		User_id:   880,
		Bill_id:   pbill.Bill_id,
		User_name: "WITHDRAWER",
		Raiting:   10,
	}
	payer := &startuper.Startuper{
		User: puser,
		Bill: pbill,
	}
	transaction := &models.Transaction{
		Transaction_id:   strconv.Itoa(int(time.Now().UnixNano())),
		Transaction_type: models.WITHDRAW,
		Payer:            payer,
		Reciever:         payer,
	}
	system := &system.System{}
	general.MakeTransaction(system, transaction)
	if transaction.Err != nil {
		save(transaction)
		return
	}
	save(transaction)
}

func TestTransferWithTransactionError(t *testing.T) {
	sbill := models.Bills{
		Bill_id: 374,
		Balance: 3300000,
	}
	suser := models.Users{
		User_id:   87543,
		User_name: "SOMEUSER",
		Raiting:   20,
		Bill_id:   sbill.Bill_id,
	}
	payer := &startuper.Startuper{
		User: suser,
		Bill: sbill,
	}
	ibill := models.Bills{
		Bill_id: 456,
		Balance: 850000,
	}
	iuser := models.Users{
		User_id:   54321,
		User_name: "RICHMAN",
		Raiting:   30,
		Bill_id:   ibill.Bill_id,
	}
	reciever := &investor.Investor{
		User: iuser,
		Bill: ibill,
	}

	transaction := &models.Transaction{
		Transaction_id:   strconv.Itoa(int(time.Now().UnixNano())),
		Payer:            payer,
		Reciever:         reciever,
		Transaction_sum:  200000,
		Accepted:         true,
		Success:          false,
		Transaction_type: models.TRANSFER,
	}

	system := &system.System{}
	general.RequestTansaction(reciever, transaction)
	if transaction.Err != nil {
		transaction.Err = errors.New("Request transaction error")
		return
	}
	general.AcceptTransaction(payer, transaction)
	if !transaction.Accepted {
		transaction.Err = errors.New("Denied by payer")
		save(transaction)
		return
	}
	transaction.Err = errors.New("Simulated error")
	general.MakeTransaction(system, transaction)
	if transaction.Err != nil {
		// В процессе транзакции произошла ошибка
		// Сохранить в бд. Если не получилось в бд, сохранить в кэш
		save(transaction)
		return
	}
	save(transaction)
}

func TestTransferWithTransactionAndDatabaseError(t *testing.T) {
	sbill := models.Bills{
		Bill_id: 374,
		Balance: 3300000,
	}
	suser := models.Users{
		User_id:   87543,
		User_name: "SOMEUSER",
		Raiting:   20,
		Bill_id:   sbill.Bill_id,
	}
	payer := &startuper.Startuper{
		User: suser,
		Bill: sbill,
	}
	ibill := models.Bills{
		Bill_id: 456,
		Balance: 850000,
	}
	iuser := models.Users{
		User_id:   54321,
		User_name: "RICHMAN",
		Raiting:   30,
		Bill_id:   ibill.Bill_id,
	}
	reciever := &investor.Investor{
		User: iuser,
		Bill: ibill,
	}

	transaction := &models.Transaction{
		Transaction_id:   strconv.Itoa(int(time.Now().UnixNano())),
		Payer:            payer,
		Reciever:         reciever,
		Transaction_sum:  200000,
		Accepted:         true,
		Success:          false,
		Transaction_type: models.TRANSFER,
	}

	system := &system.System{}
	general.RequestTansaction(reciever, transaction)
	if transaction.Err != nil {
		transaction.Err = errors.New("Request transaction error")
		return
	}
	general.AcceptTransaction(payer, transaction)
	if !transaction.Accepted {
		transaction.Err = errors.New("Denied by payer")
		save(transaction)
		return
	}
	transaction.Err = errors.New("Simulated error")
	database.DB.Err = errors.New("DATABASE UNAVAILABLE")
	general.MakeTransaction(system, transaction)
	if transaction.Err != nil {
		// В процессе транзакции произошла ошибка
		// Сохранить в бд. Если не получилось в бд, сохранить в кэш
		save(transaction)
		return
	}
	save(transaction)
}

func showCache(m map[string]any) {
	for _, v := range m {
		fmt.Println(
			v.(*models.Transaction).Transaction_id,
			v.(*models.Transaction).Payer,
			v.(*models.Transaction).Reciever,
			v.(*models.Transaction).Transaction_type,
			v.(*models.Transaction).Success,
			v.(*models.Transaction).Accepted,
			v.(*models.Transaction).Err,
			v.(*models.Transaction).Transaction_sum,
		)
	}
}

func TestCache(t *testing.T) {
	showCache(cache.TransactionsCache.Storage)
}
