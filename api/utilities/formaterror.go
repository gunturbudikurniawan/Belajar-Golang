package utilities

import (
	"errors"
	"strings"
)

// var errorMessages = make(map[string]string)

// var err error

//FormatError hahhahhahahhah
func FormatError(errString string) error {

	if strings.Contains(errString, "username") {
		return errors.New("Username Already Taken")
	}

	if strings.Contains(errString, "email") {
		return errors.New("Email Already Taken")

	}
	if strings.Contains(errString, "phone") {
		return errors.New("Phone Already Taken")

	}
	if strings.Contains(errString, "hashedPassword") {
		return errors.New("Incorrect Password")

	}
	if strings.Contains(errString, "record not found") {
		return errors.New("No Record Found")

	}

	return errors.New("Incorrect Details")
}
