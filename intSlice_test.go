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

func TestNegIntSliceIndex(t *testing.T) {
	t.Log("Testing: negative values IntSlice for Index")
	n := negints.Index(-34)
	if n != 7 {
		t.Error("Index() not expected value of 7: " + strconv.Itoa(n))
	}
}

func TestNegIntSliceSum(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Sum")
	n := negints.Sum()
	if n != -1132 {
		t.Error("Sum() not expected value of -1132: " + strconv.Itoa(n))
	}
}

func TestNegIntSliceLen(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Len")
	n := negints.Len()
	if n != 13 {
		t.Error("Len() not expected value of 13: " + strconv.Itoa(n))
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

func TestIntSliceIndex(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Index")
	n := ints.Index(935)
	if n != 6 {
		t.Error("Index() not expected value of 7: " + strconv.Itoa(n))
	}
}

func TestIntSliceSum(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Sum")
	n := ints.Sum()
	if n != 1132 {
		t.Error("Sum() not expected value of 1132: " + strconv.Itoa(n))
	}
}

func TestIntSliceLen(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Len")
	n := ints.Len()
	if n != 13 {
		t.Error("Len() not expected value of 13: " + strconv.Itoa(n))
	}
}
