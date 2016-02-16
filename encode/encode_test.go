package encode

import (
	"reflect"
	"testing"
)

var parseCases = []struct {
	arg      []string
	expected map[string]string
}{
	{[]string{"encode", "Hello, World", "0001"},
		map[string]string{"id": "0001", "task": "encode", "text": "Hello, World"}},
	{[]string{"encode", "Foo bar", "0002"},
		map[string]string{"id": "0002", "task": "encode", "text": "Foo bar"}},
	{[]string{"encode", "golang", "0003"},
		map[string]string{"id": "0003", "task": "encode", "text": "golang"}},
	{[]string{"encode", "", "0004"},
		map[string]string{"id": "0004", "task": "encode", "text": ""}},
}

func TestParse(t *testing.T) {
	for _, test := range parseCases {
		obs := parse(test.arg)
		equality := reflect.DeepEqual(obs, test.expected)
		if !equality {
			t.Errorf("failed to parse %v. expected %v, got %v", test.arg, test.expected, obs)
		}
	}
}
