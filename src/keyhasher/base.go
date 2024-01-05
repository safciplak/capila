package keyhasher

import (
	"errors"

	"github.com/speps/go-hashids"
)

// "fixed" types you can choose from
const (
	None hashType = iota
	HashID
)

const hashIDLength = 10

type hashType int

//nolint:gochecknoglobals // needs to be refactored
var hashIDer *hashids.HashID

/*
Initialize initializes the hasher.
do this just once (per hashType) at the program start
*/
func Initialize(wantType hashType, salt string) error {
	if salt == "" {
		return errors.New("no salt specified")
	}

	var err error

	if wantType == HashID {
		err = initHashID(salt)
	} else {
		err = errors.New("keyhasher > Initialize : wrong hashType specified")
	}

	return err
}

/*
Encode encodes a uint number to a hash.
*/
func Encode(number uint) (string, error) {
	// only written for hashID a.t.m.
	if hashIDer == nil {
		return "", errors.New("keyhasher > Encode : not initialized yet")
	}

	numbers := []int{int(number)}

	return hashIDer.Encode(numbers)
}

/*
Decode decodes a hash into an uint number.
*/
func Decode(hashedValue string) (uint, error) {
	// only written for hashID a.t.m.
	if hashIDer == nil {
		return 0, errors.New("keyhasher > Decode : not initialized yet")
	}

	slice, err := hashIDer.DecodeWithError(hashedValue)

	if err != nil {
		return 0, errors.New("keyhasher > Decode : no valid hash")
	}

	return uint(slice[0]), nil
}

/*
initHashID initializes the type "hashID"
*/
func initHashID(salt string) error {
	config := hashids.NewData()
	config.Salt = salt
	config.MinLength = hashIDLength

	var err error

	hashIDer, err = hashids.NewWithData(config)

	return err
}
