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

var s = StringSlice{}

var testStrings = s.New(slice{"fred", "bob", "al", "ron", "ted"})

// IntSlice tests with negative integers
//---------------------------------------------------------------------------------------------------- <-100

func TestStringSliceNew(t *testing.T) {
	t.Log("Testing: New values slice")
	target := slice{"fred", "bob", "al", "ron", "ted", "rob"}
	newStrings := testStrings.New(slice{"fred", "bob", "al", "ron", "ted", "rob"})
	for i := range newStrings.All() {
		if newStrings.All()[i] != target[i] {
			t.Error("New() order and content not as expected")
		}
	}
}

func TestStringSliceSort(t *testing.T) {
	t.Log("Testing: Sort values slice")
	target := slice{"al", "bob", "fred", "ron", "ted"}
	testStrings.Sort()
	for i := range testStrings.All() {
		if testStrings.All()[i] != target[i] {
			t.Error("Sort() order and content not as expected")
		}
	}
}

func TestStringSliceReverse(t *testing.T) {
	t.Log("Testing: Reverse values slice")
	target := slice{"ted", "ron", "fred", "bob", "al"}
	testStrings.Reverse()
	for i := range testStrings.All() {
		if testStrings.All()[i] != target[i] {
			t.Error("Reverse() order and content not as expected")
		}
	}
}

func TestStringSliceSlice(t *testing.T) {
	t.Log("Testing: Slice values slice")
	target := slice{"ron", "fred", "bob"}
	newStrings := testStrings.Slice(1, 4)
	for i := range newStrings.All() {
		if newStrings.All()[i] != target[i] {
			t.Error("Slice() order and content not as expected")
		}
	}
}
