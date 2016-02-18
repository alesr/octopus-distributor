package encode

import "golang.org/x/crypto/bcrypt"

func Exec(request []string, resultCh chan map[string]string) {

	enc := parse(request)
	encodeText(enc)

	// fmt.Println(enc)

	resultCh <- enc
}

// Parse the request
func parse(e []string) map[string]string {

	enc := make(map[string]string)
	enc["id"] = e[2]
	enc["task"] = e[0]
	enc["text"] = e[1]

	return enc
}

func encodeText(enc map[string]string) {

	hashedTxt, err := bcrypt.GenerateFromPassword([]byte(enc["text"]), 4)
	if err != nil {
		enc["result"] = err.Error()
		return
	}
	enc["result"] = string(hashedTxt)
}
