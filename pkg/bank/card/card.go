package card

import (
	"bank/pkg/bank/types"
)

// IssueCard issues cards
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       	1000,
		PAN:      	"5058 xxxx xxxx 0001",
		Balance:  	0,
		Currency: 	currency,
		Color:    	color,
		Name:     	name,
		Active:   	true,
	}
	return card
}

// Deposit transfers money to card
func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}

	if amount < 0 {
		return
	}

	if amount > 50_000_00 {
		return
	}

	card.Balance += amount

	return
}

// AddBonus addes bonuses 
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	percentMoney := card.MinBalance * types.Money(percent) / 100
	bonus := percentMoney * types.Money(daysInMonth) / types.Money(daysInYear)

	if !card.Active {
		return
	}

	if card.Balance < 0 { 
		return
	}

	if bonus > 5_000_00 {
		return
	}
	card.Balance += bonus	
	return
}


// Total counts sum of all money in cards
// Negative sums and sums from blocked cards will be ignored
func Total(cards []types.Card)  types.Money {
	sum := types.Money(0)
	for _, operation := range cards {
		
		if !operation.Active {
			continue
		}
		if operation.Balance < 0 {
			continue
		}
		sum += operation.Balance
	}
	
	return sum
}
