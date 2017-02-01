package game

import (
	"github.com/gsamuel/weerwolvensim/character"
	"github.com/gsamuel/weerwolvensim/village"
)

type StepType int

const (
	ASSIGNMENT StepType = iota
	EXECUTION
	VALIDATION
)

func (s StepType) String() string {
	if s == ASSIGNMENT {
		return "assignment"
	}

	if s == EXECUTION {
		return "execution"
	}

	if s == VALIDATION {
		return "validation"
	}

	return "unknown"
}

type Step interface {
	Check(*village.Village, *character.Pool) bool
	CharType() character.Type
	Type() StepType
}

type Assignment struct {
	t character.Type
}

func (this *Assignment) Check(v *village.Village, pool *character.Pool) bool {
	if !pool.Has(this.t) {
		return false
	}

	c, err := pool.Find(this.t)

	if err != nil {
		panic(err)
	}

	if !c.Assigned() {
		return true
	}

	return false
}

func (this *Assignment) CharType() character.Type {
	return this.t
}

func (this *Assignment) Type() StepType {
	return ASSIGNMENT
}

func NewAssignment(t character.Type) Step {
	return &Assignment{t: t}
}

type Execution struct {
	t character.Type
}

func (this *Execution) Check(v *village.Village, pool *character.Pool) bool {
	return true
}

func (this *Execution) CharType() character.Type {
	return this.t
}

func (this *Execution) Type() StepType {
	return EXECUTION
}

func NewExecution(t character.Type) Step {
	return &Execution{t: t}
}

type Validation struct {
	t character.Type
}

func (this *Validation) Check(v *village.Village, pool *character.Pool) bool {
	if !pool.Has(this.t) {
		return false
	}

	c, err := pool.Find(this.t)

	if err != nil {
		panic(err)
	}

	err = c.Validate()

	if err != nil {
		panic(err)
	}

	return false
}

func (this *Validation) CharType() character.Type {
	return this.t
}

func (this *Validation) Type() StepType {
	return VALIDATION
}

func NewValidation(t character.Type) Step {
	return &Validation{t: t}
}
