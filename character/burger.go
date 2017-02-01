package character

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/mark"
	"github.com/gsamuel/weerwolvensim/resident"
)

type Burger struct {
	resident *resident.Resident
	target   *resident.Resident
}

func (this *Burger) Assigned() bool {
	return this.resident != nil
}

func (this *Burger) Assign(resident *resident.Resident) error {
	if this.resident != nil {
		return fmt.Errorf("can't assign. This burger is already assigned")
	}

	if resident.Assigned() {
		return fmt.Errorf("can't assign. target player(%d) is already assigned a character(%s)", resident.Id(), resident.Character())
	}

	resident.SetCharacter(mark.Burger)
	this.resident = resident
	return nil
}

func (this *Burger) Execute(targets ...*resident.Resident) error {
	l := len(targets)
	if l != 1 {
		return fmt.Errorf("can't execute burger. needs 1 target but found: %d", l)
	}

	target := targets[0]

	target.Add(mark.Lynched)
	this.target = target
	return nil
}

func (this *Burger) Validate() error {
	if this.target == nil {
		return fmt.Errorf("can't validate burger. No target selected")
	}

	if this.target.Has(mark.Lynched) {
		this.target.SetAlive(false)
	}

	return nil
}

func (this *Burger) Type() Type {
	return BURGER
}

func NewBurger() Character {
	return &Burger{}
}
