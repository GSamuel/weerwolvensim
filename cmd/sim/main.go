package main

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/model/character"
	"github.com/gsamuel/weerwolvensim/model/game"
)

const (
	CHAR = "Character"
)

func main() {
	names := []string{"Gideon", "Rosanne", "Naomi", "Sundri", "Sjoerd"}
	characters := []int{character.ZIENER, character.WEERWOLF, character.BURGER, character.CUPIDO, character.HEKS}

	g, err := game.New(names, characters)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	g.Update()

	/*
		//try assigning random characters to players
		as.Assign(character.ZIENER, 0)
		as.Assign(character.ZIENER, 0)
		as.Assign(character.BURGER, 0)

		as.Assign(character.ZIENER, 1)
		as.Assign(character.BURGER, 2)
		as.Assign(character.HEKS, 3)
		as.Assign(character.WEERWOLF, 4)

		//print all players and print if they are assigned a character
		for i, _ := range spelers {
			fmt.Println(spelers[i].Has(CHAR))
			if spelers[i].Has(CHAR) {
				value, _ := spelers[i].Value(CHAR)
				fmt.Println(value)
			}
		}*/
}
