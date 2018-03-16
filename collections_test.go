package collections

import (
	"strconv"
	"testing"
)

var ints = IntSlice{1, 4, 2, 5, 6, 8, 935, 34, 2, 34, 5, 96, 0}

var negints = IntSlice{-1, -4, -2, -5, -6, -8, -935, -34, -2, -34, -5, -96, 0}

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
		t.Error("MinNonZero() not expected value of -87.44527846153846: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestNegFltSliceIndex(t *testing.T) {
	t.Log("Testing: negative values FltSlice for Index")
	n := negflts.Index(-34.54)
	if n != 9 {
		t.Error("MinNonZero() not expected value of 9: " + strconv.Itoa(n))
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
		t.Error("MinNonZero() not expected value of -87.44527846153846: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestFltSliceIndex(t *testing.T) {
	t.Log("Testing: positive values FltSlice for Index")
	n := flts.Index(34.54)
	if n != 9 {
		t.Error("MinNonZero() not expected value of 9: " + strconv.Itoa(n))
	}
}

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
	if n != -87 {
		t.Error("MinNonZero() not expected value of -87: " + strconv.Itoa(n))
	}
}

func TestNegIntSliceIndex(t *testing.T) {
	t.Log("Testing: negative values IntSlice for Index")
	n := negints.Index(-34)
	if n != 7 {
		t.Error("MinNonZero() not expected value of 7: " + strconv.Itoa(n))
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
	if n != 87 {
		t.Error("MinNonZero() not expected value of 87: " + strconv.Itoa(n))
	}
}

func TestIntSliceIndex(t *testing.T) {
	t.Log("Testing: positive values IntSlice for Index")
	n := ints.Index(935)
	if n != 6 {
		t.Error("MinNonZero() not expected value of 7: " + strconv.Itoa(n))
	}
}
