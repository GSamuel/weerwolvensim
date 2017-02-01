package game

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/character"
	"github.com/gsamuel/weerwolvensim/resident"
	"github.com/gsamuel/weerwolvensim/village"
)

type Game struct {
	phase   Phase
	steps   []Step
	step    int
	village *village.Village
	pool    *character.Pool
	config  Config
}

func (this *Game) CurrentStep() Step {
	this.Update()
	return this.steps[this.step]
}

func (this *Game) Update() {
	current := this.steps[this.step].Check(this.village, this.pool)

	for !current {
		this.nextStep()
		current = this.steps[this.step].Check(this.village, this.pool)
	}
}

func (this *Game) nextStep() {
	this.step = this.step + 1
	if this.step >= len(this.steps) {
		this.step = 0
		this.phase = this.phase.Next()
	}
}

func (this *Game) Execute(char character.Type, targets ...int) error {

	step := this.CurrentStep()

	if step.Type() != EXECUTION {
		return fmt.Errorf("cannot execute action when current step is %s", step.Type())
	}

	if step.CharType() != char {
		return fmt.Errorf("cannot execute action for %s when current execution is for %s", char, step.CharType())
	}

	c, err := this.pool.Find(char)

	if err != nil {
		panic(err)
	}

	ress := []*resident.Resident{}

	for _, t := range targets {
		r, err := this.village.Get(t)
		if err != nil {
			return err
		}
		ress = append(ress, r)
	}

	err = c.Execute(ress...)

	if err != nil {
		return err
	}

	this.nextStep()

	return nil
}

func (this *Game) Assign(char character.Type, id int) error {
	r, err := this.village.Get(id)
	if err != nil {
		return fmt.Errorf("player with id %d does not exist", r.Id())
	}

	if r.Assigned() {
		return fmt.Errorf("can't assign %s to player with id %d. already has character %s assigned.", char, r.Id(), r.Character())
	}

	chars := this.pool.FindWith(func(c character.Character) bool { return !c.Assigned() && c.Type() == char })
	if len(chars) == 0 {
		return fmt.Errorf("no characters of type %s left to assign", char)
	}

	return chars[0].Assign(r)
}

func (this *Game) Print() {
	for _, r := range this.village.Residents() {
		fmt.Println(r)
	}
	fmt.Println("--------------------------------")
}

type Phase int

func (p Phase) Round() int {
	return int(p) / 2
}

func (p Phase) Next() Phase {
	return Phase(int(p) + 1)
}

func (p Phase) Night() bool {
	return int(p)%2 == 0
}

func (p Phase) Day() bool {
	return int(p)%2 == 1
}
