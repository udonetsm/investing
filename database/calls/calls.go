package calls

import "github.com/udonetsm/investing/models"

// Сохранить в базу информацию о транзакции
// Транзакция сохраняется в любом случае, даже если произошла ошибка
// Если транзакция успешная, то поле Success переводится в true
// По умолчанию поле success = false.
// Это значит транзакция неуспешная по умолчанию.
func SaveTransactios(transaction models.Transaction) error {
	var err error = nil
	// err:= пытаемся сохранить транзакцию в бд. Можем получить ошибку
	// это временно, пока не сделаю функцию сохранения транзакции в бд
	if err != nil {
		return err
	}
	// Сохранить в базе, продумать формат хранения внутренних транзакций.
	return nil
}
