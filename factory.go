package main

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/Henry-Sarabia/article"
)

// Item is implemented by any type that can compose an item's properties into
// a proper output string.
type Item interface {
	Compose() string
}

// Factory generates items
type Factory struct {
	items     items
	materials categories
	contents  categories
	classes   classes
}

func (f Factory) chooseItem() (item, error) {
	return f.items.choose()
}

// Item generates a new item
func (f Factory) Item() (string, error) {
	item, err := f.chooseItem()
	if err != nil {
		return "", err
	}

	m, err := item.chooseMaterial()
	if err != nil {
		return "", err
	}
	material := f.materials[m]

	c, err := item.chooseContent()
	if err != nil {
		return "", err
	}
	content := f.contents[c]

	class := f.classes[item.Class]

	return parse(item.Base, material, content, class)
}

func parse(base string, mat category, cont category, class class) (string, error) {
	tok, err := tokenize(class.Format)
	if err != nil {
		return "", err
	}
	for i := len(tok) - 1; i >= 0; i-- {
		switch tok[i] {
		case "<article>":
			tok[i] = article.Indefinite(tok[i+1])
		case "<material>":
			m, err := mat.random()
			if err != nil {
				return "", err
			}
			tok[i] = m
		case "<base>":
			tok[i] = base
		case "<verb>":
			v, err := class.randomVerb()
			if err != nil {
				return "", err
			}
			tok[i] = v
		case "<content>":
			c, err := cont.random()
			if err != nil {
				return "", err
			}
			tok[i] = c
		default:
			return "", fmt.Errorf("unexpected token '%v' in format", tok[i])
		}
	}
	return strings.Join(tok, " "), nil
}

func tokenize(format string) ([]string, error) {
	tok := strings.Fields(format)
	if tok[len(tok)-1] == "<article>" {
		return nil, errors.New("article token cannot be the last token in a format")
	}
	return tok, nil
}

// New returns a new *Factory initialized with the given items, material
// categories, and content categories.
func New(i, m, c, cl io.Reader) (*Factory, error) {
	var err error
	f := &Factory{}
	f.items, err = readItems(i)
	if err != nil {
		return nil, err
	}

	f.materials, err = readCategories(m)
	if err != nil {
		return nil, err
	}

	f.contents, err = readCategories(c)
	if err != nil {
		return nil, err
	}

	f.classes, err = readClasses(cl)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Load returns a new Factory with the given filenames loaded into their
// appropriate fields.
func Load(items, materials, contents, classes string) (Factory, error) {
	i, err := loadItems(items)
	if err != nil {
		return Factory{}, nil
	}

	m, err := loadCategories(materials)
	if err != nil {
		return Factory{}, nil
	}

	c, err := loadCategories(contents)
	if err != nil {
		return Factory{}, nil
	}

	cl, err := loadClasses(classes)
	if err != nil {
		return Factory{}, nil
	}

	return Factory{
		items:     i,
		materials: m,
		contents:  c,
		classes:   cl,
	}, nil
}
