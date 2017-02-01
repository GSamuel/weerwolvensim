package character

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/mark"
	"github.com/gsamuel/weerwolvensim/resident"
)

type Weerwolf struct {
	resident *resident.Resident
	target   *resident.Resident
}

func (this *Weerwolf) Assigned() bool {
	return this.resident != nil
}

func (this *Weerwolf) Assign(resident *resident.Resident) error {
	if this.resident != nil {
		return fmt.Errorf("can't assign. This weerwolf is already assigned")
	}
	if resident.Assigned() {
		return fmt.Errorf("can't assign. target player(%d) is already assigned a character(%s)", resident.Id(), resident.Character())
	}

	resident.SetCharacter(mark.Weerwolf)
	this.resident = resident
	return nil
}

func (this *Weerwolf) Execute(targets ...*resident.Resident) error {
	if len(targets) != 1 {
		return fmt.Errorf("can't execute. weerwolf can only have 1 target. found %d", len(targets))
	}

	target := targets[0]

	if target.Character() == mark.Weerwolf {
		return fmt.Errorf("werewolves are not allowed to attack other werewolves")
	}

	if target.Has(mark.Attacked) {
		//potential bug because not monitoring target
		return nil
	}

	target.Add(mark.Attacked)
	this.target = target
	return nil
}

func (this *Weerwolf) Validate() error {
	if this.target == nil {
		return fmt.Errorf("can't validate. werewolf has no target selected")
	}

	if this.target.Has(mark.Attacked) {
		this.target.SetAlive(false)
	}

	return nil
}

func (this *Weerwolf) Type() Type {
	return WEERWOLF
}

func NewWeerwolf() Character {
	return &Weerwolf{}
}
