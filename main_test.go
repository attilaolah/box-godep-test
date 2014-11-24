package potato_test

import (
	"testing"

	"github.com/attilaolah/box-godep-test"
)

func TestPerm(t *testing.T) {
	m := potato.Perm()

	if x := m.MapFrom(m.MapTo(37)); x != 37 {
		t.Errorf("expected 37, got %d", x)
	}
}
