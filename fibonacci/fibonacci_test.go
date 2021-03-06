package fibonacci

import (
	"reflect"
	"testing"
)

var parseCases = []struct {
	arg      []string
	expected map[string]string
}{
	{[]string{"fibonacci", "5", "0001"},
		map[string]string{"id": "0001", "task": "fibonacci", "n": "5"}},
	{[]string{"fibonacci", "10", "0002"},
		map[string]string{"id": "0002", "task": "fibonacci", "n": "10"}},
	{[]string{"fibonacci", "4", "0003"},
		map[string]string{"id": "0003", "task": "fibonacci", "n": "4"}},
	{[]string{"fibonacci", "27", "0004"},
		map[string]string{"id": "0004", "task": "fibonacci", "n": "27"}},
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

var fibCases = []struct {
	arg, expected int
}{
	{0, 0},
	{1, 1},
	{10, 55},
	{27, 196418},
}

func TestNthFibonacci(t *testing.T) {
	for _, test := range fibCases {
		obs := nthFibonacci(test.arg)
		if obs != test.expected {
			t.Errorf("failed to calculate fibonacci %v. expected %s, got %s", test.arg, test.expected, obs)
		}
	}
}
