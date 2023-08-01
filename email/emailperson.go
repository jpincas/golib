package email

import (
	"regexp"
	"strings"

	"github.com/yagniltd/golib/str"
)

// Note: i don't know how to not capture the enclosing <> and "",
// so i just have to trim them after capture
var (
	emailPersonEmailRegex = regexp.MustCompile(`<(.+)>`)
	emailPersonNameRegex  = regexp.MustCompile(`"(.+)"`)
)

// EmailPerson could be a sender or recipient
type EmailPerson struct {
	Name  string          `json:"name" bson:"name"`
	Email str.CleanString `json:"email" bson:"email"`
}

type MainContact = EmailPerson
type BillingContact = EmailPerson

type EmailPersons []EmailPerson

func EmailPersonParse(s string) EmailPerson {
	ep := EmailPerson{
		Name:  strings.Trim(emailPersonNameRegex.FindString(s), `"`),
		Email: str.NewCleanString(strings.Trim(emailPersonEmailRegex.FindString(s), `<>`)),
	}

	return ep
}

func EmailPersonsParse(ss []string) (res EmailPersons) {
	for _, s := range ss {
		if ep := EmailPersonParse(s); !ep.IsBlank() {
			res = append(res, ep)
		}
	}

	return
}

func singleRecipientEmailPerson(ep EmailPerson) EmailPersons {
	return EmailPersons{ep}
}

func singleRecipient(name string, email str.CleanString) EmailPersons {
	return EmailPerson{name, email}.single()
}

func singleRecipientWithNoName(email str.CleanString) EmailPersons {
	return singleEmailPersonWithNoName(email).single()
}

func singleEmailPersonWithNoName(email str.CleanString) EmailPerson {
	return EmailPerson{
		Email: email,
	}
}

func noCCs() EmailPersons {
	return EmailPersons{}
}

func (ep EmailPerson) single() EmailPersons {
	return EmailPersons{ep}
}

type CCEmailList []str.CleanString

func (ccs CCEmailList) ToEmailPersonList() (emailPersons EmailPersons) {
	for _, cc := range ccs {
		emailPersons = append(emailPersons, singleEmailPersonWithNoName(cc))
	}

	return emailPersons
}

func (ep EmailPerson) IsBlank() bool {
	return ep.Email == ""
}

func (ep EmailPerson) FormatForSending() string {
	if ep.Email == "" {
		return ""
	}

	if ep.Name == "" {
		return ep.Email.String()
	}

	return FormatNameAndEmailForSending(ep.Name, ep.Email)
}

func (eps EmailPersons) FormatForSending() string {
	l := []string{}
	for _, ep := range eps {
		if f := ep.FormatForSending(); f != "" {
			l = append(l, f)
		}
	}

	if len(l) == 0 {
		return ""
	}

	// The space after the comma is optional but helps improve readability
	return strings.Join(l, ", ")
}

func (eps EmailPersons) toAddresses() (out []str.CleanString) {
	for _, ep := range eps {
		out = append(out, ep.Email)
	}

	return
}
