package arithmetic

import (
	"reflect"
	"testing"
)

var parseCases = []struct {
	arg      []string
	expected map[string]string
}{
	{[]string{"add", "1", "2", "0001"}, map[string]string{"id": "0001", "task": "add", "a": "1", "b": "2"}},
	{[]string{"sub", "3", "4", "0002"}, map[string]string{"id": "0002", "task": "sub", "a": "3", "b": "4"}},
	{[]string{"mult", "5", "6", "0003"}, map[string]string{"id": "0003", "task": "mult", "a": "5", "b": "6"}},
	{[]string{"div", "7", "8", "0004"}, map[string]string{"id": "0004", "task": "div", "a": "7", "b": "8"}},
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
