package collections

import (
	"strconv"
	"testing"
)

var bools = BoolSlice{true, true, false, true, false, false, false, true}

// BoolSlice tests
//-------------------------------------------------------------------------------------------------- <-100
func TestBoolSliceIndex(t *testing.T) {
	t.Log("Testing: BoolSlice for Index")
	n := bools.Index(false)
	if n != 2 {
		t.Error("Index() not expected value of 2: " + strconv.Itoa(n))
	}
}

func TestBoolSliceIndexFail(t *testing.T) {
	target := BoolSlice{false, false, false}
	t.Log("Testing: BoolSlice for Index failure")
	n := target.Index(true)
	if n != -1 {
		t.Error("Index() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestBoolSliceSliceIndexFail2(t *testing.T) {
	t.Log("Testing: BoolSlice for IndexRight")
	n := bools.Index("ted")
	if n != -1 {
		t.Error("Index() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestBoolSliceIndexRight(t *testing.T) {
	t.Log("Testing: BoolSlice for index")
	n := bools.IndexRight(false)
	if n != 6 {
		t.Error("IndexRight() not expected value of 6: " + strconv.Itoa(n))
	}
}

func TestBoolSliceIndexRightFail(t *testing.T) {
	target := BoolSlice{false, false, false}
	t.Log("Testing: BoolSlice for IndexRight failure")
	n := target.IndexRight(true)
	if n != -1 {
		t.Error("IndexRight() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestBoolSliceSliceIndexRightFail2(t *testing.T) {
	t.Log("Testing: FloatBoolSlice64Slice for IndexRight")
	n := bools.IndexRight("ted")
	if n != -1 {
		t.Error("IndexRight() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestBoolSliceLen(t *testing.T) {
	t.Log("Testing: BoolSlice for Len")
	n := bools.Len()
	if n != 8 {
		t.Error("Len() not expected value of 8: " + strconv.Itoa(n))
	}
}

func TestBoolSliceSort(t *testing.T) {
	t.Log("Testing: BoolSlice for Sort true-first")
	sorted := BoolSlice{true, true, true, true, false, false, false, false}
	bools.Sort()
	t.Log("Bools: ", bools)
	t.Log("Target: ", sorted)
	for i := range bools {
		if bools[i] != sorted[i] {
			t.Error("Sort() order not as expected, true before false")
		}
	}
}

func TestBoolSliceReverse(t *testing.T) {
	t.Log("Testing: BoolSlice for Sort false-first")
	sorted := BoolSlice{false, false, false, false, true, true, true, true}
	bools.Reverse()
	t.Log("Bools: ", bools)
	t.Log("Target: ", sorted)
	for i := range bools {
		if bools[i] != sorted[i] {
			t.Error("Reverse() order not as expected, false before true")
		}
	}
}

func TestBoolSliceMirror(t *testing.T) {
	t.Log("Testing: BoolSlice for Mirror")
	sorted := BoolSlice{true, true, true, true, false, false, false, false}
	bools.Mirror()
	t.Log("Bools: ", bools)
	t.Log("Target: ", sorted)
	for i := range bools {
		if bools[i] != sorted[i] {
			t.Error("Mirror() element positions do not match expected positions")
		}
	}
}

func TestBoolSliceTruncateLeft(t *testing.T) {
	t.Log("Testing: BoolSlice for TruncateLeft")
	target := BoolSlice{true, true, true, true, false, false}
	bools.TruncateLeft(6)
	if bools.Len() != 6 {
		t.Error("TruncateLeft() not expected value length of 6: " + strconv.Itoa(bools.Len()))
	}
	for i := range bools {
		if bools[i] != target[i] {
			t.Error("TruncateLeft() not expected values & positions")
		}
	}
}

func TestBoolSliceTruncateRight(t *testing.T) {
	t.Log("Testing: BoolSlice for TruncateRight")
	target := BoolSlice{true, true, false, false}
	bools.TruncateRight(4)
	if bools.Len() != 4 {
		t.Error("TruncateRight() not expected value length of 4: " + strconv.Itoa(bools.Len()))
	}
	for i := range bools {
		if bools[i] != target[i] {
			t.Error("TruncateRight() not expected values & positions")
		}
	}
}

func TestBoolSliceLess(t *testing.T) {
	t.Log("Testing: BoolSlice for Less")
	n := bools.Less(1, 2)
	if n != true {
		t.Error("Less() not expected value of true: " + strconv.FormatBool(n))
	}
}

func TestBoolSliceSwap(t *testing.T) {
	t.Log("Testing: BoolSlice for Swap")
	target := BoolSlice{true, false, true, false}
	bools.Swap(1, 2)
	for i := range bools {
		if bools[i] != target[i] {
			t.Error("Swap() not expected values & positions")
		}
	}
}
