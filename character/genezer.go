package character

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/mark"
	"github.com/gsamuel/weerwolvensim/resident"
)

type Genezer struct {
	resident *resident.Resident
	target   *resident.Resident
}

func (this *Genezer) Assigned() bool {
	return this.resident != nil
}

func (this *Genezer) Assign(resident *resident.Resident) error {
	if this.resident != nil {
		return fmt.Errorf("can't assign. This genezer is already assigned")
	}
	if resident.Assigned() {
		return fmt.Errorf("can't assign. target player(%d) is already assigned a character(%s)", resident.Id(), resident.Character())
	}

	resident.SetCharacter(mark.Genezer)
	this.resident = resident
	return nil
}

func (this *Genezer) Execute(targets ...*resident.Resident) error {
	if len(targets) != 1 {
		return fmt.Errorf("can't execute genezer. number of targets %d instead of 1", len(targets))
	}

	target := targets[0]

	if this.target != nil {
		if this.target.Equal(target) {
			return fmt.Errorf("can't execute. genezer can't select the same target twice in a row")
		}
	}

	if target.Has(mark.Protected) {
		//genezer doesn't monitor protected target now
		return nil
	}

	target.Add(mark.Protected)
	this.target = target
	return nil
}

func (this *Genezer) Validate() error {
	if this.target == nil {
		return fmt.Errorf("can't validate. no target selected for genezer")
	}

	if this.target.Has(mark.Protected) {
		this.target.Remove(mark.Attacked)
	}
	return nil
}

func (this *Genezer) Type() Type {
	return GENEZER
}

func NewGenezer() Character {
	return &Genezer{}
}
