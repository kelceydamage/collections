package collections

import (
	"strconv"
	"testing"
)

var s = IntSlice{1, 2, 3, 4, 33, 2, -1, -8}

// Function interface tests
//-------------------------------------------------------------------------------------------------- <-100
func TestIndex(t *testing.T) {
	t.Log("Testing: Index func interface")
	n := Index(&s, 33)
	if n != 4 {
		t.Error("Index() not expected value of 4: " + strconv.Itoa(n))
	}
}

func TestIndexRight(t *testing.T) {
	t.Log("Testing: IndexRight func interface")
	n := IndexRight(&s, 33)
	if n != 4 {
		t.Error("IndexRight() not expected value of 4: " + strconv.Itoa(n))
	}
}

func TestLess(t *testing.T) {
	t.Log("Testing: Less func interface")
	n := Less(&s, 3, 4)
	if n != true {
		t.Error("Less() not expected value of 4: " + strconv.FormatBool(n))
	}
}

func TestSort(t *testing.T) {
	t.Log("Testing: Sort func interface")
	target := IntSlice{-8, -1, 1, 2, 2, 3, 4, 33}
	Sort(&s)
	for i := range s {
		if (*&s)[i] != (*&target)[i] {
			t.Error("Sort() order not as expected, ascending")
		}
	}
}

func TestReverse(t *testing.T) {
	t.Log("Testing: Reverse func interface")
	target := IntSlice{33, 4, 3, 2, 2, 1, -1, -8}
	Reverse(&s)
	for i := range s {
		if (*&s)[i] != (*&target)[i] {
			t.Error("Reverse() order not as expected, descending")
		}
	}
}

func TestMirror(t *testing.T) {
	t.Log("Testing: Mirror func interface")
	target := IntSlice{-8, -1, 1, 2, 2, 3, 4, 33}
	Mirror(&s)
	for i := range s {
		if (*&s)[i] != (*&target)[i] {
			t.Error("Mirror() order not as expected")
		}
	}
}

func TestLen(t *testing.T) {
	t.Log("Testing: Len func interface")
	n := Len(&s)
	if n != 8 {
		t.Error("Len() not expected value of 8: " + strconv.Itoa(n))
	}
}

func TestTruncateLeft(t *testing.T) {
	t.Log("Testing: TruncateLeft func interface")
	target := IntSlice{-8, -1, 1, 2, 2, 3}
	TruncateLeft(&s, 6)
	for i := range s {
		if (*&s)[i] != (*&target)[i] {
			t.Error("TruncateLeft() order not as expected")
		}
	}
}

func TestTruncateRight(t *testing.T) {
	t.Log("Testing: TruncateRight func interface")
	target := IntSlice{1, 2, 2, 3}
	TruncateRight(&s, 4)
	for i := range s {
		if (*&s)[i] != (*&target)[i] {
			t.Error("TruncateRight() order not as expected")
		}
	}
}

func TestSwap(t *testing.T) {
	t.Log("Testing: Swap func interface")
	target := IntSlice{1, 2, 3, 2}
	Swap(&s, 2, 3)
	for i := range s {
		if (*&s)[i] != (*&target)[i] {
			t.Error("Swap() order not as expected")
		}
	}
}
