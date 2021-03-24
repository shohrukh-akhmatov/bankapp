package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleDeposit_positive() {
	card := types.Card{Balance: 10_000_00, Active: true}
	Deposit (&card, 10_000_00)

	fmt.Println(card.Balance)
	// Output: 2000000
}
func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit (&card, 20_000_00)

	fmt.Println(card.Balance)
	// Output: 2000000
}
func ExampleDeposit_limit() { 
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit (&card, 51_000_00)

	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_positive() { 
	card := types.Card{Balance: 10_000_00, MinBalance: 10_000_00, Active: true,} 
	AddBonus (&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1002465
}
func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 10_000_00, MinBalance: 10_000_00, Active: false}
	AddBonus (&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 1000000
}
func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -10_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus (&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: -1000000
}

func ExampleAddBonus_limitAbove() {
	card := types.Card{Balance: 10_000_00, MinBalance: 900_000_000_00, Active: true}
	AddBonus (&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: -20_000_00,
			Active: true,
		},
		{
			Balance: 30_000_00,
			Active: false, 
		},
		{
			Balance: 25_000_00,
			Active: true,
		},
	}
	
	fmt.Println(Total(cards))
		//Output:3500000
	
}