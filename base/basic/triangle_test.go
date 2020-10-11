package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{3000, 4000, 5000},
	}

	for _, test := range tests {
		if actual := calTriangle(test.a, test.b); actual != test.c {
			t.Errorf("callTriangle(%d,%d); got %d; expected %d", test.a, test.b, actual, test.c)
		}
	}
}