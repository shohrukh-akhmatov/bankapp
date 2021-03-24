package payment 

import (
	"bank/pkg/bank/types"
)

//Max returns max payment from cards
func Max (payments [] types.Payment) types.Payment {
	max := payments[0]
	for _, operation := range payments {

		if max.Amount < operation.Amount {
			max = operation
		}
	}
	return max
}