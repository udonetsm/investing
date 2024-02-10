// Startuper only functions

package startuper

import (
	"github.com/udonetsm/investing/models"
)

type Investor models.Investors
type Startuper models.Startupers

// Startuper может вернуть деньги инветору.
func (from Startuper) DoTransaction(transaction models.Transaction) models.Transaction {
	// err:= Проверяем баланс Payer-а
	if transaction.Err != nil {
		// Ошибку проверки баланса можно кастомизировать здесь
		return transaction
	}
	// Если средств достаточно и ошибки нет то делаем транзакцию
	from.Bill.Balance -= transaction.Sum
	transaction.Payer = from
	return transaction
}

// Проверяет, получил ли инветор деньги от стартапера.
// Можно, через nats, прикрутить подписку на обновление баланса
func (to Investor) Recieve(transaction models.Transaction) models.Transaction {
	var err error
	// err = Проверка транзакции(деньги поступили?).
	if err != nil {
		// Деньги не пришли
		transaction.Err = err
		return transaction
	}
	// Деньги пришли
	to.Bill.Balance += transaction.Sum
	transaction.Reciever = to
	return transaction
}
