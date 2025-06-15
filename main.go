package main

import (
	"fmt"
	"go-thai-currency-converter/converter"

	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(20),
		decimal.NewFromFloat(21),
		decimal.NewFromFloat(101),
		decimal.NewFromFloat(1001),
		decimal.NewFromFloat(1000001),
		decimal.NewFromFloat(0.01),
		decimal.NewFromFloat(0.21),
		decimal.NewFromFloat(999999.99),
		decimal.NewFromFloat(12345678.90),
		decimal.NewFromFloat(99999999.99),
		decimal.NewFromFloat(1000000000),
		decimal.NewFromFloat(9999999999.99),
		decimal.NewFromFloat(123456789012.34),
	}

	for _, input := range inputs {
		fmt.Println(input)
		// convert decimal to thai text (baht) and print the result here
		result := converter.ThaiCurrencyConverter(input)
		fmt.Println(result)
		fmt.Println()
	}
}
