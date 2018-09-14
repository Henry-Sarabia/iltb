package main

import (
	"encoding/json"
	"io"
	"os"
)

type category struct {
	Name  string   `json:"name"`
	Types []string `json:"types"`
}

func (c category) random() (string, error) {
	return randomString(c.Types)
}

type categories map[string]category

func loadCategories(filename string) (categories, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readCategories(f)
}

func readCategories(r io.Reader) (categories, error) {
	c := []category{}

	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return nil, err
	}

	cs := categories{}
	for _, v := range c {
		cs[v.Name] = v
	}

	return cs, nil
}
