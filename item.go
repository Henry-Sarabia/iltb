package main

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"os"
)

type item struct {
	Base      string   `json:"base"`
	Class     string   `json:"class"`
	Materials []string `json:"materials"`
	Contents  []string `json:"contents"`
}

func (i item) chooseMaterial() (string, error) {
	return randomString(i.Materials)
}

func (i item) chooseContent() (string, error) {
	return randomString(i.Contents)
}

type items []item

func loadItems(filename string) (items, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readItems(f)
}

func readItems(r io.Reader) (items, error) {
	i := items{}

	if err := json.NewDecoder(r).Decode(&i); err != nil {
		return nil, err
	}

	return i, nil
}

func (i items) choose() (item, error) {
	if len(i) < 1 {
		return item{}, errors.New("items list cannot be empty")
	}

	r := rand.Intn(len(i))
	return i[r], nil
}
