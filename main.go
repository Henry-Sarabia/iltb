package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	err := testLoadFactory()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func testLoadFactory() error {
	f, err := Load("recipes.json", "materials.json", "contents.json", "classes.json")
	if err != nil {
		return err
	}

	for j := 0; j < 10; j++ {
		i, err := f.Item()
		if err != nil {
			return err
		}
		fmt.Println(i)
	}

	return nil

}

func testSimpleFactory() error {
	i, err := os.Open("recipes.json")
	if err != nil {
		return err
	}
	defer i.Close()

	m, err := os.Open("materials.json")
	if err != nil {
		return err
	}
	defer m.Close()

	c, err := os.Open("contents.json")
	if err != nil {
		return err
	}
	defer c.Close()

	cl, err := os.Open("classes.json")
	if err != nil {
		return err
	}
	defer cl.Close()

	f, err := New(i, m, c, cl)
	if err != nil {
		return err
	}

	for j := 0; j < 10; j++ {
		fmt.Println(f.Item())
	}

	return nil
}
