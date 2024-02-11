// Startuper only functions

package startuper

import (
	"github.com/udonetsm/investing/actions/general"
	"github.com/udonetsm/investing/models"
)

type Investor models.Investors
type Startuper models.Startupers

// Перевести деньги инвестору(на его внутренний счет)
func (s *Startuper) Pay(transaction *models.Transaction) {
	if transaction.Err != nil {
		return
	}
	if transaction.Accepted {
		s.Bill.Balance -= transaction.Transaction_sum
	}
}

// Пополнить баланс стартапера(его внутренний счет)
func (s *Startuper) Recieve(transaction *models.Transaction) {
	if transaction.Err != nil {
		return
	}
	s.Bill.Balance += transaction.Transaction_sum
	// Деньги поступили на счет. Транзакция успешна
	transaction.Success = true
}

// Запросить у инветора транзакцию.
// Транзикия должна собраться на основе того
// какие значения будут введены стартапером
// в пользовательские формы на сайте или в приложении.
// Доступно несколько видов транзакций
// См. transactioner/transactioner.go или interfaces/interfaces.go:Transactioner
func (s *Startuper) RequestTransaction(transaction *models.Transaction) {
	if transaction.Err != nil {
		return
	}
	transaction.Err = general.MakeRequest(transaction)
}

// Перед тем как совершать транзакцию, нужно получить подтверждение от стратапера
// Эта функция задействуется, когда инвестор запрашивает у стартапера внутренний первод
// Инветсор создает транзакцию, указывает конкретного стартапера плательщиком
// Указвает сумму. Транзакция ожидает подтверждения от стартапера, что он готов
// совершить перевод инвестору. Это задействуется когда инвестор просит стартапера
// вернуть инвестированные деньги. Сумма может быть любая но
// не больше инвестированной и не меньше чем есть на счете стратапера+процент.
// Процент прибыли рассчитывает система.
func (s *Startuper) AcceptTransaction(transaction *models.Transaction) {
	// Бизнес-логика подтверждения операции
	if transaction.Err != nil {
		return
	}
	// Показать запрос отправителю
	// Получить от него явное согласие на транзакцию.
}

// Пополняет счет с внешних ресурсов(карт, кошельков и т.д)
func (s *Startuper) Topup(transaction *models.Transaction) {
	if transaction.Err != nil {
		return
	}
	// Бизнес логика пополнения внутреннего баланса
}

// Снимает деньги с внутреннего счета на счет в банке, кошелек и т.д
func (s *Startuper) Withdraw(transaction *models.Transaction) {
	if transaction.Err != nil {
		return
	}
	// Бизнес логика снятия денег со внутреннего счета
}
