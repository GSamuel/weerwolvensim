package game

import (
	"github.com/gsamuel/weerwolvensim/character"
)

type Config struct {
	weerwolven int
	burgers    int
	cupido     bool
	genezer    bool
	heks       bool
}

func (this Config) Weerwolven() int {
	return this.weerwolven
}

func (this Config) Burgers() int {
	return this.burgers
}

func (this Config) Cupido() bool {
	return this.cupido
}

func (this Config) Genezer() bool {
	return this.genezer
}

func (this Config) Heks() bool {
	return this.heks
}

func (this Config) SetWeerwolven(value int) Config {
	if value >= 0 {
		this.weerwolven = value
	}
	return this

}

func (this Config) SetBurgers(value int) Config {
	if value >= 0 {
		this.burgers = value
	}

	return this
}

func (this Config) SetCupido(value bool) Config {
	this.cupido = value
	return this
}

func (this Config) SetGenezer(value bool) Config {
	this.genezer = value
	return this
}

func (this Config) SetHeks(value bool) Config {
	this.heks = value
	return this
}

func (this Config) Count() int {
	count := 0

	count = count + this.weerwolven
	count = count + this.burgers

	if this.cupido {
		count = count + 1
	}
	if this.genezer {
		count = count + 1
	}
	if this.heks {
		count = count + 1
	}

	return count

}

func (this Config) CharacterTypes() []character.Type {

	types := []character.Type{}

	for i := 0; i < this.weerwolven; i++ {
		types = append(types, character.WEERWOLF)
	}

	for i := 0; i < this.burgers; i++ {
		types = append(types, character.BURGER)
	}

	if this.cupido {
		types = append(types, character.CUPIDO)
	}

	if this.genezer {
		types = append(types, character.GENEZER)
	}

	if this.heks {
		types = append(types, character.HEKS)
	}

	return types
}
