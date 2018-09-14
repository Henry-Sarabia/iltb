package main

import (
	"encoding/json"
	"io"
	"os"
)

type class struct {
	Name    string   `json:"name"`
	Format  string   `json:"format"`
	Example string   `json:"example"`
	Verb    []string `json:"verb"`
}

func (c class) randomVerb() (string, error) {
	return randomString(c.Verb)
}

type classes map[string]class

func loadClasses(filename string) (classes, error) {
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

func readClasses(r io.Reader) (classes, error) {
	c := []class{}
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return nil, err
	}

	cs := classes{}
	for _, v := range c {
		cs[v.Name] = v
	}

	return cs, nil
}
