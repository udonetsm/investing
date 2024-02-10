// Interfaces
// Several else are in models. Not here to avoid cycle import
package interfaces

type Saver interface {
	Save(any) error
}

type BalanceUpdater interface {
	BalanceUpdate(any) error
}
