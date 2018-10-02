package iltb

import (
	"math/rand"
	"os"
	"testing"
)

const testMaterialFile = "testdata/category_material_test.json"
const testContentFile = "testdata/category_content_test.json"

func TestCategoryRandom(t *testing.T) {
	rand.Seed(1)

	c, err := loadCategories(testMaterialFile)
	if err != nil {
		t.Fatal(err)
	}

	wood, err := c["wood"].random()
	if err != nil {
		t.Fatal(err)
	}

	if wood != "corkwood" {
		t.Errorf("got: <%v>, want: <%v>", wood, "corkwood")
	}

	rand.Seed(2)

	metal, err := c["precious metal"].random()
	if err != nil {
		t.Fatal(err)
	}

	if metal != "rose gold" {
		t.Errorf("got: <%v>, want: <%v>", metal, "rose gold")
	}
}

func TestReadCategories(t *testing.T) {
	f, err := os.Open(testMaterialFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	c, err := readCategories(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(c) != 4 {
		t.Errorf("got: <%v>, want: <%v>", len(c), 4)
	}

	if c["precious metal"].Name != "precious metal" {
		t.Errorf("got: <%v>, want: <%v>", c["precious metal"].Name, "precious metal")
	}
}

func TestLoadCategories(t *testing.T) {
	c, err := loadCategories(testMaterialFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(c) != 4 {
		t.Errorf("got: <%v>, want: <%v>", len(c), 4)
	}

	if c["wood"].Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", c["wood"].Name, "wood")
	}
}
