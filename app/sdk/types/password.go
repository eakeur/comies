package types

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type Password string

const encryptingCost = 7

var (
	ErrInvalidPassword = errors.New("the password must have minimum eight and maximum 16 characters, at least one uppercase letter, one lowercase letter, one number and one special character")
	ErrWrongPassword   = errors.New("the password provided is different from the stored one")
	passwordTests      = []*regexp.Regexp{
		regexp.MustCompile("\\d"),
		regexp.MustCompile("[a-z]"),
		regexp.MustCompile("[A-Z]"),
		regexp.MustCompile("[@#$%^&+=*]"),
		regexp.MustCompile("^.{8,16}$"),
	}
)

func NewPassword(from string) (Password, error) {

	for _, test := range passwordTests {
		if !test.MatchString(from) {
			return "", ErrInvalidPassword
		}
	}

	var arr []byte
	arr, err := bcrypt.GenerateFromPassword([]byte(from), encryptingCost)
	if err != nil {
		return "", err
	}

	return Password(arr), nil

}

func MustCreatePassword(from string) Password {
	pwd, err := NewPassword(from)
	if err != nil {
		panic(err)
	}

	return pwd
}

func SetPassword(from string, password *Password) error {
	pwd, err := NewPassword(from)
	if err != nil {
		return err
	}
	*password = pwd

	return nil
}

func (p Password) Compare(input string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(input))
	if err != nil {
		return ErrWrongPassword
	}

	return nil
}
