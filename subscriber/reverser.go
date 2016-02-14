package subscriber

import (
	"fmt"
	"strings"
)

// reverse request data
type reverse struct {
	request      []string
	id           int
	text, result string
	err          error
}

func reverser(revCh chan reverse) {
	for {
		select {
		case r := <-revCh:
			go func(revCh chan reverse) {

				if err := r.parse(); err != nil {
					r.err = err
					return
				}
				r.stringReverse()
				fmt.Println(r)
			}(revCh)
		}
	}
}

// Parse the request as Reverse struct
func (r *reverse) parse() error {
	r.text = r.request[1]
	return nil
}

func (r *reverse) stringReverse() {

	strLen := len(r.text)

	// The reverse of a empty string is a empty string
	if strLen == 0 {
		r.result = r.text
	}

	// Same above
	if strLen == 1 {
		r.result = r.text
	}

	// Convert s into unicode points
	s := []rune(r.text)

	// Last index
	rLen := len(s) - 1

	// String new home
	rev := []string{}

	for i := rLen; i >= 0; i-- {
		rev = append(rev, string(s[i]))
	}

	r.result = strings.Join(rev, "")
}
