package general

import (
	"github.com/udonetsm/investing/interfaces"
	"github.com/udonetsm/investing/models"
)

// The function Saves any object into any spaces.
// To use it, space object sould implement its method
// Save(any)error
func SaveSomething(saver interfaces.Saver, transaction *models.Transaction) {
	saver.Save(transaction)
}

// Задействовать методы переданного интерфейсного типа и передать им транзакцию для обрабобтки
func RequestTansaction(requester models.BaseUser, transaction *models.Transaction) {
	requester.RequestTransaction(transaction)
}

// Задействовать методы переданного интерфейсного типа и передать им транзакцию для обрабобтки
func AcceptTransaction(accepter models.BaseUser, transaction *models.Transaction) {
	accepter.AcceptTransaction(transaction)
}

// Задействовать методы переданного интерфейсного типа и передать им транзакцию для обрабобтки
func MakeTransaction(transactioner interfaces.TransferTopupWithdrawTransactioner, transaction *models.Transaction) {
	transactioner.TransferTransaction(transaction)
}

// Делает плательщику запрос на подтверждение операции
func MakeRequest(transaction *models.Transaction) error {
	return nil
}
