//-------------------------------------------------------------------------------------------------- <-100
// Author: Kelcey Damage
// Go: 1.10

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Doc
//-------------------------------------------------------------------------------------------------- <-100

// Doc (90 char length for optimal godoc code-block parsing)                              | <- 90
//-------------------------------------------------------------------------------------------------- <-100

package collections

import (
	"strconv"
	"testing"
)

var ints = slice{1, 4, 2, -5, 6, 18, 935, -134, 2, -34, -5, 96, 0}

// IntSlice tests with negative integers
//---------------------------------------------------------------------------------------------------- <-100

func TestSliceSet(t *testing.T) {
	t.Log("Testing: Set values slice")
	ints.Set(slice{1, 4, 2, -5, 6, 18, 935, -134, 2, -34, -5, 96, 0})
	if ints.Len() != 13 {
		t.Error("Set() Len not expected value of 14: " + strconv.Itoa(ints.Len()))
	}
}

func TestSliceGet(t *testing.T) {
	t.Log("Testing: Get values slice")
	n := ints.Get(2)
	if n.(int) != 2 {
		t.Error("Get() not expected value of 2: " + strconv.Itoa(n.(int)))
	}
}

func TestSliceAppend(t *testing.T) {
	t.Log("Testing: Append values slice")
	ints.Append(2)
	if ints.Len() != 14 {
		t.Error("Append() Len not expected value of 14: " + strconv.Itoa(ints.Len()))
	}
}

func TestSliceAll(t *testing.T) {
	t.Log("Testing: All values slice")
	target := slice{1, 4, 2, -5, 6, 18, 935, -134, 2, -34, -5, 96, 0, 2}
	contents := ints
	for i := range contents {
		if contents[i] != target[i] {
			t.Error("All() order and content not as expected")
		}
	}
}

func TestSliceLen(t *testing.T) {
	t.Log("Testing: Len values slice")
	if ints.Len() != 14 {
		t.Error("Len() not expected value of 14: " + strconv.Itoa(ints.Len()))
	}
}

func TestSliceOverwrite(t *testing.T) {
	t.Log("Testing: Overwrite values slice")
	target := slice{0, 14, 1, 4, 2, -5, 6, 18, 935, -134, 2, -34, -5, 96}
	contents := ints
	ints.Overwrite(target)
	for i := range contents {
		if ints[i] != target[i] {
			t.Error("Overwrite() order and content not as expected")
		}
	}
}

func TestSliceSwap(t *testing.T) {
	t.Log("Testing: Swap values slice")
	target := slice{14, 0, 1, 4, 2, -5, 6, 18, 935, -134, 2, -34, -5, 96}
	contents := ints
	ints.Swap(0, 1)
	for i := range contents {
		if ints[i] != target[i] {
			t.Error("Swap() order and content not as expected")
		}
	}
}

func TestSliceTruncateLeft(t *testing.T) {
	t.Log("Testing: TruncateLeft values slice")
	target := slice{14, 0, 1, 4, 2, -5}
	ints.TruncateLeft(6)
	for i := range ints {
		if ints[i] != target[i] {
			t.Error("TruncateLeft() order and content not as expected")
		}
	}
}

func TestSliceTruncateRight(t *testing.T) {
	t.Log("Testing: TruncateRight values slice")
	target := slice{1, 4, 2, -5}
	ints.TruncateRight(4)
	for i := range ints {
		if ints[i] != target[i] {
			t.Error("TruncateRight() order and content not as expected")
		}
	}
}

func TestSliceMirror(t *testing.T) {
	t.Log("Testing: Mirror values slice")
	target := slice{-5, 2, 4, 1}
	contents := ints
	ints.Mirror()
	for i := range contents {
		if ints[i] != target[i] {
			t.Error("Mirror() order and content not as expected")
		}
	}
}

func TestSliceIndex(t *testing.T) {
	t.Log("Testing: Index values slice")
	n := ints.Index(2)
	if n != 1 {
		t.Error("Index() not expected value of 1: " + strconv.Itoa(n))
	}
}

func TestSliceIndexFail(t *testing.T) {
	t.Log("Testing: IndexFail values slice")
	n := ints.Index(33)
	if n != -1 {
		t.Error("IndexFail() not expected value of -1: " + strconv.Itoa(n))
	}
}

func TestSliceIndexRight(t *testing.T) {
	t.Log("Testing: IndexRight values slice")
	n := ints.IndexRight(4)
	if n != 2 {
		t.Error("IndexRight() not expected value of 2: " + strconv.Itoa(n))
	}
}

func TestSliceIndexRightFail(t *testing.T) {
	t.Log("Testing: IndexRightFail values slice")
	n := ints.IndexRight(33)
	if n != -1 {
		t.Error("IndexRightFail() not expected value of -1: " + strconv.Itoa(n))
	}
}
