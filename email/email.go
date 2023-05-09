package email

import (
	"fmt"
	"regexp"

	"github.com/yagniltd/golib/str"
)

const (
	//EmailRegex = `.+@.+\..+`
	//EmailRegex = `/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/`
	EmailRegex = `(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`
)

var (
	compiledEmailRegex = regexp.MustCompile(EmailRegex)
)

func FormatNameAndEmailForSending(name string, email str.CleanString) string {
	// Wrapping in quotes is optional, but required if there are any punctuation
	// marks in the name, so we'll just always do it
	return fmt.Sprintf(`"%s" <%s>`, name, email.String())
}

func IsValidEmail(s string) bool {
	return compiledEmailRegex.MatchString(s)

}
