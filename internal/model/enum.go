package model

import (
	"fmt"
	"io"
	"strconv"
)

const (
	UserTypeAdmin UserType = "ADMIN"
	UserTypeGuest UserType = "GUEST"
)

type UserType string

func (t UserType) IsValid() bool {
	switch t {
	case UserTypeAdmin, UserTypeGuest:
		return true
	default:
		return false
	}
}

func (t UserType) String() string {
	return string(t)
}

func (t UserType) MarshalGQL(w io.Writer) {
	_, _ = fmt.Fprint(w, strconv.Quote(t.String()))
}

func (t *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}
	*t = UserType(str)
	if !t.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}

	return nil
}
