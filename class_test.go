package main

import (
	"os"
	"testing"
)

const testClassFile = "class_test.json"

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

	if len(c) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(c), 3)
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

	if len(c) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(c), 3)
	}

	if c["clothing"].Format != "<article> <material> <base> <verb> <content>" {
		t.Errorf("got: <%v>, want: <%v>", c["clothing"].Format, "<article> <material> <base> <verb> <content>")
	}
}
