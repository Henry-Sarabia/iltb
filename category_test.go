package main

import (
	"os"
	"testing"
)

const testCategoryFile = "category_test.json"

func TestReadCategories(t *testing.T) {
	f, err := os.Open(testCategoryFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	c, err := readCategories(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(c) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(c), 3)
	}

	if c["precious metal"].Name != "precious metal" {
		t.Errorf("got: <%v>, want: <%v>", c["precious metal"].Name, "precious metal")
	}
}

func TestLoadCategories(t *testing.T) {
	c, err := loadCategories(testCategoryFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(c) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(c), 3)
	}

	if c["wood"].Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", c["wood"].Name, "wood")
	}
}
