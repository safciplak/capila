package convert

import (
	"errors"
	"fmt"
	"math"

	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
)

const convertRate = 100.0

type Number interface {
	constraints.Integer | constraints.Float
}

// FromCents converts any integer cents to float value
// Deprecated: FromCents use PriceFromCents instead
func FromCents[To constraints.Float, From constraints.Integer](from From) (To, error) {
	to := To(from)
	if From(to) != from {
		return to, errors.New("conversion out of range")
	}

	return to / convertRate, nil
}

// Deprecated: ToCents use PriceToCents instead
func ToCents[From Number, To constraints.Integer](from From) (To, error) {
	res := math.RoundToEven(float64(from) * convertRate)
	to := To(res)
	if From(to)/convertRate != from {
		return to, errors.New("conversion out of range")
	}

	return to, nil
}

// Disclaimer: This function does not provide a guarantee if you exceed max safe values.
// PriceToCents given price value converts to integer cents
// Max Safe values;
// for int32 -> 21474836
// for int64 -> 92233720368547758
func PriceToCents[From constraints.Float, To constraints.Integer](from From) To {
	stringPrice := fmt.Sprintf("%.2f", from)
	decimalPrice := decimal.RequireFromString(stringPrice)
	hundred := decimal.New(convertRate, 0)
	res := decimalPrice.Mul(hundred)
	to := To(res.BigInt().Int64())

	return to
}

// Disclaimer: This function does not provide a guarantee if you exceed max safe values.
// PriceFromCents given price cent values convert to float cents
// Max Safe values;
// for int32 -> 21474836
// for int64 -> 92233720368547758
func PriceFromCents[From constraints.Integer, To constraints.Float](from From) To {
	stringPrice := fmt.Sprintf("%d", from)
	decimalPrice := decimal.RequireFromString(stringPrice)
	hundred := decimal.New(convertRate, 0)
	res := decimalPrice.Div(hundred)

	float64Value, _ := res.BigFloat().Float64()
	to := To(float64Value)

	return to
}
