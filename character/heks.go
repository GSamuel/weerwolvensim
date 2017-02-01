package character

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/mark"
	"github.com/gsamuel/weerwolvensim/resident"
)

type Heks struct {
	resident   *resident.Resident
	target1    *resident.Resident
	target2    *resident.Resident
	lifeElixer bool
	poison     bool
}

func (this *Heks) Assigned() bool {
	return this.resident != nil
}

func (this *Heks) Assign(resident *resident.Resident) error {
	if this.resident != nil {
		return fmt.Errorf("can't assign. This heks is already assigned")
	}
	if resident.Assigned() {
		return fmt.Errorf("can't assign. target player(%d) is already assigned a character(%s)", resident.Id(), resident.Character())
	}

	resident.SetCharacter(mark.Heks)
	this.resident = resident
	return nil
}

func (this *Heks) Execute(targets ...*resident.Resident) error {
	if len(targets) > 2 {
		return fmt.Errorf("heks can't have more then 2 targets selected. found: %d", len(targets))
	}

	var saveTarget *resident.Resident
	var killTarget *resident.Resident

	for i, _ := range targets {
		if targets[i].Has(mark.Attacked) {
			if saveTarget != nil {
				return fmt.Errorf("can't execute. heks can't use life elixer on two players")
			}
			saveTarget = targets[i]
		} else {
			if killTarget != nil {
				return fmt.Errorf("can't execute. heks can't poison a target marked as attacked")
			}
			killTarget = targets[i]
		}
	}

	if saveTarget != nil && !this.lifeElixer || killTarget != nil && !this.poison {
		return fmt.Errorf("heks already used one of the elixers she wants to use.")
	}

	if saveTarget != nil {
		saveTarget.Add(mark.Revived)
		this.target1 = saveTarget
		this.lifeElixer = false
	}

	if killTarget != nil {
		killTarget.Add(mark.Poisoned)
		this.target2 = killTarget
		this.poison = false
	}
	return nil
}

func (this *Heks) Validate() error {
	if this.target1 == nil && this.target2 == nil {
		return nil //Have to look at this later.
	}

	if this.target1 != nil {
		if this.target1.Has(mark.Revived) {
			this.target1.Remove(mark.Attacked)
		}
	}

	if this.target2 != nil {
		if this.target2.Has(mark.Poisoned) {
			this.target2.SetAlive(false)
		}
	}

	return nil
}

func (this *Heks) Type() Type {
	return HEKS
}

func NewHeks() Character {
	return &Heks{lifeElixer: true, poison: true}
}
