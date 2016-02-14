package subscriber

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Encode request data
type encode struct {
	request      []string
	id           int
	text, result string
	err          error
}

func runEncode(revCh chan encode) {
	for {
		select {
		case e := <-encCh:
			e.parser()
			e.encodeText()
			fmt.Println(e)
		}
	}
}

// Parse the request as Encode struct
func (enc *encode) parser() error {
	enc.text = enc.request[1]
	return nil
}

func (enc *encode) encodeText() {

	hashedTxt, err := bcrypt.GenerateFromPassword([]byte(enc.text), 4)
	if err != nil {
		enc.err = err
		return
	}
	enc.result = string(hashedTxt)
}

func (enc encode) String() string {
	return fmt.Sprintf("id %d   %s   ->    %s", enc.id, enc.text, enc.result)
}
