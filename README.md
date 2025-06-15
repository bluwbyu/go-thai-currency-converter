# Thai Currency Converter

Go library สำหรับแปลงจำนวนเงินเป็นข้อความภาษาไทย

## Quick Start

### Installation

```bash
git clone https://github.com/bluwbyu/go-thai-currency-converter.git
cd go-thai-currency-converter
go mod tidy
```

### Usage

```go
import (
    "go-thai-currency-converter/converter"
    "github.com/shopspring/decimal"
)

amount := decimal.NewFromFloat(1234.75)
result := converter.ThaiCurrencyConverter(amount)
// Output: หนึ่งพันสองร้อยสามสิบสี่บาทเจ็ดสิบห้าสตางค์
```

## Commands

### Run Demo

```bash
go run main.go
```

## Examples

| Input      | Output                                              |
| ---------- | --------------------------------------------------- |
| `1234`     | `หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน`                   |
| `33333.75` | `สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์` |
| `20`       | `ยี่สิบบาทถ้วน`                                     |
| `101`      | `หนึ่งร้อยเอ็ดบาทถ้วน`                              |
