package character

import (
	"github.com/gsamuel/weerwolvensim/resident"
)

type Character interface {
	Assigned() bool
	Assign(*resident.Resident) error
	Execute(...*resident.Resident) error
	Validate() error
	Type() Type
}
