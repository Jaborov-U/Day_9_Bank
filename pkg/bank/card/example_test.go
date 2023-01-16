package card

import (
	"bank/pkg/bank/types"
	"fmt"
)



func ExamplePaymentSourses() {
	cards := []types.Card{
		{
			Balance: 5_000,
			Active: true,
		},
		{
			Balance: -10,
			Active: true,
		},
	}

	paySourse := PaymentSourses(cards)

	fmt.Println(len(paySourse))
	
	//Output:1
}
