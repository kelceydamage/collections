package collections

import (
	"strconv"
	"testing"
)

var ints = IntSlice{1, 4, 2, 5, 6, 8, 935, 34, 2, 34, 5, 96, 0}

var negints = IntSlice{-1, -4, -2, -5, -6, -8, -935, -34, -2, -34, -5, -96, 0}

// IntSlice tests with negative integers
//---------------------------------------------------------------------------------------------------- <-100

func TestNegIntSliceMax(t *testing.T) {
	t.Log("Testing: negative values IntSlice for Max")
	n, _ := negints.Max()
	if n != 0 {
		t.Error("Max() not expected value of 0: " + strconv.Itoa(n))
	}
}

func TestNegIntSliceMaxNonZero(t *testing.T) {
	t.Log("Testing: negative values IntSlice for MaxNonZero")
	n, _ := negints.MaxNonZero()
	if n != -1 {
		t.Error("MaxNonZero() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestNegIntSliceMin(t *testing.T) {
	t.Log("Testing: negative values IntSlice for Min")
	n, _ := negints.Min()
	if n != -935 {
		t.Error("Min() not expected value of -935: " + strconv.Itoa(n))
	}
}

func TestNegIntSliceMinNonZero(t *testing.T) {
	t.Log("Testing: negative values IntSlice for MinNonZero")
	n, _ := negints.MinNonZero()
	if n != -935 {
		t.Error("MinNonZero() not expected value of -935: " + strconv.Itoa(n))
	}
}

func TestNegIntSliceAvg(t *testing.T) {
	t.Log("Testing: negative values IntSlice for Avg")
	n := negints.Avg()
	if n != -87.07692307692308 {
		t.Error("Avg() not expected value of -87.07692307692308: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestNegIntSliceSum(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Sum")
	n := negints.Sum()
	if n != -1132 {
		t.Error("Sum() not expected value of -1132: " + strconv.Itoa(n))
	}
}

// IntSlice tests with positive integers
//---------------------------------------------------------------------------------------------------- <-100

func TestIntSliceMax(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Max")
	n, _ := ints.Max()
	if n != 935 {
		t.Error("Max() not expected value of 935: " + strconv.Itoa(n))
	}
}

func TestIntSliceMaxNonZero(t *testing.T) {
	t.Log("Testing: positive values IntSlice for MaxNonZero")
	n, _ := ints.MaxNonZero()
	if n != 935 {
		t.Error("MaxNonZero() not expected value of 935: " + strconv.Itoa(n))
	}
}

func TestIntSliceMin(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Min")
	n, _ := ints.Min()
	if n != 0 {
		t.Error("Min() not expected value of 0: " + strconv.Itoa(n))
	}
}

func TestIntSliceMinNonZero(t *testing.T) {
	t.Log("Testing: positive values IntSlice for MinNonZero")
	n, _ := ints.MinNonZero()
	if n != 1 {
		t.Error("MinNonZero() not expected value of 1: " + strconv.Itoa(n))
	}
}

func TestIntSliceAvg(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Avg")
	n := ints.Avg()
	if n != 87.07692307692308 {
		t.Error("Avg() not expected value of 87.07692307692308: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestIntSliceSum(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Sum")
	n := ints.Sum()
	if n != 1132 {
		t.Error("Sum() not expected value of 1132: " + strconv.Itoa(n))
	}
}

// IntSlice non value sensitive tests
//---------------------------------------------------------------------------------------------------- <-100

func TestIntSliceLen(t *testing.T) {
	t.Log("Testing: Values IntSlice for Len")
	n := ints.Len()
	if n != 13 {
		t.Error("Len() not expected value of 13: " + strconv.Itoa(n))
	}
}

func TestIntSliceIndex(t *testing.T) {
	t.Log("Testing: IntSlice for index")
	n := ints.Index(8)
	if n != 5 {
		t.Error("Index() not expected value of 5: " + strconv.Itoa(n))
	}
}

func TestIntSliceIndexFail(t *testing.T) {
	t.Log("Testing: IntSlice for index failure")
	n := ints.Index(47)
	if n != -1 {
		t.Error("Index() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestIntSliceIndexRight(t *testing.T) {
	t.Log("Testing: IntSlice for IndexRight")
	n := ints.IndexRight(2)
	if n != 8 {
		t.Error("IndexRight() not expected value of 8: " + strconv.Itoa(n))
	}
}

func TestIntSliceIndexRightFail(t *testing.T) {
	t.Log("Testing: IntSlice for IndexRight failure")
	n := ints.IndexRight(47)
	if n != -1 {
		t.Error("IndexRight() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestIntSliceSortTrue(t *testing.T) {
	t.Log("Testing: IntSlice for Sort ascending")
	sorted := IntSlice{0, 1, 2, 2, 4, 5, 5, 6, 8, 34, 34, 96, 935}
	ints.Sort(true)
	t.Log("Ints: ", ints)
	t.Log("Target: ", sorted)
	for i := range ints {
		if ints[i] != sorted[i] {
			t.Error("Sort() order not as expected, ascending")
		}
	}
}

func TestIntSliceSortFalse(t *testing.T) {
	t.Log("Testing: IntSlice for Sort descending")
	sorted := IntSlice{935, 96, 34, 34, 8, 6, 5, 5, 4, 2, 2, 1, 0}
	ints.Sort(false)
	t.Log("Ints: ", ints)
	t.Log("Target: ", sorted)
	for i := range ints {
		if ints[i] != sorted[i] {
			t.Error("Sort() order not as expected, descending")
		}
	}
}

func TestIntSliceReverse(t *testing.T) {
	t.Log("Testing: IntSlice for Reverse")
	sorted := IntSlice{0, 1, 2, 2, 4, 5, 5, 6, 8, 34, 34, 96, 935}
	ints.Reverse()
	t.Log("Ints: ", ints)
	t.Log("Target: ", sorted)
	for i := range ints {
		if ints[i] != sorted[i] {
			t.Error("Reverse() element positions do not match expected positions")
		}
	}
}

func TestIntSliceTruncateLeft(t *testing.T) {
	t.Log("Testing: IntSlice for TruncateLeft")
	target := IntSlice{0, 1, 2, 2, 4, 5}
	ints.TruncateLeft(6)
	if ints.Len() != 6 {
		t.Error("TruncateLeft() not expected value length of 6: " + strconv.Itoa(strs.Len()))
	}
	for i := range ints {
		if ints[i] != target[i] {
			t.Error("TruncateLeft() not expected values & positions")
		}
	}
}

func TestIntSliceTruncateRight(t *testing.T) {
	t.Log("Testing: IntSlice for TruncateRight")
	target := IntSlice{2, 2, 4, 5}
	ints.TruncateRight(4)
	if ints.Len() != 4 {
		t.Error("TruncateRight() not expected value length of 4: " + strconv.Itoa(strs.Len()))
	}
	for i := range ints {
		if ints[i] != target[i] {
			t.Error("TruncateRight() not expected values & positions")
		}
	}
}

func TestIntSliceLess(t *testing.T) {
	t.Log("Testing: IntSlice for Less")
	n := ints.Less(1, 2)
	if n != true {
		t.Error("Less() not expected value of true: " + strconv.FormatBool(n))
	}
}

func TestIntSliceSwap(t *testing.T) {
	t.Log("Testing: IntSlice for Swap")
	target := IntSlice{2, 4, 2, 5}
	ints.Swap(1, 2)
	for i := range ints {
		if ints[i] != target[i] {
			t.Error("Swap() not expected values & positions")
		}
	}
}
