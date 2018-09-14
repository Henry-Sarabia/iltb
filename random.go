package main

import (
	"math/rand"

	"github.com/pkg/errors"
)

func randomString(s []string) (string, error) {
	if len(s) < 1 {
		return "", errors.New("string slice is empty")
	}

	r := rand.Intn(len(s))
	return s[r], nil
}

type randomizer interface {
	random() interface{}
}
