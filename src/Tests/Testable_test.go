package mymath

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(2, 3) == 5 &&
	Add(2, -3) == -1 &&
	Add(-3, 2) == -1 &&
	Add(-2, -3) == -5 &&
	Add(2, 0) == 2 {
	} else {
		t.Fail()
	}
}

