package main

import (
	"fmt"
	"io"
	"math/rand"
	"strings"

	"github.com/Henry-Sarabia/article"
	"github.com/pkg/errors"
)

// Factory generates items
type Factory struct {
	recipeList  []recipe
	materialMap map[string]category
	contentMap  map[string]category
	classMap    map[string]class
}

func (f *Factory) randomRecipe() (recipe, error) {
	if len(f.recipeList) < 1 {
		return recipe{}, errors.New("recipe slice is empty")
	}

	i := rand.Intn(len(f.recipeList))
	return f.recipeList[i], nil
}

// Item generates a new item
func (f *Factory) Item() (string, error) {
	r, err := f.randomRecipe()
	if err != nil {
		return "", err
	}

	m, err := r.material()
	if err != nil {
		return "", errors.Wrapf(err, "cannot retrieve material from recipe '%v'", r.Base)
	}
	mat, ok := f.materialMap[m]
	if !ok {
		return "", errors.Errorf("cannot find material '%v' in material map", m)
	}

	c, err := r.content()
	if err != nil {
		return "", errors.Wrapf(err, "cannot retrieve content from recipe '%v'", r.Base)
	}
	cont, ok := f.contentMap[c]
	if !ok {
		return "", errors.Errorf("cannot find content '%v' in content map", c)
	}

	cl, ok := f.classMap[r.Class]
	if !ok {
		return "", errors.Errorf("cannot find class '%v'", r.Class)
	}

	return compose(r.Base, mat, cont, cl)
}

func compose(base string, mat category, cont category, cl class) (string, error) {
	tok, err := tokenize(cl.Format)
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
				return "", errors.Wrap(err, "cannot retrieve random material type")
			}
			tok[i] = m

		case "<base>":
			tok[i] = base

		case "<verb>":
			v, err := cl.randomVerb()
			if err != nil {
				return "", errors.Wrap(err, "cannot retrieve random verb")
			}
			tok[i] = v

		case "<content>":
			c, err := cont.random()
			if err != nil {
				return "", errors.Wrap(err, "cannot retrieve random content type")
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
func New(r, m, c, cl io.Reader) (*Factory, error) {
	var err error
	f := &Factory{}
	f.recipeList, err = readRecipes(r)
	if err != nil {
		return nil, err
	}

	f.materialMap, err = readCategories(m)
	if err != nil {
		return nil, err
	}

	f.contentMap, err = readCategories(c)
	if err != nil {
		return nil, err
	}

	f.classMap, err = readClasses(cl)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Load returns a new Factory with the given filenames loaded into their
// appropriate fields.
func Load(items, materialMap, contentMap, classMap string) (*Factory, error) {
	i, err := loadRecipes(items)
	if err != nil {
		return nil, err
	}

	m, err := loadCategories(materialMap)
	if err != nil {
		return nil, err
	}

	c, err := loadCategories(contentMap)
	if err != nil {
		return nil, err
	}

	cl, err := loadClasses(classMap)
	if err != nil {
		return nil, err
	}

	return &Factory{
		recipeList:  i,
		materialMap: m,
		contentMap:  c,
		classMap:    cl,
	}, nil
}
