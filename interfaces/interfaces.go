// Interfaces
// Several else are in models. Not here to avoid cycle import
package interfaces

import "github.com/udonetsm/investing/models"

type Saver interface {
	Save(*models.Transaction)
}

type Updater interface {
}

type SaveUpdater interface {
	Saver
	Updater
}

// Полный список транзакций, которые делает система.
type TransaferTransactioner interface {
	TransferTransaction(*models.Transaction)
}

type TopupTrnsactioner interface {
	TopupTransaction(*models.Transaction)
}

type WithdrawTransactioner interface {
	WithdrawTransaction(*models.Transaction)
}

type TransferTopupWithdrawTransactioner interface {
	TransaferTransactioner
	TopupTrnsactioner
	WithdrawTransactioner
}
