package utilities

import "testing"

var checkFileTestCases = []struct {
	filepath string
	exist    bool
}{
	{"../generator/data/text.in", true},
	{"../generator/data/text.text", false},
}

func TestCheckFile(t *testing.T) {
	for _, test := range checkFileTestCases {
		obs, _ := checkFile(test.filepath)
		if obs != test.exist {
			t.Errorf("for: %s, expected: %t, got %t", test.filepath, obs, test.exist)
		}
	}
}
