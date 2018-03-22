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
	"testing"
)

var i = IntSlice{}

var testInts = i.New(slice{1, 4, 2, -5, 6, 18, 935, -134, 2, -34, -5, 96, 0})

// IntSlice tests with negative integers
//---------------------------------------------------------------------------------------------------- <-100

func TestIntSliceNew(t *testing.T) {
	t.Log("Testing: New values slice")
	target := slice{1, 4, 2, -5, 6, 18, 935, -134, 2, -34}
	newInts := testInts.New(slice{1, 4, 2, -5, 6, 18, 935, -134, 2, -34})
	for i := range newInts.All() {
		if newInts.All()[i] != target[i] {
			t.Error("New() order and content not as expected")
		}
	}
}

func TestIntSliceSort(t *testing.T) {
	t.Log("Testing: Sort values slice")
	target := slice{-134, -34, -5, -5, 0, 1, 2, 2, 4, 6, 18, 96, 935}
	testInts.Sort()
	for i := range testInts.All() {
		if testInts.All()[i] != target[i] {
			t.Error("Sort() order and content not as expected")
		}
	}
}

func TestIntSliceReverse(t *testing.T) {
	t.Log("Testing: Reverse values slice")
	target := slice{935, 96, 18, 6, 4, 2, 2, 1, 0, -5, -5, -34, -134}
	testInts.Reverse()
	for i := range testInts.All() {
		if testInts.All()[i] != target[i] {
			t.Error("Reverse() order and content not as expected")
		}
	}
}

func TestIntSliceSlice(t *testing.T) {
	t.Log("Testing: Slice values slice")
	target := slice{96, 18, 6, 4}
	newInts := testInts.Slice(1, 4)
	for i := range newInts.All() {
		if newInts.All()[i] != target[i] {
			t.Error("Slice() order and content not as expected")
		}
	}
}
