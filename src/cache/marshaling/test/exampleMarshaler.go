package test

import (
	"errors"
)

type binaryMarshaler struct {
	ValueA string
}

// MarshalBinary implements encoding.BinaryMarshaler
//
//nolint:unparam // implementing interface
func (m *binaryMarshaler) MarshalBinary() (data []byte, err error) {
	return []byte(m.ValueA), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
//
//nolint:unparam // implementing interface
func (m *binaryMarshaler) UnmarshalBinary(data []byte) error {
	m.ValueA = string(data)

	return nil
}

// This object is considered completely valid when its Value starts with 'Va'
// it is considered valid for marshaling when the second character is an 'a'
// it is considered valid for unmarshaling when the first character is a 'V'

type binaryMarshalerWithValidation struct {
	Value string
}

// MarshalBinary implements encoding.BinaryMarshaler
func (m *binaryMarshalerWithValidation) MarshalBinary() ([]byte, error) {
	data := []byte(m.Value)
	if len(data) < 2 || data[1] != 'a' {
		return nil, errors.New("invalid Object to marshal")
	}

	return data, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (m *binaryMarshalerWithValidation) UnmarshalBinary(data []byte) error {
	if len(data) == 0 || data[0] != 'V' {
		return errors.New("invalid Object to unmarshal")
	}

	m.Value = string(data)

	return nil
}
