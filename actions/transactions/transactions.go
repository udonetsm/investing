// All transactions
package transactions

import (
	"github.com/udonetsm/investing/models"
)

func TransferMoney(transaction models.Transaction) models.Transaction {
	// Делаем оплату со счета отправителя
	transaction = transaction.Payer.DoTransaction(transaction)
	// Если в процессе перевода возникла проблема,
	// то проверять баланс получателя незачем
	// Просто возвращаем транзакцию с ошибкой
	if transaction.Err != nil {
		return transaction
	}
	//Проверяем что получаеть получил необходимую сумму
	// (нужная транзакция есть на счете получателя)
	transaction = transaction.Reciever.Recieve(transaction)
	return transaction
}
