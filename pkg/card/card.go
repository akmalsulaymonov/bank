package card

import (
	"github.com/akmalsulaymonov/bank/pkg/bank/types"
)

const withdrawLimit = 20_000_00

func Withdraw(card *types.Card, amount types.Money) {
	if card.Active && card.Balance >= amount && amount > 0 && amount <= withdrawLimit {
		card.Balance -= amount
	}
}
