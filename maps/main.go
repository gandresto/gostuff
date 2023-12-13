package main

import (
	"errors"
)

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func deleteIfNeccessary(users map[string]user, name string) (deleted bool, err error) {
	usr, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if usr.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

type stringOrNumberOrBoolean interface {
	string | float64 | int | bool
}

func main() {
	users := map[string]user{
		"hello": {
			name:                 "se",
			number:               54648123,
			scheduledForDeletion: false,
		},
	}
	deleteIfNeccessary(users, "hello")
}
