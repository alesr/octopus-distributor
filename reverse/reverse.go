package reverse

import (
	"fmt"
	"strings"
)

func Exec(revCh chan []string) {

	r := <-revCh

	rev := parse(r)
	reverseText(rev)
	fmt.Println(rev)
}

// Parse the request
func parse(r []string) map[string]string {

	rev := make(map[string]string)
	rev["id"] = r[2]
	rev["task"] = r[0]
	rev["text"] = r[1]

	return rev
}

func reverseText(rev map[string]string) {

	strLen := len(rev["text"])

	// The reverse of a empty string is a empty string
	if strLen == 0 || strLen == 1 {
		rev["result"] = rev["text"]
	}

	// Convert s into unicode points
	s := []rune(rev["text"])

	// Last index
	rLen := len(s) - 1

	// String new home
	reverse := []string{}

	for i := rLen; i >= 0; i-- {
		reverse = append(reverse, string(s[i]))
	}
	rev["result"] = strings.Join(reverse, "")
}
