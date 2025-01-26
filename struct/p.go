package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	name   string
	age    int
	skills []string
}

type Animal struct {
	dog  Dog
	kind string
}

type Skills []string

func (s Skills) findSkill(skill string) bool {
	for _, _skill := range s {
		if _skill == skill {
			return true
		}
	}
	return false
}

func newAnimal(dog Dog, kind string) *Animal {
	return &Animal{
		dog:  dog,
		kind: kind,
	}
}

func main() {

	s := Skills{"run", "jump"}
	b := s.findSkill("run2")

	dog := Dog{
		name:   "Max",
		age:    2,
		skills: []string{"run", "jump"},
	}

	a := newAnimal(dog, "狗狗")

	d, err := json.Marshal(a)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Hello, World!", b, a, d)
}
