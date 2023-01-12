package card

import "bank/pkg/bank/types"

func PaymentSourses(cards []types.Card) []types.PaymentSourse {
	paysourse := make([]types.PaymentSourse, 0, len(cards))

	for _, card := range cards {
		if card.Balance <= 0 {
			continue
		}
		if !card.Active {
			continue
		}

		paysourse = append(paysourse, types.PaymentSourse{
			Type: "card",
			Number: string(card.PAN),
			Balance: card.Balance,
		})
	}

	return paysourse
	
}
