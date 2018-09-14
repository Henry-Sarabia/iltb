package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

type category struct {
	Name  string   `json:"name"`
	Types []string `json:"types"`
}

func (c category) random() (string, error) {
	t, err := randomString(c.Types)
	if err != nil {
		return "", errors.Wrap(err, "types slice is empty")
	}
	return t, nil
}

func loadCategories(filename string) (map[string]category, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readCategories(f)
}

func readCategories(r io.Reader) (map[string]category, error) {
	var c []category

	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return nil, err
	}

	cs := make(map[string]category)
	for _, v := range c {
		cs[v.Name] = v
	}

	return cs, nil
}
