package generator

import "testing"

var randomErrorCases = []struct {
	arg int
	err error
}{
	{-2, errRandomFuncError},
	{-1, errRandomFuncError},
	{0, errRandomFuncError},
	{1, nil},
	{2, nil},
}

// not possibile to test randoms, so at least let's check if error is working
func randomTest(t *testing.T) {
	for _, test := range randomErrorCases {
		_, err := random(test.arg)
		if err != nil {
			if err != test.err {
				t.Error(err)
			}
		}
	}
}
