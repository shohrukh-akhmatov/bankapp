package payment

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID: 001,
			Amount: 10,
		},
		{
			ID: 002,
			Amount: 20,
		},
		{
			ID: 003,
			Amount: 20,
		}, 
		{
			ID: 004,
			Amount: 99,
		},
		{
			ID: 99,
			Amount: 100,
		},
		{
			ID: 100,
			Amount: 100,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	//Output:99
	
}