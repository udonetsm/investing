// Investors only fuctions
package investor

import (
	"github.com/udonetsm/investing/models"
)

type Investor models.Investors
type Startuper models.Startupers

// Investor может перевести деньги стартаперу.
func (from Investor) DoTransaction(transaction models.Transaction) models.Transaction {
	// Если возникла какая-либо ошибка, транзакция не осуществляется
	// и ошибка возвращается "наверх". Возможно не хватает средств или банк отклонил соединение
	if transaction.Err != nil {
		// Ошибку проверки баланса можно кастомизировать здесь
		return models.Transaction{Err: transaction.Err}
	}
	// Если средств достаточно и ошибки нет то запрашиваем у банка транзакцию.
	if transaction.Err != nil {
		// Неудача во время перевода
		return models.Transaction{Err: transaction.Err}
	}
	// Если все ок, то делаем лог и сообщение об успехе
	// fmt.Printf("%v send %.2f to %v...", s, sum, to)
	// И записываем транзакцию в базу, в таблицу Transactions
	return transaction
}

func (to Startuper) Recieve(transaction models.Transaction) models.Transaction {
	if transaction.Err != nil {
		return models.Transaction{Err: transaction.Err}
	}
	return transaction
}
