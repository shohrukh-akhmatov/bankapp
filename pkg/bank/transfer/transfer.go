package transfer

// Total is
func Total(amount int) (total int) {

	if amount >= 500000 {
		total = (amount*5)/1000 + amount
	} else {
		total = amount
	}
	return total
}
