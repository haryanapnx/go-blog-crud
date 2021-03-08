package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	var msg string = "already taken"

	if strings.Contains(err, "username") {
		return errors.New("Username" + msg)
	}

	if strings.Contains(err, "email") {
		return errors.New("Email" + msg)
	}

	if strings.Contains(err, "title") {
		return errors.New("Title" + msg)
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorect Password")
	}
	return errors.New("Incorect Detail")
}
