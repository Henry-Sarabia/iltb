package iltb

import (
	"io"
	"math/rand"

	"github.com/pkg/errors"
)

// Factory generates Items
type Factory struct {
	recipeList         []recipe
	availableMaterials map[string]category
	availableContents  map[string]category
	availableClasses   map[string]class
}

// Item represents a mundane RPG item complete with description and suggested
// gold value.
type Item struct {
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}

// New returns a new *Factory initialized with the given recipes, material
// categories, content categories, and classes.
func New(r, m, c, cl io.Reader) (*Factory, error) {
	var err error
	f := &Factory{}
	f.recipeList, err = readRecipes(r)
	if err != nil {
		return nil, err
	}

	f.availableMaterials, err = readCategories(m)
	if err != nil {
		return nil, err
	}

	f.availableContents, err = readCategories(c)
	if err != nil {
		return nil, err
	}

	f.availableClasses, err = readClasses(cl)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// FromFiles returns a new Factory with the given filenames loaded into their
// appropriate fields.
func FromFiles(recipesFile, materialsFile, contentsFile, classesFile string) (*Factory, error) {
	r, err := loadRecipes(recipesFile)
	if err != nil {
		return nil, err
	}

	m, err := loadCategories(materialsFile)
	if err != nil {
		return nil, err
	}

	c, err := loadCategories(contentsFile)
	if err != nil {
		return nil, err
	}

	cl, err := loadClasses(classesFile)
	if err != nil {
		return nil, err
	}

	return &Factory{
		recipeList:         r,
		availableMaterials: m,
		availableContents:  c,
		availableClasses:   cl,
	}, nil
}

// Item generates a new item
func (f *Factory) Item() (*Item, error) {
	r, err := f.randomRecipe()
	if err != nil {
		return nil, err
	}

	s, err := f.prepare(r)
	if err != nil {
		return nil, err
	}

	desc, err := compose(s)
	if err != nil {
		return nil, err
	}

	val, err := appraise(s)
	if err != nil {
		return nil, err
	}

	return &Item{Description: desc, Value: val}, nil
}

func (f *Factory) prepare(r recipe) (*stage, error) {
	s := stage{}
	s.base = r.Base
	s.value = r.Value

	m, err := r.material()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot retrieve material from recipe '%v'", r.Base)
	}
	s.material, err = f.getMaterial(m)
	if err != nil {
		return nil, err
	}

	c, err := r.content()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot retrieve content from recipe '%v'", r.Base)
	}
	s.content, err = f.getContent(c)
	if err != nil {
		return nil, err
	}

	s.class, err = f.getClass(r.Class)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (f *Factory) randomRecipe() (recipe, error) {
	if len(f.recipeList) < 1 {
		return recipe{}, errors.New("recipe slice is empty")
	}

	i := rand.Intn(len(f.recipeList))
	return f.recipeList[i], nil
}

func (f *Factory) getMaterial(name string) (*category, error) {
	m, ok := f.availableMaterials[name]
	if !ok {
		return nil, errors.Errorf("cannot find material '%v' in available materials", name)
	}

	return &m, nil
}

func (f *Factory) getContent(name string) (*category, error) {
	c, ok := f.availableContents[name]
	if !ok {
		return nil, errors.Errorf("cannot find content '%v' in available contents", name)
	}

	return &c, nil
}

func (f *Factory) getClass(name string) (*class, error) {
	c, ok := f.availableClasses[name]
	if !ok {
		return nil, errors.Errorf("cannot find class '%v' in available classes", name)
	}

	return &c, nil
}
