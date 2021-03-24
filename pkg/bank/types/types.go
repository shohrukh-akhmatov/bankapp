package types

// Money presents money in minimal units
type Money int64

// Currency presents code of currencies
type Currency string

// Code of currencies
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN presents number of card
type PAN string

// Card presents information about card
type Card struct {
	ID       	int
	PAN     	PAN
	Balance  	Money
	Currency 	Currency
	Color    	string
	Name     	string
	Active   	bool
	MinBalance 	Money
}


// Payment presents information about transfers
type Payment struct {
	ID 		int
	Amount 	Money
}


