// Interfaces
package interfaces

type Getter interface {
	Get(string) error
}

type Payer interface {
	Pay(Getter, float32) (string, error)
}
