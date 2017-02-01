package village

import (
	"fmt"
	"github.com/gsamuel/weerwolvensim/resident"
)

type Village struct {
	residents []*resident.Resident
}

func (this *Village) Get(id int) (*resident.Resident, error) {
	for i, _ := range this.residents {
		if this.residents[i].Id() == id {
			return this.residents[i], nil
		}
	}
	return nil, fmt.Errorf("Resident with id: %d Not Found", id)
}

func (this *Village) Residents() []*resident.Resident {
	return this.residents
}

func (this *Village) Count() int {
	return len(this.residents)
}
