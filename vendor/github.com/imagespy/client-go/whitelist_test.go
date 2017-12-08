package imagespy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWhitelisted(t *testing.T) {
	cases := []struct {
		expect_err bool
		image      string
		result     bool
	}{
		{
			expect_err: false,
			image:      "imagespy/imagespy",
			result:     true,
		},
		{
			expect_err: false,
			image:      "private.registry:5000/app",
			result:     false,
		},
		{
			expect_err: false,
			image:      "app",
			result:     true,
		},
		{
			expect_err: true,
			image:      "app@sha256:abc",
			result:     false,
		},
	}

	for _, c := range cases {
		ok, err := isRegistryWhitelisted(c.image, DefaultRegistryWhitelist)
		assert.Equal(t, c.result, ok)
		if c.expect_err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
