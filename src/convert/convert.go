package convert

import (
	"strconv"
)

// NewString turns a normal string into a pointer string.
func NewString(value string) *string {
	return &value
}

// NewBool turns a normal bool into a pointer bool.
func NewBool(value bool) *bool {
	return &value
}

// NewFloat32 turns a normal float32 into a pointer float32.
func NewFloat32(value float32) *float32 {
	return &value
}

// NewFloat64 turns a normal float64 into a pointer float64.
func NewFloat64(value float64) *float64 {
	return &value
}

// NewUint turns a normal uint into a pointer uint.
func NewUint(value uint) *uint {
	return &value
}

// NewInteger turns a normal integer into a pointer integer
func NewInteger(value int) *int {
	return &value
}

// NewInteger32 turns a normal int32 into a pointer int32
func NewInteger32(value int32) *int32 {
	return &value
}

// NewInteger64 turns a normal int64 into a pointer int64
func NewInteger64(value int64) *int64 {
	return &value
}

// StringToBool converts a string to a boolean value defaulting to false.
func StringToBool(value string) bool {
	valueAsBool, _ := strconv.ParseBool(value)
	return valueAsBool
}

// PointerToString converts a string pointer to a string.
func PointerToString(stringPointer *string) string {
	var result string

	if stringPointer != nil {
		result = *stringPointer
	}

	return result
}

// IntPointerToStringPointer converts an int pointer to a string pointer
func IntPointerToStringPointer(intPointer *int) *string {
	return NewString(strconv.Itoa(*intPointer))
}
