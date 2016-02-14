package subscriber

import (
	"fmt"
	"strings"
)

func reverser(revCh chan Reverse) {
	for {
		select {
		case r := <-revCh:
			go func(revCh chan Reverse) {
				r.result = stringReverse(r.text)
				fmt.Println(r)
			}(revCh)
		}

	}
}

func stringReverse(s string) string {

	strLen := len(s)

	// The reverse of a empty string is a empty string
	if strLen == 0 {
		return s
	}

	// Same above
	if strLen == 1 {
		return s
	}

	// Convert s into unicode points
	r := []rune(s)

	// Last index
	rLen := len(r) - 1

	// String new home
	rev := []string{}

	for i := rLen; i >= 0; i-- {
		rev = append(rev, string(r[i]))
	}

	return strings.Join(rev, "")
}
