package collections

import (
	"strconv"
	"testing"
)

var flts = FltSlice{1.0, 4.3, 2.5, 5.2, 6.567, 8.2, 935.342, 34.98762, 2.43, 34.54, 5.5, 96.222, 0}

var negflts = FltSlice{-1.0, -4.3, -2.5, -5.2, -6.567, -8.2, -935.342, -34.98762, -2.43, -34.54, -5.5, -96.222, 0}

// FltSlice tests with negative integers
//---------------------------------------------------------------------------------------------------- <-100
func TestNegFltSliceMax(t *testing.T) {
	t.Log("Testing: negative values FltSlice for Max")
	n, _ := negflts.Max()
	if n != 0.0 {
		t.Error("Max() not expected value of 0: " + strconv.FormatFloat(n, 'f', 2, 64))
	}
}

func TestNegFltSliceMaxNonZero(t *testing.T) {
	t.Log("Testing: negative values FltSlice for MaxNonZero")
	n, _ := negflts.MaxNonZero()
	if n != -1.0 {
		t.Error("MaxNonZero() not expected value of -1: " + strconv.FormatFloat(n, 'f', 2, 64))
	}
}

func TestNegFltSliceMin(t *testing.T) {
	t.Log("Testing: negative values FltSlice for Min")
	n, _ := negflts.Min()
	if n != -935.342 {
		t.Error("Min() not expected value of -935.342: " + strconv.FormatFloat(n, 'f', 3, 64))
	}
}

func TestNegFltSliceMinNonZero(t *testing.T) {
	t.Log("Testing: negative values FltSlice for MinNonZero")
	n, _ := negflts.MinNonZero()
	if n != -935.342 {
		t.Error("MinNonZero() not expected value of -935.342: " + strconv.FormatFloat(n, 'f', 3, 64))
	}
}

func TestNegFltSliceAvg(t *testing.T) {
	t.Log("Testing: negative values FltSlice for Avg")
	n := negflts.Avg()
	if n != -87.44527846153846 {
		t.Error("Avg() not expected value of -87.44527846153846: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestNegFltSliceIndex(t *testing.T) {
	t.Log("Testing: negative values FltSlice for Index")
	n := negflts.Index(-34.54)
	if n != 9 {
		t.Error("Index() not expected value of 9: " + strconv.Itoa(n))
	}
}

func TestNegFltSliceSum(t *testing.T) {
	t.Log("Testing: positive values FltSlice for Sum")
	n := negflts.Sum()
	if n != -1136.78862 {
		t.Error("Sum() not expected value of -1136.78862: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestNegFltSliceLen(t *testing.T) {
	t.Log("Testing: positive values FltSlice for Len")
	n := negflts.Len()
	if n != 13 {
		t.Error("Len() not expected value of 13: " + strconv.Itoa(n))
	}
}

// FltSlice tests with positive integers
//---------------------------------------------------------------------------------------------------- <-100

func TestFltSliceMax(t *testing.T) {
	t.Log("Testing: positive values FltSlice for Max")
	n, _ := flts.Max()
	if n != 935.342 {
		t.Error("Max() not expected value of 935.342: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestFltSliceMaxNonZero(t *testing.T) {
	t.Log("Testing: positive values FltSlice for MaxNonZero")
	n, _ := flts.MaxNonZero()
	if n != 935.342 {
		t.Error("MaxNonZero() not expected value of 935.342: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestFltSliceMin(t *testing.T) {
	t.Log("Testing: positive values FltSlice for Min")
	n, _ := flts.Min()
	if n != 0.0 {
		t.Error("Min() not expected value of 0.0: " + strconv.FormatFloat(n, 'f', 3, 64))
	}
}

func TestFltSliceMinNonZero(t *testing.T) {
	t.Log("Testing: positive values FltSlice for MinNonZero")
	n, _ := flts.MinNonZero()
	if n != 1.0 {
		t.Error("MinNonZero() not expected value of 1.0: " + strconv.FormatFloat(n, 'f', 3, 64))
	}
}

func TestFltSliceAvg(t *testing.T) {
	t.Log("Testing: positive values FltSlice for Avg")
	n := flts.Avg()
	if n != 87.44527846153846 {
		t.Error("Avg() not expected value of -87.44527846153846: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestFltSliceIndex(t *testing.T) {
	t.Log("Testing: positive values FltSlice for Index")
	n := flts.Index(34.54)
	if n != 9 {
		t.Error("Index() not expected value of 9: " + strconv.Itoa(n))
	}
}

func TestFltSliceSum(t *testing.T) {
	t.Log("Testing: positive values FltSlice for Sum")
	n := flts.Sum()
	if n != 1136.78862 {
		t.Error("Sum() not expected value of 1136.78862: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestFltSliceLen(t *testing.T) {
	t.Log("Testing: positive values FltSlice for Len")
	n := flts.Len()
	if n != 13 {
		t.Error("Len() not expected value of 13: " + strconv.Itoa(n))
	}
}
