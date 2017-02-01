package resident

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/mark"
)

type Resident struct {
	id        int
	name      string
	alive     bool
	character mark.Character
	marks     []mark.Mark
}

func (this *Resident) Id() int {
	return this.id
}

func (this *Resident) Name() string {
	return this.name
}

func (this *Resident) Alive() bool {
	return this.alive
}

func (this *Resident) Dead() bool {
	return !this.Alive()
}

func (this *Resident) SetAlive(alive bool) {
	this.alive = alive
}

func (this *Resident) Character() mark.Character {
	return this.character
}

func (this *Resident) SetCharacter(character mark.Character) {
	this.character = character
}

func (this *Resident) Assigned() bool {
	return this.character != mark.Unassigned
}

func (this *Resident) Add(mark mark.Mark) {
	this.marks = append(this.marks, mark)
}

func (this *Resident) Has(mark mark.Mark) bool {
	for _, m := range this.marks {
		if m == mark {
			return true
		}
	}
	return false
}

func (this *Resident) Remove(mark mark.Mark) {
	for i, m := range this.marks {
		if m == mark {
			j := len(this.marks) - 1
			this.marks[i] = this.marks[j]
			this.marks = this.marks[:j]
		}

	}
}

func (this *Resident) Equal(resident *Resident) bool {
	return this.Id() == resident.Id()
}

func (this *Resident) String() string {
	return fmt.Sprintf("{%d %s %v %s %v}", this.id, this.name, this.alive, this.character, this.marks)
}

func New(id int, name string) *Resident {
	return &Resident{id: id, name: name, alive: true, character: mark.Unassigned, marks: []mark.Mark{}}
}
