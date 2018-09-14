package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

type class struct {
	Name    string   `json:"name"`
	Format  string   `json:"format"`
	Example string   `json:"example"`
	Verb    []string `json:"verb"`
}

func (c class) randomVerb() (string, error) {
	v, err := randomString(c.Verb)
	if err != nil {
		return "", errors.Wrap(err, "verb slice is empty")
	}
	return v, err
}

func loadClasses(filename string) (map[string]class, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c, err := readClasses(f)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func readClasses(r io.Reader) (map[string]class, error) {
	var c []class
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return nil, err
	}

	cs := make(map[string]class)
	for _, v := range c {
		cs[v.Name] = v
	}

	return cs, nil
}
