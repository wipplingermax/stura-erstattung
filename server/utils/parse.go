package utils

import (
	"fmt"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
)

// format: not empty
func ParseFirstName(s *string) (err error) {

	p := strings.TrimSpace(*s)

	if p == "" {
		return fmt.Errorf("Firstname invalid")
	}

	s = &p
	return nil
}

// format: not empty
func ParseLastName(s *string) (err error) {

	p := strings.TrimSpace(*s)

	if p == "" {
		return fmt.Errorf("Lastname invalid")
	}

	s = &p
	return nil
}

// format: 7 digit number
func ParseMatriculationNumber(s *string) (err error) {

	p := strings.ReplaceAll(*s, " ", "")

	if _, err := regexp.MatchString("/^[0-9]{7}$/", p); err != nil {
		return fmt.Errorf("Matriculationnumber invalid")
	}

	s = &p
	return nil
}

// format: 2 letter, 3 numbers - like "ab123"
func ParseUniID(s *string) (err error) {

	p := strings.ToLower(strings.ReplaceAll(*s, " ", ""))

	if val, _ := regexp.MatchString("/^[a-z]{2}[0-9]{3}$/", p); !val {
		return fmt.Errorf("Uni-ID invalid")
	}

	s = &p
	return nil
}

func ParseEmail(s *string) (err error) {

	p := *s
	var m *mail.Address

	if m, err = mail.ParseAddress(p); err != nil {
		return fmt.Errorf("E-Mail invalid")
	}

	p = m.String()
	s = &p
	return nil
}

// format: no spaces
func ParsePhone(s *string) (err error) {

	p := strings.ReplaceAll(*s, " ", "")

	s = &p
	return nil
}

// format: Uppercase 22 digit
func ParseIBAN(s *string) (err error) {

	p := strings.ToUpper(strings.ReplaceAll(*s, " ", ""))

	// german IBAN
	if val, _ := regexp.MatchString("/^DE/", p); !val {
		return fmt.Errorf("IBAN invalid - german IBAN needed")
	}

	// 22 digits
	if val, _ := regexp.MatchString("/^DE[0-9]{20}$/", p); !val {
		return fmt.Errorf("IBAN invalid")
	}

	// validate
	// 1. move the four initial characters to the end of the string
	// 2. replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11 (DE = 1314)
	// 3. treat the account number as a large integer, divide it by 97 and use the remainder for check
	// a valid Iban is indicated by a remainder value of 1

	var checksum string
	checksum = p[4:] + "1314" + p[2:3]
	if res, _ := strconv.Atoi(checksum); res%97 != 1 {
		return fmt.Errorf("Request IBAN invalid")
	}

	s = &p
	return nil
}

//format:
func ParseBIC(s *string) (err error) {

	p := strings.ToUpper(strings.ReplaceAll(*s, " ", ""))

	if val, _ := regexp.MatchString("/^[0-9A-Z]{8:11}$/", p); !val {
		return fmt.Errorf("Bic invalid")
	}

	s = &p
	return nil
}

// format: not empty
func ParseAccountOwner(s *string) (err error) {

	p := strings.TrimSpace(*s)

	if p == "" {
		return fmt.Errorf("Lastname invalid")
	}

	s = &p
	return nil
}
