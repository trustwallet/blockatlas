package spamfilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isURL(t *testing.T) {
	assert.True(t, isURL("http://www.trust.com"))
	assert.True(t, isURL("www.trust.com"))
	assert.True(t, isURL("trust.com"))
	assert.True(t, isURL("saa,c trust.com"))
}
