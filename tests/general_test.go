// Testing of general functions
package tests

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/udonetsm/investing/actions/general"
	"github.com/udonetsm/investing/actions/startuper"
	"github.com/udonetsm/investing/actions/transactions"
	"github.com/udonetsm/investing/cache"
	"github.com/udonetsm/investing/database"
	"github.com/udonetsm/investing/models"
)

func TestStartuperToInvestorWithTransferError(t *testing.T) {
	suser := models.Users{
		User_id:   98765,
		User_name: "POORMAN",
		Raiting:   12,
	}
	sbill := models.Bills{
		Bill_id: 123,
		Balance: 300000,
	}
	payer := startuper.Startuper{
		User: suser,
		Bill: sbill,
	}
	iuser := models.Users{
		User_id:   54321,
		User_name: "RICHMAN",
		Raiting:   30,
	}
	ibill := models.Bills{Bill_id: 456, Balance: 850000}
	reciever := startuper.Investor{
		User: iuser,
		Bill: ibill,
	}
	transaction := models.Transaction{
		Transaction_id: strconv.Itoa(int(time.Now().UnixNano())),
		Payer:          payer,
		Reciever:       reciever,
		Sum:            200000,
		Err:            errors.New("TRANSACTION ERROR IMITATE"),
	}
	transaction = transactions.TransferMoney(transaction)
	if transaction.Err != nil {
		fmt.Printf("Transaction %s losst...", transaction.Transaction_id)
	}
	err := general.SaveSomething(&database.DB, transaction)
	if err != nil {
		err := general.SaveSomething(&cache.TC, transaction)
		if err != nil {
			fmt.Print("Не получилось сохранить в базу и кэш...")
		} else {
			fmt.Print("Saved into the cache...")
		}
	} else {
		fmt.Print("Saved into the database...")
	}
	if transaction.Err == nil {
		database.DB.Err = general.UpdateBalance(&database.DB, ibill)
		database.DB.Err = general.UpdateBalance(&database.DB, sbill)
		fmt.Printf("Transaction %s OK...", transaction.Transaction_id)
	}
	fmt.Println("\n", transaction)
}

func TestStartuperToInvestorWithDatabaseError(t *testing.T) {
	suser := models.Users{
		User_id:   98765,
		User_name: "POORMAN",
		Raiting:   12,
	}
	sbill := models.Bills{
		Bill_id: 123,
		Balance: 300000,
	}
	payer := startuper.Startuper{
		User: suser,
		Bill: sbill,
	}
	iuser := models.Users{
		User_id:   54321,
		User_name: "RICHMAN",
		Raiting:   30,
	}
	ibill := models.Bills{Bill_id: 456, Balance: 850000}
	reciever := startuper.Investor{
		User: iuser,
		Bill: ibill,
	}
	transaction := models.Transaction{
		Transaction_id: strconv.Itoa(int(time.Now().UnixNano())),
		Payer:          payer,
		Reciever:       reciever,
		Sum:            200000,
	}

	transaction = transactions.TransferMoney(transaction)
	if transaction.Err != nil {
		fmt.Printf("Transaction %s losst...", transaction.Transaction_id)
	}
	err := general.SaveSomething(&database.DB, transaction)
	if err != nil {
		err := general.SaveSomething(&cache.TC, transaction)
		if err != nil {
			fmt.Print("Не получилось сохранить в базу и кэш...")
		} else {
			fmt.Print("Saved into the cache...")
		}
	} else {
		fmt.Print("Saved into the database...")
	}
	if transaction.Err == nil {
		database.DB.Err = general.UpdateBalance(&database.DB, ibill)
		database.DB.Err = general.UpdateBalance(&database.DB, sbill)
		fmt.Printf("Transaction %s OK...", transaction.Transaction_id)
	}
	fmt.Println("\n", transaction)
}

func TestStartuperToInvestorWithCacheError(t *testing.T) {
	suser := models.Users{
		User_id:   98765,
		User_name: "POORMAN",
		Raiting:   12,
	}
	sbill := models.Bills{
		Bill_id: 123,
		Balance: 300000,
	}
	payer := startuper.Startuper{
		User: suser,
		Bill: sbill,
	}
	iuser := models.Users{
		User_id:   54321,
		User_name: "RICHMAN",
		Raiting:   30,
	}
	ibill := models.Bills{Bill_id: 456, Balance: 850000}
	reciever := startuper.Investor{
		User: iuser,
		Bill: ibill,
	}
	transaction := models.Transaction{
		Transaction_id: strconv.Itoa(int(time.Now().UnixNano())),
		Payer:          payer,
		Reciever:       reciever,
		Sum:            200000,
	}
	// Immitate cache error
	cache.TC.Err = errors.New("IMMITATE CACHE ERROR")
	transaction = transactions.TransferMoney(transaction)
	if transaction.Err != nil {
		fmt.Printf("Transaction %s losst...", transaction.Transaction_id)
	}
	err := general.SaveSomething(&database.DB, transaction)
	if err != nil {
		err := general.SaveSomething(&cache.TC, transaction)
		if err != nil {
			fmt.Print("Не получилось сохранить в базу и кэш...")
		} else {
			fmt.Print("Saved into the cache...")
		}
	} else {
		fmt.Print("Saved into the database...")
	}
	if transaction.Err == nil {
		database.DB.Err = general.UpdateBalance(&database.DB, ibill)
		database.DB.Err = general.UpdateBalance(&database.DB, sbill)
		fmt.Printf("Transaction %s OK...", transaction.Transaction_id)
	}
	fmt.Println("\n", transaction)
}

func TestStartuperToInvestorWithoutErrors(t *testing.T) {
	suser := models.Users{
		User_id:   98765,
		User_name: "POORMAN",
		Raiting:   12,
	}
	sbill := models.Bills{
		Bill_id: 123,
		Balance: 300000,
	}
	payer := startuper.Startuper{
		User: suser,
		Bill: sbill,
	}
	iuser := models.Users{
		User_id:   54321,
		User_name: "RICHMAN",
		Raiting:   30,
	}
	ibill := models.Bills{Bill_id: 456, Balance: 850000}
	reciever := startuper.Investor{
		User: iuser,
		Bill: ibill,
	}
	transaction := models.Transaction{
		Transaction_id: strconv.Itoa(int(time.Now().UnixNano())),
		Payer:          payer,
		Reciever:       reciever,
		Sum:            200000,
	}
	transaction = transactions.TransferMoney(transaction)
	if transaction.Err != nil {
		log.Printf("Transaction %s losst...", transaction.Transaction_id)
	}
	err := general.SaveSomething(&database.DB, transaction)
	if err != nil {
		err := general.SaveSomething(&cache.TC, transaction)
		if err != nil {
			fmt.Print("Не получилось сохранить в базу и кэш...")
		} else {
			fmt.Print("Saved into the cache...")
		}
	} else {
		fmt.Print("Saved into the database...")
	}
	if transaction.Err == nil {
		database.DB.Err = general.UpdateBalance(&database.DB, ibill)
		database.DB.Err = general.UpdateBalance(&database.DB, sbill)
		fmt.Printf("Transaction %s OK...", transaction.Transaction_id)
	}
	fmt.Println("\n", transaction)
}

func showCache(m map[string]any) {
	for _, v := range m {
		log.Println(v)
	}
}

func TestCache(t *testing.T) {
	showCache(cache.TC.Storage)
}
