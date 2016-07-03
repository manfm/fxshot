package numbers

import (
	"testing"
)

func TestRound(t *testing.T) {
	if Round(1.12345678) != 1 {
		t.Error("Failed round down")
	}

	if Round(1.56789) != 2 {
		t.Error("Failed round up")
	}
}

func TestRoundPlus(t *testing.T) {
	if RoundPlus(1.12345678, 2) != 1.12 {
		t.Error("Failed round down")
	}

	if RoundPlus(1.56789, 2) != 1.57 {
		t.Error("Failed round up")
	}
}
