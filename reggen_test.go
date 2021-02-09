package reggen

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	regex string
}

var cases = []testCase{
	{`123[0-2]+.*\w{3}`},
	{`^\d{1,2}[/](1[0-2]|[1-9])[/]((19|20)\d{2})$`},
	{`^((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])$`},
	{`^\d+$`},
	{`\D{3}`},
	{`((123)?){3}`},
	{`(ab|bc)def`},
	{`[^abcdef]{5}`},
	{`[^1]{3,5}`},
	{`[[:upper:]]{5}`},
	{`[^0-5a-z\s]{5}`},
	{`Z{2,5}`},
	{`[a-zA-Z]{100}`},
	{`^[a-z]{5,10}@[a-z]{5,10}\.(com|net|org)$`},
}

func TestGenerate(t *testing.T) {
	for _, test := range cases {
		for i := 0; i < 10; i++ {
			r, err := NewGenerator(test.regex)
			assert.NoError(t, err)

			res := r.Generate(10)

			re, err := regexp.Compile(test.regex)
			assert.NoError(t, err)

			assert.True(t, re.MatchString(res), "Generated data does not match regex. Regex: %q output: %q", test.regex, res)
		}
	}
}

func TestSeed(t *testing.T) {
	g1, err := NewGenerator(cases[0].regex)
	assert.NoError(t, err)

	g2, err := NewGenerator(cases[0].regex)
	assert.NoError(t, err)

	currentTime := time.Now().UnixNano()

	g1.Seed(currentTime)
	g2.Seed(currentTime)

	for i := 0; i < 10; i++ {
		assert.Equal(t, g1.Generate(100), g2.Generate(100))
	}

	g1.Seed(123)
	g2.Seed(456)

	for i := 0; i < 10; i++ {
		assert.NotEqual(t, g1.Generate(100), g2.Generate(100))
	}
}

func BenchmarkGenerate(b *testing.B) {
	r, err := NewGenerator(`^[a-z]{5,10}@[a-z]+\.(com|net|org)$`)
	if err != nil {
		b.Fatal("Error creating generator: ", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Generate(10)
	}
}
