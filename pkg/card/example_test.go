package card

import (
	"fmt"

	"github.com/akmalsulaymonov/bank/pkg/bank/types"
)

func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 0, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 0
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 40_000_00, Active: true}
	Withdraw(&card, 30_000_00)
	fmt.Println(card.Balance)
	// Output: 4000000
}
