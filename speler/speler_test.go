package speler

import (
	"testing"
)

func TestNew(t *testing.T) {
	speler := New()
	if speler == nil {
		t.Errorf("Factory returning nil instead of speler")
	}

}
