package funcs

import (
	"fmt"

	"github.com/udonetsm/investing/models"
)

func SaveTransaction(transaction models.Transaction) (err error) {
	// Здесь нужно делать запрос на сохранение транзакции в бд
	if transaction.Err != nil {
		return
	}
	return
}

func UpdateBalances(bill models.Bills) (err error) {
	// Здесь нужно делать запрос на обновление баланса пользователя в БД
	fmt.Print("Balace updated...")
	return
}
