// General functions
package general

import (
	"fmt"

	"github.com/udonetsm/investing/interfaces"
	"github.com/udonetsm/investing/models"
)

type Payer models.DBEntries

func (s Payer) Get(transaction_id string) (err error) {
	// err=Проверить наличие транзакции на счете для Getter-a
	fmt.Print("Проверка транзакции...")
	return
}

// err = Проверка банковского счета стартапера на наличие нужной суммы
func (s Payer) Pay(to interfaces.Getter, sum float32) (transaction_id string, err error) {
	// Если возникла какая-либо ошибка, транзакция не осуществляется
	// и ошибка возвращается "наверх". Возможно не хватает средств или банк отклонил соединение
	if err != nil {
		// Ошибку проверки баланса можно кастомизировать здесь
		return
	}
	// Если средств достаточно и ошибки нет то запрашиваем у банка транзакцию.
	// transaction_id, err = Делаем внешнюю транзакцию и получаем ее идентификатор
	if err != nil {
		// Неудача во время перевода
		return
	}
	// Если все ок, то делаем лог и сообщение об успехе
	// fmt.Printf("%v send %.2f to %v...", s, sum, to)
	// И записываем транзакцию в базу, в таблицу Transactions
	return
}
