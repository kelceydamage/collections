package collections

import (
	"strconv"
	"testing"
)

var flts = Float64Slice{1.0, 4.3, 2.5, 5.2, 6.567, 8.2, 935.342, 34.98762, 2.43, 34.54, 5.5, 96.222, 0}

var negflts = Float64Slice{-1.0, -4.3, -2.5, -5.2, -6.567, -8.2, -935.342, -34.98762, -2.43, -34.54, -5.5, -96.222, 0}

// Float64Slice tests with negative integers
//-------------------------------------------------------------------------------------------------- <-100
func TestNegFloat64SliceMax(t *testing.T) {
	t.Log("Testing: negative values Float64Slice for Max")
	n, _ := negflts.Max()
	if n != 0.0 {
		t.Error("Max() not expected value of 0: " + strconv.FormatFloat(n, 'f', 2, 64))
	}
}

func TestNegFloat64SliceMaxNonZero(t *testing.T) {
	t.Log("Testing: negative values Float64Slice for MaxNonZero")
	n, _ := negflts.MaxNonZero()
	if n != -1.0 {
		t.Error("MaxNonZero() not expected value of -1: " + strconv.FormatFloat(n, 'f', 2, 64))
	}
}

func TestNegFloat64SliceMin(t *testing.T) {
	t.Log("Testing: negative values Float64Slice for Min")
	n, _ := negflts.Min()
	if n != -935.342 {
		t.Error("Min() not expected value of -935.342: " + strconv.FormatFloat(n, 'f', 3, 64))
	}
}

func TestNegFloat64SliceMinNonZero(t *testing.T) {
	t.Log("Testing: negative values Float64Slice for MinNonZero")
	n, _ := negflts.MinNonZero()
	if n != -935.342 {
		t.Error("MinNonZero() not expected value of -935.342: " + strconv.FormatFloat(n, 'f', 3, 64))
	}
}

func TestNegFloat64SliceAvg(t *testing.T) {
	t.Log("Testing: negative values Float64Slice for Avg")
	n := negflts.Avg()
	if n != -87.44527846153846 {
		t.Error("Avg() not expected value of -87.44527846153846: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestNegFloat64SliceSum(t *testing.T) {
	t.Log("Testing: positive values Float64Slice for Sum")
	n := negflts.Sum()
	if n != -1136.78862 {
		t.Error("Sum() not expected value of -1136.78862: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

// Float64Slice tests with positive integers
//---------------------------------------------------------------------------------------------------- <-100

func TestFloat64SliceMax(t *testing.T) {
	t.Log("Testing: positive values Float64Slice for Max")
	n, _ := flts.Max()
	if n != 935.342 {
		t.Error("Max() not expected value of 935.342: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestFloat64SliceMaxNonZero(t *testing.T) {
	t.Log("Testing: positive values Float64Slice for MaxNonZero")
	n, _ := flts.MaxNonZero()
	if n != 935.342 {
		t.Error("MaxNonZero() not expected value of 935.342: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestFloat64SliceMin(t *testing.T) {
	t.Log("Testing: positive values Float64Slice for Min")
	n, _ := flts.Min()
	if n != 0.0 {
		t.Error("Min() not expected value of 0.0: " + strconv.FormatFloat(n, 'f', 3, 64))
	}
}

func TestFloat64SliceMinNonZero(t *testing.T) {
	t.Log("Testing: positive values Float64Slice for MinNonZero")
	n, _ := flts.MinNonZero()
	if n != 1.0 {
		t.Error("MinNonZero() not expected value of 1.0: " + strconv.FormatFloat(n, 'f', 3, 64))
	}
}

func TestFloat64SliceAvg(t *testing.T) {
	t.Log("Testing: positive values Float64Slice for Avg")
	n := flts.Avg()
	if n != 87.44527846153846 {
		t.Error("Avg() not expected value of -87.44527846153846: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

func TestFloat64SliceSum(t *testing.T) {
	t.Log("Testing: positive values Float64Slice for Sum")
	n := flts.Sum()
	if n != 1136.78862 {
		t.Error("Sum() not expected value of 1136.78862: " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}

// Float64Slice non value sensitive tests
//---------------------------------------------------------------------------------------------------- <-100

func TestFloat64SliceLen(t *testing.T) {
	t.Log("Testing: Values Float64Slice for Len")
	n := flts.Len()
	if n != 13 {
		t.Error("Len() not expected value of 13: " + strconv.Itoa(n))
	}
}

func TestFloat64SliceIndex(t *testing.T) {
	t.Log("Testing: Float64Slice for index")
	n := flts.Index(5.2)
	if n != 3 {
		t.Error("Index() not expected value of 5: " + strconv.Itoa(n))
	}
}

func TestFloat64SliceIndexFail(t *testing.T) {
	t.Log("Testing: Float64Slice for index failure")
	n := flts.Index(10)
	if n != -1 {
		t.Error("Index() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestFloat64SliceIndexRight(t *testing.T) {
	t.Log("Testing: Float64Slice for IndexRight")
	n := flts.IndexRight(5.5)
	if n != 10 {
		t.Error("IndexRight() not expected value of 8: " + strconv.Itoa(n))
	}
}

func TestFloat64SliceIndexRightFail(t *testing.T) {
	t.Log("Testing: Float64Slice for IndexRight failure")
	n := flts.IndexRight(10)
	if n != -1 {
		t.Error("IndexRight() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestFloat64SliceSortTrue(t *testing.T) {
	t.Log("Testing: Float64Slice for Sort ascending")
	sorted := Float64Slice{0, 1, 2.43, 2.5, 4.3, 5.2, 5.5, 6.567, 8.2, 34.54, 34.98762, 96.222, 935.342}
	flts.Sort(true)
	t.Log("Floats: ", flts)
	t.Log("Target: ", sorted)
	for i := range flts {
		if flts[i] != sorted[i] {
			t.Error("Sort() order not as expected, ascending")
		}
	}
}

func TestFloat64SliceSortFalse(t *testing.T) {
	t.Log("Testing: Float64Slice for Sort descending")
	sorted := Float64Slice{935.342, 96.222, 34.98762, 34.54, 8.2, 6.567, 5.5, 5.2, 4.3, 2.5, 2.43, 1, 0}
	flts.Sort(false)
	t.Log("Floats: ", flts)
	t.Log("Target: ", sorted)
	for i := range flts {
		if flts[i] != sorted[i] {
			t.Error("Sort() order not as expected, descending")
		}
	}
}

func TestFloat64SliceReverse(t *testing.T) {
	t.Log("Testing: Float64Slice for Reverse")
	sorted := Float64Slice{0, 1, 2.43, 2.5, 4.3, 5.2, 5.5, 6.567, 8.2, 34.54, 34.98762, 96.222, 935.342}
	flts.Reverse()
	t.Log("Floats: ", flts)
	t.Log("Target: ", sorted)
	for i := range flts {
		if flts[i] != sorted[i] {
			t.Error("Reverse() element positions do not match expected positions")
		}
	}
}
