package utils

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testObject struct {
	name string
	val1 int64
	val2 uint32
}

type testObject2 struct {
	Name string
	Val1 int64
	Val2 uint32
}

func TestSerialize(t *testing.T) {
	t.Parallel()

	obj := testObject2{"valk", int64(10), uint32(20)}
	var decodedObj testObject2

	serializedBytes, err := Serialize(obj)
	decodeErr := Deserialize(serializedBytes, &decodedObj)

	assert.Nil(t, err)
	assert.Nil(t, decodeErr)
	assert.NotNil(t, serializedBytes)
	assert.Equal(t, obj, decodedObj)

	if !reflect.DeepEqual(obj, decodedObj) {
		t.Fatalf("TestSerialize: expected %v, got %v", obj, decodedObj)
	}
}

func TestSerialize_PrimitiveTypes(t *testing.T) {
	t.Parallel()

	serializedBytes, err := Serialize(int64(10))
	serializedBytes2, err2 := Serialize(uint64(10))

	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.NotNil(t, serializedBytes)
	assert.NotNil(t, serializedBytes2)
}

func TestSerialize_AsBytes(t *testing.T) {
	t.Parallel()

	myVal := []byte("hello")
	serializedBytes, err := Serialize(myVal)

	assert.Nil(t, err)
	assert.Equal(t, myVal, serializedBytes)
}

func TestSerialize_Error(t *testing.T) {
	t.Parallel()

	obj := testObject{"valk", int64(10), uint32(20)}

	serializedBytes, err := Serialize(obj)

	assert.Nil(t, serializedBytes)
	assert.NotNil(t, err)
}

func TestDeserialize(t *testing.T) {
	t.Parallel()

	var obj testObject2

	serializedBytes, _ := Serialize(testObject2{"valk", int64(10), uint32(20)})

	err := Deserialize(serializedBytes, &obj)

	assert.Nil(t, err)
	assert.Equal(t, "valk", obj.Name)
	assert.Equal(t, int64(10), obj.Val1)
	assert.Equal(t, uint32(20), obj.Val2)
}

func TestDeserialize_Error(t *testing.T) {
	t.Parallel()

	var obj testObject

	serializedBytes, _ := Serialize(testObject2{"valk", int64(10), uint32(20)})

	err := Deserialize(serializedBytes, &obj)

	assert.NotNil(t, err)
}

func TestDeserialize_PrimitiveTypes(t *testing.T) {
	t.Parallel()

	var myVal int64
	var myVal2 uint64

	err := Deserialize([]byte(strconv.Itoa(10)), &myVal)
	err2 := Deserialize([]byte(strconv.FormatUint(uint64(20), 10)), &myVal2)

	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.Equal(t, int64(10), myVal)
	assert.Equal(t, uint64(20), myVal2)
}

func TestDeserialize_AsBytes(t *testing.T) {
	t.Parallel()

	var myVal []byte
	err := Deserialize([]byte("hello"), &myVal)

	assert.Nil(t, err)
	assert.Equal(t, []byte("hello"), myVal)
}
