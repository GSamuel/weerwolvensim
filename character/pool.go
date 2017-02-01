package character

import (
	"fmt"
)

type Pool struct {
	characters []Character
}

func (this *Pool) FindAll(t Type) []Character {

	chars := []Character{}

	for i, _ := range this.characters {
		if this.characters[i].Type() == t {
			chars = append(chars, this.characters[i])
		}
	}

	return chars
}

func (this *Pool) Find(t Type) (Character, error) {
	for i, _ := range this.characters {
		if this.characters[i].Type() == t {
			return this.characters[i], nil
		}
	}
	return nil, fmt.Errorf("could not find character %s.", t)
}

func (this *Pool) Has(t Type) bool {
	for i, _ := range this.characters {
		if this.characters[i].Type() == t {
			return true
		}
	}
	return false
}

func (this *Pool) Count(t Type) int {
	count := 0

	for i, _ := range this.characters {
		if this.characters[i].Type() == t {
			count = count + 1
		}
	}

	return count
}

func (this *Pool) FindWith(f func(Character) bool) []Character {
	chars := []Character{}

	for i, _ := range this.characters {
		if f(this.characters[i]) {
			chars = append(chars, this.characters[i])
		}
	}

	return chars
}
