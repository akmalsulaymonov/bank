package types

//Money for Balance
type Money int64

// Category type
type Category string

// Payment type
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}

//Currency for type money
type Currency string

// Constants
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// PAN type
type PAN string

// Card type
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

// Payment source type
type PaymentSource struct {
	Type    string // 'card'
	Number  string // номер вида '5058 xxxx xxxx 8888'
	Balance Money  // баланс в дирамах
}
