//---------------------------------------------------------------------------------------------------- <-100
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
//---------------------------------------------------------------------------------------------------- <-100

//---------------------------------------------------------------------------------------------------- <-100

package collections

import (
	"sort"
)

// Code
//---------------------------------------------------------------------------------------------------- <-100

// IntSlice is a slice of ints with common methods.
type IntSlice struct {
	slice
}

// Sort inplace sorts the slice ascending.
//
// Implements sort.Sort() method.
func (s *IntSlice) Sort() {
	sort.Sort(intSliceAsc{*s})
}

// Reverse inplace sorts the slice descending.
//
// Implements sort.Sort() method.
func (s *IntSlice) Reverse() {
	sort.Sort(intSliceDsc{*s})
}

type intSliceDsc struct{ IntSlice }

// Less implementation for sort. Return true if value at index [i] is greater then value at index [j].
func (s intSliceDsc) Less(i, j int) bool {
	return s.IntSlice.All()[i].(int) > s.IntSlice.All()[j].(int)
}

type intSliceAsc struct{ IntSlice }

// Less implementation for sort. Return true if value at index [i] is less then value at index [j].
func (s intSliceAsc) Less(i, j int) bool {
	return s.IntSlice.All()[i].(int) < s.IntSlice.All()[j].(int)
}

// Slice returns a new slice with the range of data between indexes [i] and [j] int the original.
func (s *IntSlice) Slice(i, j int) IntSlice {
	return IntSlice{(*s).All()[i:j]}
}
