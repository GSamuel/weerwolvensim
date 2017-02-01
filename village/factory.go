package village

import (
	"github.com/gsamuel/weerwolvensim/resident"
)

func New(residentNames []string) *Village {
	residents := []*resident.Resident{}

	for i, name := range residentNames {
		residents = append(residents, resident.New(i, name))
	}

	return &Village{residents: residents}
}
