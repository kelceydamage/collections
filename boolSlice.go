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
	"sort"
)

// Code
//-------------------------------------------------------------------------------------------------- <-100

// BoolSlice is a slice of booleans with common methods.
type BoolSlice struct {
	slice
}

// New instantiates a new BoolSlice based on the passed slice.
func (s *BoolSlice) New(x slice) BoolSlice {
	i := BoolSlice{}
	i.slice = x
	return i
}

// Sort inplace sorts the slice ascending.
//
// Implements sort.Sort() method.
func (s *BoolSlice) Sort() {
	sort.Sort(*s)
}

// Reverse inplace sorts the slice descending.
//
// Implements sort.Sort() method.
func (s *BoolSlice) Reverse() {
	sort.Sort(*s)
	s.Mirror()
}

// Less implementation for sort. Return true if value at index [i] not equal to [j]. This will sort
// the slice so that all true booleans appear before false booleans.
func (s BoolSlice) Less(i, j int) bool {
	return s.All()[i].(bool) != s.All()[j].(bool)
}

// Slice returns a new slice with the range of data between indexes [i] and [j] int the original.
func (s *BoolSlice) Slice(i, j int) BoolSlice {
	return BoolSlice{(*s).All()[i:j]}
}
