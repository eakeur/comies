package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPassword(t *testing.T) {
	t.Parallel()

	type (
		test struct {
			name    string
			from    string
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should create encrypted password",
			from: "1922Eakeur!(@@",
		},
		{
			name:    "should fail because its missing number",
			from:    "AAAAEakeur!(@@",
			wantErr: ErrInvalidPassword,
		},
		{
			name:    "should fail because its missing symbols",
			from:    "1922Eakeur1922",
			wantErr: ErrInvalidPassword,
		},
		{
			name:    "should fail because its missing upper case letters",
			from:    "1922eakeur!(@@",
			wantErr: ErrInvalidPassword,
		},
		{
			name:    "should fail because its missing lower case letters",
			from:    "1922EAKEUR!(@@",
			wantErr: ErrInvalidPassword,
		},
		{
			name:    "should fail because its shorter than wanted",
			from:    "19Ea@!",
			wantErr: ErrInvalidPassword,
		},
		{
			name:    "should fail because its longer than wanted",
			from:    "1922Eakeur!(@@@@@@@@",
			wantErr: ErrInvalidPassword,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			pwd, err := NewPassword(c.from)
			assert.ErrorIs(t, err, c.wantErr)

			if err == nil {
				err := pwd.Compare(c.from)
				assert.ErrorIs(t, err, c.wantErr)
			}

		})
	}
}

func TestSetPassword(t *testing.T) {
	t.Parallel()

	const pass = "1922Eakeur!(@@"

	placeholder := struct {
		password Password
	}{}

	err := SetPassword(pass, &placeholder.password)
	assert.ErrorIs(t, err, nil)

	err = placeholder.password.Compare(pass)
	assert.ErrorIs(t, err, nil)
}

func TestPassword_Compare(t *testing.T) {
	t.Parallel()

	type (
		test struct {
			name    string
			from    string
			compare string
			wantErr error
		}
	)

	cases := []test{
		{
			name:    "should return nil error for right password",
			from:    "1922Eakeur!(@@",
			compare: "1922Eakeur!(@@",
		},
		{
			name:    "should return ErrWrongPassword for wrong password",
			from:    "1922Eakeur!(@@",
			compare: "1922Eakeur!!!!",
			wantErr: ErrWrongPassword,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			pwd, _ := NewPassword(c.from)
			assert.ErrorIs(t, pwd.Compare(c.compare), c.wantErr)

		})
	}
}
