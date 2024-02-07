// All transactions
package transactions

import (
	"fmt"
	"strconv"
	"time"

	"github.com/udonetsm/investing/interfaces"
	"github.com/udonetsm/investing/models"
)

func Transaction(from interfaces.Payer, to interfaces.Getter, sum float32) models.Transaction {
	transaction_id, err := from.Pay(to, sum)
	if err != nil {
		return models.Transaction{}
	}
	err = to.Get(transaction_id)
	if err != nil {
		// Сообщение о неудаче
		return models.Transaction{}
	}
	fmt.Println("Transaction OK")
	// ВрЕменная генерация transaction_id, для тестирований
	transaction_id = strconv.Itoa(int(time.Now().UnixNano()))
	// Схранить транзакцию в базе, в таблице Transactions
	return saveTransactios(transaction_id, from, to)
}

// Сохранить в базу информацию о транзакции
func saveTransactios(transaction_id string, from interfaces.Payer, to interfaces.Getter) models.Transaction {
	transaction := models.Transaction{
		Transaction_id: transaction_id,
		From:           from,
		To:             to,
	}
	return transaction
	// Сохранить в базе, продумать формат хранения внутренних транзакций.
}
