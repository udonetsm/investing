// Investors only fuctions
package investor

import (
	"github.com/udonetsm/investing/models"
)

type Investor models.Investors
type Startuper models.Startupers

// Investor может перевести деньги стартаперу.
func (from Investor) DoTransaction(transaction models.Transaction) models.Transaction {
	// err:= проверяем баланс Payer-а
	if transaction.Err != nil {
		// Ошибку проверки баланса можно кастомизировать здесь
		return transaction
	}
	// Если средств достаточно и ошибки нет то делаем транзакцию
	from.Bill.Balance -= transaction.Sum
	transaction.Payer = from
	return transaction
}

// Эта функция должна выполняться пока
// на счете получателя не появится транзакция
// Либо через done канал,
// либо через wiatGroup,
// либо через nats streaming server
func (to Startuper) Recieve(transaction models.Transaction) models.Transaction {
	var err error
	// err = Проверка пришли ли средства получателю
	if err != nil {
		// Деньги не пришли
		transaction.Err = err
		return transaction
	}
	//Деньги пришли
	to.Bill.Balance += transaction.Sum
	transaction.Reciever = to
	return transaction
}
