package iltb

import (
	"math/rand"
	"os"
	"testing"
)

const testClassFile = "testdata/class_test.json"

func TestClassRandomVerb(t *testing.T) {
	rand.Seed(1)

	c, err := loadClasses(testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	clothing, err := c["clothing"].randomVerb()
	if err != nil {
		t.Fatal(err)
	}

	if clothing != "faintly colored" {
		t.Errorf("got: <%v>, want: <%v>", clothing, "faintly colored")
	}

	rand.Seed(2)

	art, err := c["art"].randomVerb()
	if err != nil {
		t.Fatal(err)
	}

	if art != "built to look like" {
		t.Errorf("got: <%v>, want: <%v>", art, "built to look like")
	}
}

func TestReadClasses(t *testing.T) {
	f, err := os.Open(testClassFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	c, err := readClasses(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(c) != 4 {
		t.Errorf("got: <%v>, want: <%v>", len(c), 4)
	}

	if c["clothing"].Format != "<article> <material> <base> <verb> <content>" {
		t.Errorf("got: <%v>, want: <%v>", c["clothing"].Format, "<article> <material> <base> <verb> <content>")
	}
}

func TestLoadClasses(t *testing.T) {
	c, err := loadClasses(testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(c) != 4 {
		t.Errorf("got: <%v>, want: <%v>", len(c), 4)
	}

	if c["clothing"].Format != "<article> <material> <base> <verb> <content>" {
		t.Errorf("got: <%v>, want: <%v>", c["clothing"].Format, "<article> <material> <base> <verb> <content>")
	}
}
