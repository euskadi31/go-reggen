package reggen

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultGenerate(t *testing.T) {
	s, err := Generate("[A-F0-9]{12}")
	assert.NoError(t, err)

	assert.True(t, regexp.MustCompile("[A-F0-9]{12}").MatchString(s))
}

func TestDefaultGenerateWithBadRegexp(t *testing.T) {
	_, err := Generate(`(?P<name>a`)
	assert.Error(t, err)
}
