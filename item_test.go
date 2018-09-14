package main

import (
	"os"
	"testing"
)

const testItemFile = "items_test.json"

func TestReadItems(t *testing.T) {
	f, err := os.Open(testItemFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	i, err := readItems(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(i) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(i), 3)
	}

	if i[0].Base != "statue" {
		t.Errorf("got: <%v>, want: <%v>", i[0].Base, "statue")
	}
}

func TestLoadItems(t *testing.T) {
	i, err := loadItems(testItemFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(i) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(i), 3)
	}

	if i[0].Base != "statue" {
		t.Errorf("got: <%v>, want: <%v>", i[0].Base, "statue")
	}
}
