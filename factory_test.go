package iltb

import (
	"math/rand"
	"os"
	"testing"
)

func TestFactoryNew(t *testing.T) {
	r, err := os.Open(testRecipeFile)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	m, err := os.Open(testMaterialFile)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()

	c, err := os.Open(testContentFile)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	cl, err := os.Open(testClassFile)
	if err != nil {
		t.Fatal(err)
	}
	defer cl.Close()

	f, err := New(r, m, c, cl)
	if err != nil {
		t.Fatal(err)
	}

	if len(f.recipeList) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(f.recipeList), 3)
	}

	if f.recipeList[0].Base != "statue" {
		t.Errorf("got: <%v>, want <%v>", f.recipeList[0].Base, "statue")
	}

	if len(f.availableMaterials) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(f.availableMaterials), 3)
	}

	if f.availableMaterials["wood"].Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", f.availableMaterials["wood"].Name, "wood")
	}

	if len(f.availableContents) != 4 {
		t.Errorf("got: <%v>, want: <%v>", len(f.availableContents), 4)
	}

	if f.availableContents["beverage"].Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", f.availableContents["beverage"].Name, "beverage")
	}

	if len(f.availableClasses) != 4 {
		t.Errorf("got: <%v>, want: <%v>", len(f.availableClasses), 4)
	}

	if f.availableClasses["container"].Name != "container" {
		t.Errorf("got: <%v>, want: <%v>", f.availableClasses["container"].Name, "container")
	}
}

func TestFactoryFromFiles(t *testing.T) {
	f, err := FromFiles(testRecipeFile, testMaterialFile, testContentFile, testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(f.recipeList) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(f.recipeList), 3)
	}

	if f.recipeList[0].Base != "statue" {
		t.Errorf("got: <%v>, want <%v>", f.recipeList[0].Base, "statue")
	}

	if len(f.availableMaterials) != 3 {
		t.Errorf("got: <%v>, want: <%v>", len(f.availableMaterials), 3)
	}

	if f.availableMaterials["wood"].Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", f.availableMaterials["wood"].Name, "wood")
	}

	if len(f.availableContents) != 4 {
		t.Errorf("got: <%v>, want: <%v>", len(f.availableContents), 4)
	}

	if f.availableContents["beverage"].Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", f.availableContents["beverage"].Name, "beverage")
	}

	if len(f.availableClasses) != 4 {
		t.Errorf("got: <%v>, want: <%v>", len(f.availableClasses), 4)
	}

	if f.availableClasses["container"].Name != "container" {
		t.Errorf("got: <%v>, want: <%v>", f.availableClasses["container"].Name, "container")
	}
}

func TestFactoryRandomRecipe(t *testing.T) {
	rand.Seed(1)

	f, err := FromFiles(testRecipeFile, testMaterialFile, testContentFile, testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	ring, err := f.randomRecipe()
	if err != nil {
		t.Fatal(err)
	}

	if ring.Base != "ring" {
		t.Errorf("got: <%v>, want: <%v>", ring.Base, "ring")
	}

	if len(ring.MaterialChoices) != 2 {
		t.Errorf("got: <%v>, want: <%v>", len(ring.MaterialChoices), 2)
	}

	rand.Seed(2)

	wineskin, err := f.randomRecipe()
	if err != nil {
		t.Fatal(err)
	}

	if wineskin.Base != "wineskin" {
		t.Errorf("got: <%v>, want: <%v>", wineskin.Base, "wineskin")
	}

	if len(wineskin.ContentChoices) != 1 {
		t.Errorf("got: <%v>, want: <%v>", len(wineskin.ContentChoices), 1)
	}
}

func TestItem(t *testing.T) {
	rand.Seed(1)

	f, err := FromFiles(testRecipeFile, testMaterialFile, testContentFile, testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	it, err := f.Item()
	if err != nil {
		t.Fatal(err)
	}

	if it != "an elm wood ring crafted in the style of a cat" {
		t.Errorf("got: <%v>, want: <%v>", it, "an elm wood ring crafted in the style of a cat")
	}
}

func TestFactoryPrepare(t *testing.T) {
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

	if s.base != "statue" {
		t.Errorf("got: <%v>, want: <%v>", s.base, "statue")
	}

	if s.class.Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", s.class, "art")
	}
}

func TestFactoryGetMaterial(t *testing.T) {
	f, err := FromFiles(testRecipeFile, testMaterialFile, testContentFile, testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	wood, err := f.getMaterial("wood")
	if err != nil {
		t.Fatal(err)
	}

	if wood.Name != "wood" {
		t.Errorf("got: <%v>, want: <%v>", wood.Name, "wood")
	}

	if len(wood.Types) != 18 {
		t.Errorf("got: <%v>, want: <%v>", len(wood.Types), 18)
	}

	metal, err := f.getMaterial("precious metal")
	if err != nil {
		t.Fatal(err)
	}

	if metal.Name != "precious metal" {
		t.Errorf("got: <%v>, want: <%v>", metal.Name, "precious metal")
	}

	if len(metal.Types) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(metal.Types), 7)
	}
}

func TestFactoryGetContent(t *testing.T) {
	f, err := FromFiles(testRecipeFile, testMaterialFile, testContentFile, testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	bev, err := f.getContent("beverage")
	if err != nil {
		t.Fatal(err)
	}

	if bev.Name != "beverage" {
		t.Errorf("got: <%v>, want: <%v>", bev.Name, "beverage")
	}

	if len(bev.Types) != 14 {
		t.Errorf("got: <%v>, want: <%v>", len(bev.Types), 14)
	}

	hu, err := f.getContent("humanoid")
	if err != nil {
		t.Fatal(err)
	}

	if hu.Name != "humanoid" {
		t.Errorf("got: <%v>, want: <%v>", hu.Name, "humanoid")
	}

	if len(hu.Types) != 23 {
		t.Errorf("got: <%v>, want: <%v>", len(hu.Types), 23)
	}
}

func TestFactoryGetClass(t *testing.T) {
	f, err := FromFiles(testRecipeFile, testMaterialFile, testContentFile, testClassFile)
	if err != nil {
		t.Fatal(err)
	}

	art, err := f.getClass("art")
	if err != nil {
		t.Fatal(err)
	}

	if art.Name != "art" {
		t.Errorf("got: <%v>, want: <%v>", art.Name, "art")
	}

	if len(art.Verbs) != 8 {
		t.Errorf("got: <%v>, want: <%v>", len(art.Verbs), 8)
	}

	cloth, err := f.getClass("clothing")
	if err != nil {
		t.Fatal(err)
	}

	if cloth.Name != "clothing" {
		t.Errorf("got: <%v>, want: <%v>", cloth.Name, "clothing")
	}

	if len(cloth.Verbs) != 7 {
		t.Errorf("got: <%v>, want: <%v>", len(cloth.Verbs), 7)
	}
}
