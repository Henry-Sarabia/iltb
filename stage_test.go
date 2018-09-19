package iltb

import (
	"math/rand"
	"testing"
)

func TestCompose(t *testing.T) {
	rand.Seed(1)

	f, err := FromFiles(testRecipeFile, testMaterialFile, testContentFile, testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	r := f.recipeList[0]

	s, err := f.prepare(r)
	if err != nil {
		t.Fatal(err)
	}

	item, err := compose(s)
	if err != nil {
		t.Fatal(err)
	}

	if item != "a pine wood statue designed like a gynosphinx" {
		t.Errorf("got: <%v>, want: <%v>", item, "a pine wood statue designed like a gynosphinx")
	}
}

func TestTokenize(t *testing.T) {
	cl, err := loadClasses(testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	tok, err := tokenize(cl["art"].Format)
	if err != nil {
		t.Fatal(err)
	}

	if len(tok) != 6 {
		t.Errorf("got: <%v>, want: <%v>", len(tok), 6)
	}

	if tok[0] != "<article>" {
		t.Errorf("got: <%v>, want: <%v>", tok[0], "<article>")
	}

	if tok[5] != "<content>" {
		t.Errorf("got: <%v>, want: <%v>", tok[5], "<content>")
	}
}

func TestParse(t *testing.T) {
	rand.Seed(1)

	f, err := FromFiles(testRecipeFile, testMaterialFile, testContentFile, testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	r := f.recipeList[0]

	s, err := f.prepare(r)
	if err != nil {
		t.Fatal(err)
	}

	tok, err := tokenize(s.class.Format)
	if err != nil {
		t.Fatal(err)
	}

	item, err := parse(s, tok)
	if err != nil {
		t.Fatal(err)
	}

	if item != "a pine wood statue designed like a gynosphinx" {
		t.Errorf("got: <%v>, want: <%v>", item, "a pine wood statue designed like a gynosphinx")
	}
}
