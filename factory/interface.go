package factory

// Payment interface
type Payment interface {
	Pay(amount float64) error
}
