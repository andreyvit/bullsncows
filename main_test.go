package main

import (
	"testing"
)

func TestCompute1(t *testing.T) {
	tests := []struct {
		move string
		orig string
		e1   string
		e2   string
		e3   string
	}{
		{"122456", "100000", "1Б", "1Б", "1Б"},
		{"122456", "000001", "1К", "1К", "1К"},
		{"122456", "200000", "2К", "2К", "1К"},
		{"122456", "011000", "2К", "1К", "2К"},
		{"122456", "110000", "1Б 1К", "1Б", "1Б 1К"},
	}
	for _, tt := range tests {
		move := Parse(tt.move)
		orig := Parse(tt.orig)

		a1 := Compute1(&move, &orig).String()
		if a1 != tt.e1 {
			t.Errorf("** Compute1(%q, %q) == %q, wanted %q", tt.move, tt.orig, a1, tt.e1)
		} else {
			t.Logf("✓ Compute1(%q, %q) == %q", tt.move, tt.orig, a1)
		}

		a2 := Compute2(&move, &orig).String()
		if a2 != tt.e2 {
			t.Errorf("** Compute2(%q, %q) == %q, wanted %q", tt.move, tt.orig, a2, tt.e2)
		} else {
			t.Logf("✓ Compute2(%q, %q) == %q", tt.move, tt.orig, a2)
		}

		a3 := Compute3(&move, &orig).String()
		if a3 != tt.e3 {
			t.Errorf("** Compute3(%q, %q) == %q, wanted %q", tt.move, tt.orig, a3, tt.e3)
		} else {
			t.Logf("✓ Compute3(%q, %q) == %q", tt.move, tt.orig, a3)
		}
	}
}
