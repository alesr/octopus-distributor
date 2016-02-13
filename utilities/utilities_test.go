package utilities

import "testing"

var randomErrorCases = []struct {
	arg int
	err error
}{
	{-2, ErrRandomFuncError},
	{-1, ErrRandomFuncError},
	{0, ErrRandomFuncError},
	{1, nil},
	{2, nil},
}

// not possibile to test randoms, so at least let's check if error is working
func TestRandom(t *testing.T) {
	for _, test := range randomErrorCases {
		_, err := Random(test.arg)
		if err != nil {
			if err != test.err {
				t.Error(err)
			}
		}
	}
}

var checkFileTestCases = []struct {
	filepath string
	exist    bool
}{
	{"../generator/data/text.in", true},
	{"../generator/data/text.text", false},
	{"../data/text.text", false},
}

func TestCheckFile(t *testing.T) {
	for _, test := range checkFileTestCases {
		obs, _ := checkFile(test.filepath)
		if obs != test.exist {
			t.Errorf("for: %s, expected: %t, got %t", test.filepath, test.exist, obs)
		}
	}
}

var loadFileCases = []struct {
	filepath string
	content  []string
}{
	{"../generator/data/text.in",
		[]string{
			"Lorem ipsum dolor sit amet.",
			"Hello, World",
			"foo bar",
			"I'll make him an offer he can't refuse.",
			"You shall not pass.",
			"The quick brown fox jumps over the lazy dog.",
			"Papa Americano"}},
}

func TestLoadFile(t *testing.T) {
	for _, test := range loadFileCases {
		obs, err := LoadFile(test.filepath)
		if err != nil {
			t.Error(err)
		}

		for i := range obs {
			if obs[i] != test.content[i] {
				t.Error("for: %s, expected: %s, got: %s", test.filepath, string(test.content[i]), string(obs[i]))
			}
		}
	}
}
