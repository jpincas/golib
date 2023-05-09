package postcode

import (
	"errors"
	"regexp"
	"strings"
)

// https://en.wikipedia.org/wiki/Postcodes_in_the_United_Kingdom#Formatting

const (
	UKPostcodeRegex                = `^[A-Z]{1,2}[0-9][A-Z0-9]? ?[0-9][A-Z]{2}$`
	ErrorIncorrectUKPostcodeFormat = "incorrect format for UK postcode"
)

var validUKPostcode = regexp.MustCompile(UKPostcodeRegex)

type UKPostcode struct {
	OutwardCode string
	InwardCode  string
}

func ParseUKPostcode(s string) (UKPostcode, error) {
	// Remove all whitespace
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "")

	// Validate
	if isValid := validUKPostcode.MatchString(s); !isValid {
		return UKPostcode{}, errors.New(ErrorIncorrectUKPostcodeFormat)
	}

	// Let's make double sure we don't go overbounds
	// Even though the regex should always reject any postcode that is too short
	if len(s) < 3 {
		return UKPostcode{}, errors.New(ErrorIncorrectUKPostcodeFormat)
	}

	// As all formats end with 9AA, the first part of a postcode can easily be extracted by ignoring the last three characters.
	inward := s[len(s)-3:]
	outward := strings.TrimSuffix(s, inward)
	return UKPostcode{
		OutwardCode: outward,
		InwardCode:  inward,
	}, nil
}
