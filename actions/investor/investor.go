// Investors only fuctions
package investor

import (
	"github.com/udonetsm/investing/actions/general"
	"github.com/udonetsm/investing/models"
)

type Investor models.Investors
type Startuper models.Startupers

// Перевести деньги стартаперу(на его внутренний счет).
func (from *Investor) Pay(transaction *models.Transaction) {
	if transaction.Err != nil {
		return
	}
	if transaction.Accepted {
		from.Bill.Balance -= transaction.Transaction_sum
	}
}

// Пополнить баланс инвестора(его внутренний счет)
func (investor *Investor) Recieve(transaction *models.Transaction) {
	if transaction.Err != nil {
		return
	}
	investor.Bill.Balance += transaction.Transaction_sum
	// Деньги поступили на счет. Транзакция успешна
	transaction.Success = true
}

// Инвестор запрашивает от стартапера первод. Это задействуется, когда
// инвестор хочет вернуть деньги инвестированные в проект
// В этом случае подтверждать транзакцию должен стартапер
func (investor *Investor) RequestTransaction(transaction *models.Transaction) {
	// Бизнес-логика запроса перевода
	if transaction.Err != nil {
		return
	}
	transaction.Err = general.MakeRequest(transaction)
}

// Инвестор должен подтвердить транзакцию. Это задейстсвуется, когда
// стартапер запросил у инвестора перевод(инвестицию в проект)
func (investor *Investor) AcceptTransaction(transaction *models.Transaction) {
	// Бизнес-логика подтверждения операции
	if transaction.Err != nil {
		return
	}
	// Показать запрос инвестору
	// Получить от него явное согласие на транзакцию.
}

func (i *Investor) Topup(transaction *models.Transaction) {
	// Бизнес-логика пополнения баланса.
}

func (i *Investor) Withdraw(transaction *models.Transaction) {
	// Бизнес-логика снятия денег со счета
}
