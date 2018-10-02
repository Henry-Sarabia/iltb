package iltb

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

// recipe contains information for generating a base item from a given
// item class using one of the materials and details.
type recipe struct {
	Base            string   `json:"base"`
	Class           string   `json:"class"`
	Value           float64  `json:"value"`
	MaterialChoices []string `json:"materials"`
	ContentChoices  []string `json:"contents"`
}

func (r recipe) material() (string, error) {
	m, err := randomString(r.MaterialChoices)
	if err != nil {
		return "", errors.Wrap(err, "material slice is empty")
	}

	return m, nil
}

func (r recipe) content() (string, error) {
	c, err := randomString(r.ContentChoices)
	if err != nil {
		return "", errors.Wrap(err, "content slice is empty")
	}

	return c, nil
}

func loadRecipes(filename string) ([]recipe, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readRecipes(f)
}

func readRecipes(r io.Reader) ([]recipe, error) {
	var rcp []recipe

	if err := json.NewDecoder(r).Decode(&rcp); err != nil {
		return nil, err
	}

	return rcp, nil
}
