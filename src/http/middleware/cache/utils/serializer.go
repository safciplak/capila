// Package utils is a sub helper package for cache
//
// The serializer is helper for serialize/deserialize for cache data
package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"strconv"
)

// Serialize returns a []byte representing the passed value
func Serialize(value interface{}) ([]byte, error) {
	if valueAsBytes, ok := value.([]byte); ok {
		return valueAsBytes, nil
	}

	switch val := reflect.ValueOf(value); val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		return []byte(strconv.FormatInt(val.Int(), 10)), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		return []byte(strconv.FormatUint(val.Uint(), 10)), nil
	}

	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(value); err != nil {
		return nil, fmt.Errorf("serialize encode error: %w", err)
	}

	return buffer.Bytes(), nil
}

// Deserialize deserializes the passed []byte into a passed value interface{}
func Deserialize(byt []byte, value interface{}) error {
	if valueAsBytes, ok := value.(*[]byte); ok {
		*valueAsBytes = byt

		return nil
	}

	if val := reflect.ValueOf(value); val.Kind() == reflect.Ptr {
		switch ptr := val.Elem(); ptr.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(string(byt), 10, 64)

			if err != nil {
				return fmt.Errorf("parseint error: %w", err)
			}

			ptr.SetInt(intVal)

			return nil

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintVal, err := strconv.ParseUint(string(byt), 10, 64)

			if err != nil {
				return fmt.Errorf("parseUint error: %w", err)
			}

			ptr.SetUint(uintVal)

			return nil
		}
	}

	b := bytes.NewBuffer(byt)
	decoder := gob.NewDecoder(b)

	return decoder.Decode(value)
}
