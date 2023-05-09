package web

import "strings"

func StripPort(host string) string {
	s := strings.Split(host, ":")
	return s[0]
}
