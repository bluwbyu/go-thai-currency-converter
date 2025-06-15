package converter

import (
	"github.com/shopspring/decimal"
)

// Thai number words
var (
	ones = []string{"", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
)

// ThaiCurrencyConverter converts a decimal amount to Thai text with currency suffix.
//
// Example:
//
//	ThaiCurrencyConverter(decimal.NewFromInt(1234)) returns "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"
//	ThaiCurrencyConverter(decimal.NewFromFloat(33333.75)) returns "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"
func ThaiCurrencyConverter(amount decimal.Decimal) string {
	integerPart := amount.Truncate(0)
	fractionalPart := amount.Sub(integerPart).Mul(decimal.NewFromInt(100)).Truncate(0)

	result := convertNumberToThai(integerPart.IntPart()) + "บาท"

	if fractionalPart.IsZero() {
		result += "ถ้วน"
	} else {
		result += convertNumberToThai(fractionalPart.IntPart()) + "สตางค์"
	}

	return result
}

// convertNumberToThai converts an integer to Thai text.
func convertNumberToThai(num int64) string {
	switch {
	case num == 0:
		return "ศูนย์"
	case num < 0:
		return "ลบ" + convertNumberToThai(-num)
	default:
		return convertPositiveNumber(num)
	}
}

// convertPositiveNumber handles positive numbers.
func convertPositiveNumber(num int64) string {
	if num == 0 {
		return ""
	}

	result := ""
	hasHigherDigits := false

	// Handle millions
	if num >= 1000000 {
		millions := num / 1000000
		result += convertUpToThousands(millions, false) + "ล้าน"
		num %= 1000000
		hasHigherDigits = true
	}

	// Handle remaining part (up to 999,999)
	if num > 0 {
		result += convertUpToThousands(num, hasHigherDigits)
	}

	return result
}

// convertUpToThousands handles numbers from 1 to 999,999 with context awareness.
func convertUpToThousands(num int64, hasHigherDigits bool) string {
	if num == 0 {
		return ""
	}

	result := ""

	// Handle hundred thousands (แสน)
	if num >= 100000 {
		result += ones[num/100000] + "แสน"
		num %= 100000
	}

	// Handle ten thousands (หมื่น)
	if num >= 10000 {
		result += ones[num/10000] + "หมื่น"
		num %= 10000
	}

	// Handle thousands (พัน)
	if num >= 1000 {
		thousands := num / 1000
		if thousands == 1 {
			result += "หนึ่งพัน"
		} else {
			result += ones[thousands] + "พัน"
		}
		num %= 1000
	}

	// Handle hundreds (ร้อย)
	if num >= 100 {
		result += ones[num/100] + "ร้อย"
		num %= 100
	}

	// Handle tens and ones
	result += convertTensAndOnes(num, result != "" || hasHigherDigits)

	return result
}

// convertTensAndOnes handles numbers 0-99 with context awareness.
func convertTensAndOnes(num int64, hasHigherDigits bool) string {
	switch {
	case num >= 20:
		return convertTwentyToNinetyNine(num)
	case num >= 10:
		return convertTenToNineteen(num)
	case num > 0:
		return convertOnes(num, hasHigherDigits)
	default:
		return ""
	}
}

// convertTwentyToNinetyNine handles numbers 20-99.
func convertTwentyToNinetyNine(num int64) string {
	tens := num / 10
	remainder := num % 10

	var result string
	if tens == 2 {
		result = "ยี่สิบ"
	} else {
		result = ones[tens] + "สิบ"
	}

	if remainder > 0 {
		if remainder == 1 {
			result += "เอ็ด"
		} else {
			result += ones[remainder]
		}
	}

	return result
}

// convertTenToNineteen handles numbers 10-19.
func convertTenToNineteen(num int64) string {
	switch num {
	case 10:
		return "สิบ"
	case 11:
		return "สิบเอ็ด"
	default:
		return "สิบ" + ones[num%10]
	}
}

// convertOnes handles single digits 1-9 with context awareness.
func convertOnes(num int64, hasHigherDigits bool) string {
	if hasHigherDigits && num == 1 {
		return "เอ็ด"
	}
	return ones[num]
}
