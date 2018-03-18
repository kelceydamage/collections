package collections

import (
	"strconv"
	"testing"
)

var strs = StringSlice{"hi", "how", "are", "you", "hello", "world", "kitty", "cat"}

// BoolSlice tests
//-------------------------------------------------------------------------------------------------- <-100
func TestStringSliceIndex(t *testing.T) {
	t.Log("Testing: StringSlice for index")
	n := strs.Index("hello")
	if n != 4 {
		t.Error("Index() not expected value of 4: " + strconv.Itoa(n))
	}
}

func TestStringSliceIndexFail(t *testing.T) {
	t.Log("Testing: StringSlice for index failure")
	n := strs.Index("fred")
	if n != -1 {
		t.Error("Index() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestStringSliceIndexRight(t *testing.T) {
	t.Log("Testing: StringSlice for IndexRight")
	n := strs.IndexRight("you")
	if n != 3 {
		t.Error("IndexRight() not expected value of 3: " + strconv.Itoa(n))
	}
}

func TestStringSliceIndexRightFail(t *testing.T) {
	t.Log("Testing: StringSlice for IndexRight failure")
	n := strs.IndexRight("fred")
	if n != -1 {
		t.Error("IndexRight() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestStringSliceLen(t *testing.T) {
	t.Log("Testing: StringSlice for Len")
	n := strs.Len()
	if n != 8 {
		t.Error("Len() not expected value of 8: " + strconv.Itoa(n))
	}
}

func TestStringSliceSort(t *testing.T) {
	t.Log("Testing: StringSlice for Sort ascending")
	sorted := StringSlice{"are", "cat", "hello", "hi", "how", "kitty", "world", "you"}
	strs.Sort()
	t.Log("Strings: ", strs)
	t.Log("Target: ", sorted)
	for i := range strs {
		if strs[i] != sorted[i] {
			t.Error("Sort() order not as expected, ascending")
		}
	}
}

func TestStringSliceReverse(t *testing.T) {
	t.Log("Testing: StringSlice for Sort descending")
	sorted := StringSlice{"you", "world", "kitty", "how", "hi", "hello", "cat", "are"}
	strs.Reverse()
	t.Log("Strings: ", strs)
	t.Log("Target: ", sorted)
	for i := range strs {
		if strs[i] != sorted[i] {
			t.Error("Reverse() order not as expected, descending")
		}
	}
}

func TestStringSliceMirror(t *testing.T) {
	t.Log("Testing: StringSlice for Mirror")
	sorted := StringSlice{"are", "cat", "hello", "hi", "how", "kitty", "world", "you"}
	strs.Mirror()
	t.Log("Strings: ", strs)
	t.Log("Target: ", sorted)
	for i := range strs {
		if strs[i] != sorted[i] {
			t.Error("Mirror() element positions do not match expected positions")
		}
	}
}

func TestStringSliceTruncateLeft(t *testing.T) {
	t.Log("Testing: StringSlice for TruncateLeft")
	target := StringSlice{"are", "cat", "hello", "hi", "how", "kitty"}
	strs.TruncateLeft(6)
	if strs.Len() != 6 {
		t.Error("TruncateLeft() not expected value length of 6: " + strconv.Itoa(strs.Len()))
	}
	for i := range strs {
		if strs[i] != target[i] {
			t.Error("TruncateLeft() not expected values & positions")
		}
	}
}

func TestStringSliceTruncateRight(t *testing.T) {
	t.Log("Testing: StringSlice for TruncateRight")
	target := StringSlice{"hello", "hi", "how", "kitty"}
	strs.TruncateRight(4)
	if strs.Len() != 4 {
		t.Error("TruncateRight() not expected value length of 4: " + strconv.Itoa(strs.Len()))
	}
	for i := range strs {
		if strs[i] != target[i] {
			t.Error("TruncateRight() not expected values & positions")
		}
	}
}

func TestStringSliceLess(t *testing.T) {
	t.Log("Testing: StringSlice for Less")
	n := strs.Less(1, 2)
	if n != true {
		t.Error("Less() not expected value of true: " + strconv.FormatBool(n))
	}
}

func TestStringSliceSwap(t *testing.T) {
	t.Log("Testing: StringSlice for Swap")
	target := StringSlice{"hello", "how", "hi", "kitty"}
	strs.Swap(1, 2)
	for i := range strs {
		if strs[i] != target[i] {
			t.Error("Swap() not expected values & positions")
		}
	}
}
