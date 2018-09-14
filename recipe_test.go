package main

import (
	"os"
	"testing"
)

const testRecipeFile = "recipes_test.json"

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
