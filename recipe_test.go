package main

import (
	"math/rand"
	"os"
	"testing"
)

const testRecipeFile = "testdata/recipe_test.json"

func TestRecipeMaterial(t *testing.T) {
	rand.Seed(1)

	r, err := loadRecipes(testRecipeFile)
	if err != nil {
		t.Fatal(err)
	}

	wood, err := r[0].material()
	if err != nil {
		t.Fatal(err)
	}

	if wood != "wood" {
		t.Errorf("got: <%v>, want: <%v>", wood, "wood")
	}

	rand.Seed(2)

	metal, err := r[0].material()
	if err != nil {
		t.Fatal(err)
	}

	if metal != "precious metal" {
		t.Errorf("got: <%v>, want: <%v>", metal, "precious metal")
	}

}

func TestRecipeContent(t *testing.T) {
	rand.Seed(1)

	r, err := loadRecipes(testRecipeFile)
	if err != nil {
		t.Fatal(err)
	}

	monstrosity, err := r[0].content()
	if err != nil {
		t.Fatal(err)
	}

	if monstrosity != "monstrosity" {
		t.Errorf("got: <%v>, want: <%v>", monstrosity, "monstrosity")
	}

	rand.Seed(2)

	beast, err := r[0].content()
	if err != nil {
		t.Fatal(err)
	}

	if beast != "beast" {
		t.Errorf("got: <%v>, want: <%v>", beast, "beast")
	}
}

func TestReadRecipes(t *testing.T) {
	f, err := os.Open(testRecipeFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r, err := readRecipes(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(r) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(r), 3)
	}

	if r[0].Base != "statue" {
		t.Errorf("got: <%v>, want: <%v>", r[0].Base, "statue")
	}
}

func TestLoadRecipes(t *testing.T) {
	r, err := loadRecipes(testRecipeFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(r) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(r), 3)
	}

	if r[0].Base != "statue" {
		t.Errorf("got: <%v>, want: <%v>", r[0].Base, "statue")
	}
}
