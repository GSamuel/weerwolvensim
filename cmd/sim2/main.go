package main

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/character"
	"github.com/gsamuel/weerwolvensim/game"
)

func main() {

	names := []string{"Gideon", "Rosanne", "Naomi", "Sundri", "Sjoerd", "Esmee", "Anja"}
	c := game.NewConfig().SetBurgers(2).SetWeerwolven(2).SetHeks(true).SetCupido(true).SetGenezer(true)

	g, err := game.New(names, c)
	panicerr(err)

	err = g.Assign(character.CUPIDO, 0)
	panicerr(err)
	err = g.Assign(character.GENEZER, 1)
	panicerr(err)
	err = g.Assign(character.WEERWOLF, 2)
	panicerr(err)
	err = g.Assign(character.HEKS, 3)
	panicerr(err)
	//
	//
	//
	printstep(g.CurrentStep())
	g.Print()

	err = g.Execute(character.CUPIDO, 1, 2)
	printerr(err)

	err = g.Execute(character.GENEZER, 6)
	printerr(err)

	err = g.Execute(character.WEERWOLF, 0)
	printerr(err)

	err = g.Execute(character.HEKS)
	printerr(err)

	printstep(g.CurrentStep())
	g.Print()
}

func printstep(step game.Step) {
	fmt.Println("current step: ", step.Type(), step.CharType())
}

func printerr(err error) {
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}

func panicerr(err error) {
	if err != nil {
		panic(err)
	}
}
