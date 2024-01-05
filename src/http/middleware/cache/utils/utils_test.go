package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlEscape(t *testing.T) {
	t.Parallel()

	prefix := "my_test"
	u := "/hotels?size=20"
	s := URLEscape(prefix, u)

	assert.Equal(t, "my_test:%2Fhotels%3Fsize%3D20", s)
}

func TestUrlEscape_MaxLength(t *testing.T) {
	t.Parallel()

	prefix := "my_test"
	u := "https://uat2-bookingtool-platform.vandervalkonline.com/hotels/single-hotel/theatm.vandervalkonline.com/hotels/" +
		"single-hotel/theater-hotel/configurahotels/sorm.vandervalkonline.com/hotels/single-hotel/theater-hotel/configur" +
		"ahotels/single-hotel/theater-hotel/configuration/hotels?size=20"
	s := URLEscape(prefix, u)

	assert.Equal(t, "my_test:\xdbA\xcd\xcaX\xdaČ\v\\\xb6\xa4\x8b\xa1\x1b\xd3\xf2;\x82\x90", s)
}
