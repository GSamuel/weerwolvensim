package game

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/character"
	"github.com/gsamuel/weerwolvensim/village"
)

func New(residentNames []string, config Config) (*Game, error) {

	if len(residentNames) != config.Count() {
		return nil, fmt.Errorf("amount of player names (%d) and amount of characters (%d) should be the same", len(residentNames), config.Count())
	}

	vil := village.New(residentNames)
	pool := character.NewPool(config.CharacterTypes())

	return &Game{village: vil, pool: pool, config: config, steps: CreateSteps()}, nil
}

/*
genezer.Execute(r1)

	wolf.Assign(r3)
	wolf.Execute(r5)

	heks.Assign(r5)
	heks.Execute(r5, r0)

	heks.Validate()
	genezer.Validate()
	wolf.Validate()
	cupid.Validate()*/

func CreateSteps() []Step {
	steps := []Step{
		NewAssignment(character.CUPIDO),
		NewExecution(character.CUPIDO),

		NewAssignment(character.GENEZER),
		NewExecution(character.GENEZER),

		NewAssignment(character.WEERWOLF),
		NewExecution(character.WEERWOLF),

		NewAssignment(character.HEKS),
		NewExecution(character.HEKS),

		NewValidation(character.HEKS),
		NewValidation(character.GENEZER),
		NewValidation(character.WEERWOLF),
		NewValidation(character.CUPIDO),
	}
	return steps
}

func NewConfig() Config {
	return Config{}
}
