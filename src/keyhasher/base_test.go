package keyhasher_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/safciplak/capila/src/keyhasher"
)

const (
	appKey    = "123456"
	oneHashed = "DdakboKbJv"
	twoHashed = "V1oKzxE96Y"
)

func Test_Initialize_Wrong_HashType(t *testing.T) {
	err := keyhasher.Initialize(keyhasher.None, appKey)

	assert.NotNil(t, err)
	assert.Equal(t, `keyhasher > Initialize : wrong hashType specified`, err.Error())
}

func Test_Encode_Without_Init(t *testing.T) {
	hashNumber, err := keyhasher.Encode(1)

	assert.NotNil(t, err)
	assert.Equal(t, `keyhasher > Encode : not initialized yet`, err.Error())
	assert.EqualValues(t, "", hashNumber)
}

func Test_Decode_Without_Init(t *testing.T) {
	number, err := keyhasher.Decode(oneHashed)

	assert.NotNil(t, err)
	assert.Equal(t, `keyhasher > Decode : not initialized yet`, err.Error())
	assert.EqualValues(t, 0, number)
}

func Test_Initialize(t *testing.T) {
	err := keyhasher.Initialize(keyhasher.HashID, appKey)

	assert.Nil(t, err)
}

func Test_Encode(t *testing.T) {
	err := keyhasher.Initialize(keyhasher.HashID, appKey)

	assert.Nil(t, err)

	hashNumber, err := keyhasher.Encode(1)

	assert.Nil(t, err)
	assert.NotEqual(t, "", hashNumber)
	assert.Equal(t, oneHashed, hashNumber)

	hashNumber, err = keyhasher.Encode(2)

	assert.Nil(t, err)
	assert.NotEqual(t, "", hashNumber)
	assert.Equal(t, twoHashed, hashNumber)
}

func Test_Decode(t *testing.T) {
	err := keyhasher.Initialize(keyhasher.HashID, appKey)

	assert.Nil(t, err)

	number, err := keyhasher.Decode(oneHashed)

	assert.Nil(t, err)
	assert.EqualValues(t, 1, number)

	number, err = keyhasher.Decode(twoHashed)

	assert.Nil(t, err)
	assert.EqualValues(t, 2, number)
}

func Test_DecodeUintWithRandomString(t *testing.T) {
	err := keyhasher.Initialize(keyhasher.HashID, appKey)

	assert.Nil(t, err)

	_, err = keyhasher.Decode("Ditisgeenhashjonguh")

	assert.NotNil(t, err)
	assert.Equal(t, `keyhasher > Decode : no valid hash`, err.Error())
}
