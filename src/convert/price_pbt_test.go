package convert_test

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"

	"github.com/safciplak/capila/src/convert"
)

type Valuta struct {
	Main  int32
	Cents int8
}

func (v *Valuta) toFloat() (float64, error) {
	return strconv.ParseFloat(v.String(), 32)
}

func (v *Valuta) String() string {
	return fmt.Sprintf("%d.%d", v.Main, v.Cents)
}

func (v *Valuta) CentsString() string {
	return fmt.Sprintf("%d%d", v.Main, v.Cents)
}

func TestPriceToCents_onlyCentRangePrecision(t *testing.T) {
	properties := gopter.NewProperties(nil)

	// fmt.Println(math.MinInt32/100, math.MaxInt32/100)
	arbitraryValuta := gen.Struct(reflect.TypeOf(Valuta{}), map[string]gopter.Gen{
		"Cents": gen.Int8Range(0, 99),
		"Main":  gen.Int32Range(0, math.MaxInt32/100), // main not higher than math.MaxInt32/100 => 2.147.4836 +- max safe value
	})

	properties.Property("conversion and back does not lose information", prop.ForAll(func(v Valuta) string {
		input64, _ := v.toFloat()

		input := float32(input64)
		cents := convert.PriceToCents[float32, int32](input)

		asFloat := convert.PriceFromCents[int32, float32](cents)

		asCentsAgain := convert.PriceToCents[float32, int32](asFloat)

		if asFloat != input {
			return fmt.Sprintf("first conversion back is not equal to original input valuta: %+v input64: %v, input: %v, cents: %v, asFloat: %v", v, input64, input, cents, asFloat)
		}

		if cents != asCentsAgain {
			return fmt.Sprintf("second conversion back is not equal to first converted value valuta: %+v  input64: %v, input: %v, cents: %v, asFloat: %v asCentsAgain: %v", v, input64, input, cents, asFloat, asCentsAgain)
		}

		return ""
	},
		arbitraryValuta))

	properties.TestingRun(t)

}
