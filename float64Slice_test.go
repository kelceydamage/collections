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

var f = Float64Slice{}

var testFloats = f.New(slice{0.0, 1.5, 4.0, 2.4, -5.6, 6.2, -18.33})

// IntSlice tests with negative integers
//---------------------------------------------------------------------------------------------------- <-100

func TestFloat64SliceNew(t *testing.T) {
	t.Log("Testing: New values slice")
	target := slice{0.0, 1.5, 4.0, 2.4, -5.6, 6.2, -18.33}
	newFloats := testFloats.New(slice{0.0, 1.5, 4.0, 2.4, -5.6, 6.2, -18.33})
	for i := range newFloats.All() {
		if newFloats.All()[i] != target[i] {
			t.Error("New() order and content not as expected")
		}
	}
}

func TestFloat64SliceSort(t *testing.T) {
	t.Log("Testing: Sort values slice")
	target := slice{-18.33, -5.6, 0.0, 1.5, 2.4, 4.0, 6.2}
	testFloats.Sort()
	for i := range testFloats.All() {
		if testFloats.All()[i] != target[i] {
			t.Error("Sort() order and content not as expected")
		}
	}
}

func TestFloat64SliceReverse(t *testing.T) {
	t.Log("Testing: Reverse values slice")
	target := slice{6.2, 4.0, 2.4, 1.5, 0.0, -5.6, -18.33}
	testFloats.Reverse()
	for i := range testFloats.All() {
		if testFloats.All()[i] != target[i] {
			t.Error("Reverse() order and content not as expected")
		}
	}
}

func TestFloat64SliceSlice(t *testing.T) {
	t.Log("Testing: Slice values slice")
	target := slice{4.0, 2.4, 1.5, 0.0}
	newFloats := testFloats.Slice(1, 4)
	for i := range newFloats.All() {
		if newFloats.All()[i] != target[i] {
			t.Error("Slice() order and content not as expected")
		}
	}
}
