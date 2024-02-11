// Interfaces
// Several else are in models. Not here to avoid cycle import
package interfaces

import "github.com/udonetsm/investing/models"

type Saver interface {
	Save(*models.Transaction)
}

type Requester interface {
	RequestTransaction(*models.Transaction)
}

type Accepter interface {
	AcceptTransaction(*models.Transaction)
}

// Полный список транзакций, которые делает система.
type Tranactioner interface {
	TransferTransaction(*models.Transaction)
	TopupTransaction(*models.Transaction)
	WithdrawTransaction(*models.Transaction)
}
