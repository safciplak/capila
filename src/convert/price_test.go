package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromCents(t *testing.T) {
	t.Run("From Int8 Cents To Float32", func(t *testing.T) {
		res, err := FromCents[float32, int8](-17)
		assert.Nil(t, err)
		assert.Equal(t, float32(-0.17), res)
	})
	t.Run("From Int8 Cents To Float32", func(t *testing.T) {
		res, err := FromCents[float32, int8](23)

		assert.Nil(t, err)
		assert.Equal(t, float32(0.23), res)
	})
	t.Run("From Int8 Cents To Float32", func(t *testing.T) {
		res, err := FromCents[float32, int8](int8(127))

		assert.Nil(t, err)
		assert.Equal(t, float32(1.27), res)
	})
	t.Run("From Int16 Cents To Float32", func(t *testing.T) {
		res, err := FromCents[float32, int16](int16(23875))

		assert.Nil(t, err)
		assert.Equal(t, float32(238.75), res)
	})
	t.Run("From Int32 Cents To Float32", func(t *testing.T) {
		res, err := FromCents[float32, int32](int32(566789))

		assert.Nil(t, err)
		assert.Equal(t, float32(5667.89), res)
	})
	t.Run("From Int64 Cents To Float32", func(t *testing.T) {
		res, err := FromCents[float32, int64](int64(12094782))

		assert.Nil(t, err)
		assert.Equal(t, float32(120947.82), res)
	})
	t.Run("From Int8 Cents To Float64", func(t *testing.T) {
		res, err := FromCents[float64, int8](int8(127))

		assert.Nil(t, err)
		assert.Equal(t, 1.27, res)
	})
	t.Run("From Int16 Cents To Float64", func(t *testing.T) {
		res, err := FromCents[float64, int16](int16(23875))

		assert.Nil(t, err)
		assert.Equal(t, 238.75, res)
	})
	t.Run("From Int32 Cents To Float64", func(t *testing.T) {
		res, err := FromCents[float64, int32](int32(566789))

		assert.Nil(t, err)
		assert.Equal(t, 5667.89, res)
	})
	t.Run("From Int64 Cents To Float64", func(t *testing.T) {
		res, err := FromCents[float64, int64](int64(12094782))

		assert.Nil(t, err)
		assert.Equal(t, 120947.82, res)
	})
}

func TestToCents(t *testing.T) {
	t.Run("From Float32 To Int Cents", func(t *testing.T) {
		res, err := ToCents[float32, int](float32(123.00))

		assert.Nil(t, err)
		assert.Equal(t, 12300, res)
	})
	t.Run("From Float64 To Int Cents", func(t *testing.T) {
		res, err := ToCents[float64, int](1903.00)

		assert.Nil(t, err)
		assert.Equal(t, 190300, res)
	})
	t.Run("From Float32 To Int Cents", func(t *testing.T) {
		res, err := ToCents[float32, uint](1903.00)

		assert.Nil(t, err)
		assert.Equal(t, uint(190300), res)
	})
	t.Run("From Float64 To Int Cents", func(t *testing.T) {
		res, err := ToCents[float64, uint](1903.00)

		assert.Nil(t, err)
		assert.Equal(t, uint(190300), res)
	})
}

func TestPriceToCents(t *testing.T) {
	t.Run("From Float32 To Int32 Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceToCents[float32, int32](167772.67)
		assert.Equal(t, int32(16777267), res)
	})
	t.Run("From Float32 To Int64 Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceToCents[float32, int64](32112.63)
		assert.Equal(t, int64(3211263), res)
	})
	t.Run("From Float32 To Int Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceToCents[float32, int](123765.12)
		assert.Equal(t, 12376512, res)
	})
	t.Run("From Float64 To Int32 Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceToCents[float64, int32](2132122.34)
		assert.Equal(t, int32(213212234), res)
	})
	t.Run("From Float64 To Int64 Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceToCents[float64, int64](3454213.21)
		assert.Equal(t, int64(345421321), res)
	})
	t.Run("From Float64 To Int Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceToCents[float64, int](5632123.78)
		assert.Equal(t, 563212378, res)
	})
}

func TestPriceFromCents(t *testing.T) {
	t.Run("From Int32 To Float32 Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceFromCents[int32, float32](16777267)
		assert.Equal(t, float32(167772.67), res)
	})
	t.Run("From Int64 To Float32 Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceFromCents[int64, float32](3211263)
		assert.Equal(t, float32(32112.63), res)
	})
	t.Run("From Int To Float32 Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceFromCents[int, float32](12376512)
		assert.Equal(t, float32(123765.12), res)
	})
	t.Run("From Int32 To Float64 Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceFromCents[int32, float64](213212234)
		assert.Equal(t, 2132122.34, res)
	})
	t.Run("From Int64 To Float64 Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceFromCents[int64, float64](345421321)
		assert.Equal(t, 3454213.21, res)
	})
	t.Run("From Int To Float64 Cents", func(t *testing.T) {
		t.Parallel()
		res := PriceFromCents[int, float64](563212378)
		assert.Equal(t, 5632123.78, res)
	})
}
