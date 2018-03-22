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

// Float64Slice is a slice off float64 with common methods.
type Float64Slice struct {
	slice
}

// New instantiates a new Float64Slice based on the passed slice.
func (s *Float64Slice) New(x slice) Float64Slice {
	i := Float64Slice{}
	i.slice = x
	return i
}

// Sort inplace sorts the slice ascending.
//
// Implements sort.Sort() method.
func (s *Float64Slice) Sort() {
	sort.Sort(fltSliceAsc{*s})
}

// Reverse inplace sorts the slice descending.
//
// Implements sort.Sort() method.
func (s *Float64Slice) Reverse() {
	sort.Sort(fltSliceDsc{*s})
}

type fltSliceDsc struct{ Float64Slice }

// Less implementation for sort. Return true if value at index [i] is greater then value at index [j].
func (s fltSliceDsc) Less(i, j int) bool {
	return s.Float64Slice.All()[i].(float64) > s.Float64Slice.All()[j].(float64)
}

type fltSliceAsc struct{ Float64Slice }

// Less implementation for sort. Return true if value at index [i] is less then value at index [j].
func (s fltSliceAsc) Less(i, j int) bool {
	return s.Float64Slice.All()[i].(float64) < s.Float64Slice.All()[j].(float64)
}

// Slice returns a new slice with the range of data between indexes [i] and [j] int the original.
func (s *Float64Slice) Slice(i, j int) Float64Slice {
	return Float64Slice{(*s).All()[i:j]}
}
