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

var b = BoolSlice{}

var testBools = b.New(slice{true, false, false, true, false})

// IntSlice tests with negative integers
//---------------------------------------------------------------------------------------------------- <-100

func TestBoolSliceNew(t *testing.T) {
	t.Log("Testing: New values slice")
	target := slice{true, false, false, true, false, true}
	newBools := testBools.New(slice{true, false, false, true, false, true})
	for i := range newBools.All() {
		if newBools.All()[i] != target[i] {
			t.Error("New() order and content not as expected")
		}
	}
}

func TestBoolSliceSort(t *testing.T) {
	t.Log("Testing: Sort values slice")
	target := slice{false, false, false, true, true}
	testBools.Sort()
	for i := range testBools.All() {
		if testBools.All()[i] != target[i] {
			t.Error("Sort() order and content not as expected")
		}
	}
}

func TestBoolSliceReverse(t *testing.T) {
	t.Log("Testing: Reverse values slice")
	target := slice{true, true, false, false, false}
	testBools.Reverse()
	for i := range testBools.All() {
		if testBools.All()[i] != target[i] {
			t.Error("Reverse() order and content not as expected")
		}
	}
}

func TestBoolSliceSlice(t *testing.T) {
	t.Log("Testing: Slice values slice")
	target := slice{true, false, false}
	newBools := testBools.Slice(1, 4)
	for i := range newBools.All() {
		if newBools.All()[i] != target[i] {
			t.Error("Slice() order and content not as expected")
		}
	}
}
