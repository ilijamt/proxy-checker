package job

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Shim(a, b interface{}) []interface{} {
	return []interface{}{a, b}
}

func Test_DetailWithNoAuthenticationDetails(t *testing.T) {
	var item = Detail{
		Host: "http://example.com:9000",
	}
	var url, err = item.ToURL()
	assert.Nil(t, err)
	assert.NotNil(t, url)
	assert.EqualValues(t, "example.com:9000", url.Host)
	assert.EqualValues(t, "http", url.Scheme)
}

func Test_DetailWithIncompleteAuthenticationDetails(t *testing.T) {
	var item = Detail{
		Host:     "http://example.com:9000",
		Username: "user",
	}
	var url, err = item.ToURL()
	assert.Nil(t, err)
	assert.NotNil(t, url)
	assert.EqualValues(t, "example.com:9000", url.Host)
	assert.EqualValues(t, "http", url.Scheme)
	assert.Nil(t, url.User)
}

func Test_DetailWithAuthenticationDetails(t *testing.T) {
	var item = Detail{
		Host:     "http://example.com:9000",
		Username: "user",
		Password: "pass",
	}
	var url, err = item.ToURL()
	assert.Nil(t, err)
	assert.NotNil(t, url)
	assert.EqualValues(t, "example.com:9000", url.Host)
	assert.EqualValues(t, "http", url.Scheme)
	assert.NotNil(t, url.User)
	assert.EqualValues(t, "http://user:pass@example.com:9000", url.String())
}
