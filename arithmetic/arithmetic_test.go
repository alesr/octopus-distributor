package arithmetic

import (
	"reflect"
	"testing"
)

var parseCases = []struct {
	arg      []string
	expected map[string]string
}{
	{[]string{"add", "1", "2", "0001"},
		map[string]string{"id": "0001", "task": "add", "a": "1", "b": "2"}},
	{[]string{"sub", "3", "4", "0002"},
		map[string]string{"id": "0002", "task": "sub", "a": "3", "b": "4"}},
	{[]string{"mult", "5", "6", "0003"},
		map[string]string{"id": "0003", "task": "mult", "a": "5", "b": "6"}},
	{[]string{"div", "7", "8", "0004"},
		map[string]string{"id": "0004", "task": "div", "a": "7", "b": "8"}},
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

var addCases = []struct {
	arg, expected map[string]string
}{
	{map[string]string{"id": "0001", "task": "add", "a": "1", "b": "2"},
		map[string]string{"id": "0001", "task": "add", "a": "1", "b": "2", "result": "3"}},
	{map[string]string{"id": "0002", "task": "add", "a": "3", "b": "4"},
		map[string]string{"id": "0002", "task": "add", "a": "3", "b": "4", "result": "7"}},
	{map[string]string{"id": "0003", "task": "add", "a": "5", "b": "6"},
		map[string]string{"id": "0003", "task": "add", "a": "5", "b": "6", "result": "11"}},
	{map[string]string{"id": "0004", "task": "add", "a": "7", "b": "8"},
		map[string]string{"id": "0004", "task": "add", "a": "7", "b": "8", "result": "15"}},
}

func TestAdd(t *testing.T) {
	for _, test := range addCases {
		add(test.arg)
		if test.arg["result"] != test.expected["result"] {
			t.Errorf("failed to add %v. expected %s, got %s", test.arg, test.expected["result"], test.arg["result"])
		}
	}
}

var subractCases = []struct {
	arg, expected map[string]string
}{
	{map[string]string{"id": "0001", "task": "sub", "a": "1", "b": "2"},
		map[string]string{"id": "0001", "task": "sub", "a": "1", "b": "2", "result": "-1"}},
	{map[string]string{"id": "0002", "task": "sub", "a": "3", "b": "5"},
		map[string]string{"id": "0002", "task": "sub", "a": "3", "b": "5", "result": "-2"}},
	{map[string]string{"id": "0003", "task": "sub", "a": "15", "b": "6"},
		map[string]string{"id": "0003", "task": "sub", "a": "15", "b": "6", "result": "9"}},
	{map[string]string{"id": "0004", "task": "sub", "a": "10", "b": "5"},
		map[string]string{"id": "0004", "task": "sub", "a": "10", "b": "5", "result": "5"}},
}

func TestSubtract(t *testing.T) {
	for _, test := range subractCases {
		subtract(test.arg)
		if test.arg["result"] != test.expected["result"] {
			t.Errorf("failed to subtract %v. expected %s, got %s", test.arg, test.expected["result"], test.arg["result"])
		}
	}
}

var multiplyCases = []struct {
	arg, expected map[string]string
}{
	{map[string]string{"id": "0001", "task": "mult", "a": "0", "b": "2"},
		map[string]string{"id": "0001", "task": "mult", "a": "0", "b": "2", "result": "0"}},
	{map[string]string{"id": "0002", "task": "mult", "a": "3", "b": "5"},
		map[string]string{"id": "0002", "task": "mult", "a": "3", "b": "5", "result": "15"}},
	{map[string]string{"id": "0003", "task": "mult", "a": "6", "b": "6"},
		map[string]string{"id": "0003", "task": "mult", "a": "6", "b": "6", "result": "36"}},
	{map[string]string{"id": "0004", "task": "mult", "a": "10", "b": "5"},
		map[string]string{"id": "0004", "task": "mult", "a": "10", "b": "5", "result": "50"}},
}

func TestMultiply(t *testing.T) {
	for _, test := range multiplyCases {
		multiply(test.arg)
		if test.arg["result"] != test.expected["result"] {
			t.Errorf("failed to multiply %v. expected %s, got %s", test.arg, test.expected["result"], test.arg["result"])
		}
	}
}

var divideCases = []struct {
	arg, expected map[string]string
}{
	{map[string]string{"id": "0001", "task": "div", "a": "0", "b": "0"},
		map[string]string{"id": "0001", "task": "div", "a": "0", "b": "0", "result": errZeroDivError.Error()}},
	{map[string]string{"id": "0002", "task": "div", "a": "4", "b": "2"},
		map[string]string{"id": "0002", "task": "div", "a": "4", "b": "2", "result": "2.0"}},
	{map[string]string{"id": "0003", "task": "div", "a": "6", "b": "6"},
		map[string]string{"id": "0003", "task": "div", "a": "6", "b": "6", "result": "1.0"}},
	{map[string]string{"id": "0004", "task": "div", "a": "10", "b": "8"},
		map[string]string{"id": "0004", "task": "div", "a": "10", "b": "8", "result": "1.2"}},
}

func TestDivide(t *testing.T) {
	for _, test := range divideCases {
		divide(test.arg)
		if test.arg["result"] != test.expected["result"] {
			t.Errorf("failed to multiply %v. expected %s, got %s", test.arg, test.expected["result"], test.arg["result"])
		}
	}
}
