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

func runReverse(revCh chan reverse) {
	for {
		select {
		case r := <-revCh:
			r.parser()
			r.reverseText()
			fmt.Println(r)
		}
	}
}

// Parse the request as Reverse struct
func (rev *reverse) parser() {
	rev.text = rev.request[1]
}

func (rev *reverse) reverseText() {

	strLen := len(rev.text)

	// The reverse of a empty string is a empty string
	if strLen == 0 {
		rev.result = rev.text
	}

	// Same above
	if strLen == 1 {
		rev.result = rev.text
	}

	// Convert s into unicode points
	s := []rune(rev.text)

	// Last index
	rLen := len(s) - 1

	// String new home
	reverse := []string{}

	for i := rLen; i >= 0; i-- {
		reverse = append(reverse, string(s[i]))
	}
	rev.result = strings.Join(reverse, "")
}

func (rev reverse) String() string {
	return fmt.Sprintf("id %d   %s   ->    %s", rev.id, rev.text, rev.result)
}
