package card

import (
	"github.com/akmalsulaymonov/bank/pkg/bank/types"
)

const withdrawLimit = 20_000_00
const maxAmount = 50_000
const bonusLimit = 5_000

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       123,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  50_000_00,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}

	return card
}

func Withdraw(card *types.Card, amount types.Money) {
	if card.Active && card.Balance >= amount && amount > 0 && amount <= withdrawLimit {
		card.Balance -= amount
	}
}
