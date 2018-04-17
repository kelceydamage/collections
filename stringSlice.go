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

// StringSlice is a slice off string with common methods.
type StringSlice struct {
	slice
}

// Set initial values for the slice.
func (s *StringSlice) Set(x slice) {
	(*s).Set(x)
}

// Get returns the element at the given index.
//func (s *StringSlice) Get(x int) string {
//	return (*s).slice.Get(x).(string)
//} 

// New instantiates a new StringSlice based on the passed slice.
func (s *StringSlice) New(x slice) StringSlice {
	i := StringSlice{}
	i.slice = x
	return i
}

// Sort inplace sorts the slice ascending.
//
// Implements sort.Sort() method.
func (s *StringSlice) Sort() {
	sort.Sort(strSliceAsc{*s})
}

// Reverse inplace sorts the slice descending.
//
// Implements sort.Sort() method.
func (s *StringSlice) Reverse() {
	sort.Sort(strSliceDsc{*s})
}

type strSliceDsc struct{ StringSlice }

// Less implementation for sort. Return true if value at index [i] is greater then value at index [j].
func (s strSliceDsc) Less(i, j int) bool {
	return s.StringSlice.All()[i].(string) > s.StringSlice.All()[j].(string)
}

type strSliceAsc struct{ StringSlice }

// Less implementation for sort. Return true if value at index [i] is less then value at index [j].
func (s strSliceAsc) Less(i, j int) bool {
	return s.StringSlice.All()[i].(string) < s.StringSlice.All()[j].(string)
}

// Slice returns a new slice with the range of data between indexes [i] and [j] int the original.
func (s *StringSlice) Slice(i, j int) StringSlice {
	return StringSlice{(*s).All()[i:j]}
}
