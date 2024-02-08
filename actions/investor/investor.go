// Investors only fuctions
package investor

import (
	"github.com/udonetsm/investing/models"
)

type Investor models.Investors
type Startuper models.Startupers

// Investor может перевести деньги стартаперу.
func (from Investor) DoTransaction(transaction models.Transaction) models.Transaction {
	var err error
	//err = Запросить баланс банковского счета отправителя
	// и зарезервировать нужную сумму
	// Если возникла какая-либо ошибка, транзакция не осуществляется
	// и ошибка возвращается "наверх".
	// Возможно не хватает средств или банк отклонил соединение
	if err != nil {
		// Ошибку проверки баланса можно кастомизировать здесь
		transaction.Err = err
		return transaction
	}
	// Если средств достаточно и ошибки нет то делаем транзакцию
	// err:= Запрос внутренней транзакции
	if err != nil {
		transaction.Err = err
	}
	// Транзакция успешна. Раппортуем и
	// записываем транзакцию в базу, в таблицу Transactions
	return transaction
}

func (to Startuper) Recieve(transaction models.Transaction) models.Transaction {
	var err error
	// err = Проверка пришли ли средства получателю
	if err != nil {
		// Деньги не пришли
		transaction.Err = err
		return transaction
	}
	//Деньги пришли
	return transaction
}
