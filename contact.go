// Copyright 2016 Jesse Allen.

package main

import (
	"errors"
	"strings"
)

type Contact struct {
	Name    string `json:"name"`
	Address struct {
		StreetAddress string `json:"streetAddress"`
		Locality      string `json:"locality"`
		Region        string `json:"region"`
		PostalCode    string `json:"postalCode"`
	} `json:"address"`
	Phone []Phone `json:"phone"`
	Email string  `json:"email"`
}

type Phone struct {
	Type      PhoneType   `json:"type"`
	Number    PhoneNumber `json:"number"`
	Extension string      `json:"extension"`
	Primary   bool        `json:"primary"`
}

// Use to select only a single, primary phone number. Selects the first phone listed as primary.
func (c Contact) PrimaryPhone() *Phone {
	var p Phone
	for _, p = range c.Phone {
		if p.Primary {
			return &p
		}
	}
	return nil
}

type PhoneType int

//go:generate stringer -type=PhoneType
//go:generate jsonenums -type=PhoneType

const (
	unknown PhoneType = iota
	mobile
	home
	office
	other
)

func (t PhoneType) Short() string {
	typeMap := map[PhoneType]string{
		unknown: "",
		mobile:  "m",
		home:    "h",
		office:  "o",
		other:   "",
	}
	return typeMap[t]
}

func (t PhoneType) Title() string {
	return strings.ToTitle(t.String())
}

// Phone Number (for now, expects 11, 10, or 7 digit number)
type PhoneNumber string

var ERROR_NOT_NANP_NUMBER = errors.New("Phone Number is not a valid North American Numbering Plan phone number.")

// Removes non-numeric chars from phone number.
// If the resulting length is not 11, 10, or 7 digits, returns ERROR_NOT_NANP_NUMBER.
func (n PhoneNumber) cleanPhoneNumber() (PhoneNumber, error) {
	m := strings.Replace(strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return 'x'
	}, string(n)), "x", "", -1)
	if len(m) != 11 && len(m) != 10 && len(m) != 7 {
		return n, ERROR_NOT_NANP_NUMBER
	}
	return PhoneNumber(m), nil
}

// Format phone number as NANP (North American Numbering Plan) "1 (NPA) NXX-XXXX"
// If the number is not a valid NANP number, the original number is returned unformatted.
func (n PhoneNumber) FormatTraditional() string {
	m, err := n.cleanPhoneNumber()
	if err != nil { // this includes ERROR_NOT_NANP_NUMBER
		return string(n)
	}
	b := []byte(m) // make a copy to parse. We can do this by byte because 0-9 are represented as single bytes
	s := ""
	if len(b) == 11 {
		s += string(b[0])
		b = b[1:]
	}
	if len(b) == 10 {
		s += "(" + string(b[0:3]) + ") "
		b = b[3:]
	}
	if len(b) == 7 {
		s += string(b[0:3]) + "-" + string(b[3:])
	}
	return s
}

// Format phone number as NANP (North American Numbering Plan) "1-NPA-NXX-XXXX"
// If the number is not a valid NANP number, the original number is returned unformatted.
func (n PhoneNumber) FormatSeparator(sep string) string {
	m, err := n.cleanPhoneNumber()
	if err != nil { // this includes ERROR_NOT_NANP_NUMBER
		return string(n)
	}
	b := []byte(m) // make a copy to parse. We can do this by byte because 0-9 are represented as single bytes
	s := ""
	if len(b) == 11 {
		s += string(b[0]) + sep
		b = b[1:]
	}
	if len(b) == 10 {
		s += string(b[0:3]) + sep
		b = b[3:]
	}
	if len(b) == 7 {
		s += string(b[0:3]) + sep + string(b[3:])
	}
	return s
}
