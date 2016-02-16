package reverse

import (
	"reflect"
	"testing"
)

var parseCases = []struct {
	arg      []string
	expected map[string]string
}{
	{[]string{"reverse", "Hello, World", "0001"},
		map[string]string{"id": "0001", "task": "reverse", "text": "Hello, World"}},
	{[]string{"reverse", "Foo bar", "0002"},
		map[string]string{"id": "0002", "task": "reverse", "text": "Foo bar"}},
	{[]string{"reverse", "golang", "0003"},
		map[string]string{"id": "0003", "task": "reverse", "text": "golang"}},
	{[]string{"reverse", "", "0004"},
		map[string]string{"id": "0004", "task": "reverse", "text": ""}},
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

var revCases = []struct {
	arg, expected map[string]string
}{
	{map[string]string{"id": "0001", "task": "reverse", "text": "Hello, World"},
		map[string]string{"id": "0001", "task": "reverse", "text": "Hello, World", "result": "dlroW ,olleH"}},
	{map[string]string{"id": "0002", "task": "reverse", "text": "Foo bar"},
		map[string]string{"id": "0002", "task": "reverse", "text": "Foo bar", "result": "rab ooF"}},
	{map[string]string{"id": "0003", "task": "reverse", "text": "golang"},
		map[string]string{"id": "0003", "task": "reverse", "text": "golang", "result": "gnalog"}},
	{map[string]string{"id": "0004", "task": "reverse", "text": ""},
		map[string]string{"id": "0004", "task": "reverse", "text": "", "result": ""}},
}

func TestReverseText(t *testing.T) {
	for _, test := range revCases {
		reverseText(test.arg)
		if test.arg["result"] != test.expected["result"] {
			t.Errorf("failed to reverse string %v. expected %s, got %s", test.arg, test.expected["result"], test.arg["result"])
		}
	}
}
