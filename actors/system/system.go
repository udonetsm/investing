// All transactions
package system

import (
	"github.com/udonetsm/investing/models"
)

type System models.System

// Первод между пользователями. Обязательно от инвестора стартаперу или обратно.
// Переводы между инветорами запрещены.
// Переводы между стартаперами запрещены.
func (s *System) TransferTransaction(transaction *models.Transaction) {
	if transaction.Err != nil {
		return
	}
	// Делаем оплату со счета отправителя
	transaction.Payer.Pay(transaction)
	if transaction.Err != nil {
		return
	}
	// Например недостаточно средств
	if transaction.Err != nil {
		return
	}
	transaction.Reciever.Recieve(transaction)
}

// Пополнение счета
func (s *System) TopupTransaction(transaction *models.Transaction) {
	transaction.Payer.Topup(transaction)
	// ...
}

// Снятие средств со счета
func (s *System) WithdrawTransaction(transaction *models.Transaction) {
	transaction.Payer.Withdraw(transaction)
	// ...
}
