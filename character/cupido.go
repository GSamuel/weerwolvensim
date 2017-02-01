package character

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/mark"
	"github.com/gsamuel/weerwolvensim/resident"
)

type Cupido struct {
	resident *resident.Resident
	lover1   *resident.Resident
	lover2   *resident.Resident
}

func (this *Cupido) Assigned() bool {
	return this.resident != nil
}

func (this *Cupido) Assign(resident *resident.Resident) error {
	if this.resident != nil {
		return fmt.Errorf("can't assign. This cupido is already assigned")
	}
	if resident.Assigned() {
		return fmt.Errorf("can't assign. target player(%d) is already assigned a character(%s)", resident.Id(), resident.Character())
	}

	resident.SetCharacter(mark.Cupido)
	this.resident = resident
	return nil
}

func (this *Cupido) Execute(targets ...*resident.Resident) error {
	if this.lover1 != nil || this.lover2 != nil || len(targets) != 2 {
		return fmt.Errorf("cupido: lovers already selected or amount of targets not equal to 2")
	}

	if targets[0].Equal(targets[1]) {
		return fmt.Errorf("can't execute. cupido should select two different targets")
	}

	targets[0].Add(mark.Lover)
	targets[1].Add(mark.Lover)

	this.lover1 = targets[0]
	this.lover2 = targets[1]
	return nil
}

func (this *Cupido) Validate() error {
	if this.lover1 == nil || this.lover2 == nil {
		return fmt.Errorf("can't validate. no lovers selected")
	}

	if this.lover1.Alive() && this.lover2.Dead() {
		this.lover1.SetAlive(false)
	} else if this.lover1.Dead() && this.lover2.Alive() {
		this.lover2.SetAlive(false)
	}
	return nil
}

func (this *Cupido) Type() Type {
	return CUPIDO
}

func NewCupido() Character {
	return &Cupido{}
}
