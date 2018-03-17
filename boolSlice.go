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
//---------------------------------------------------------------------------------------------------- <-100

// Struct Type:
//  BoolSlice
//
// Struct Methods:
//  len()
//  Index(string)
//  IndexRight(string)
//  Sort(bool)
//  Swap(i, j)
//  Less(i, j)
//  TruncateLeft(int)
//  TruncateRight(int)
//  Reverse()

// Doc (90 char length for optimal godoc code-block parsing)                              | <- 90
//-------------------------------------------------------------------------------------------------- <-100

package collections

import "sort"

// Code
//-------------------------------------------------------------------------------------------------- <-100

// BoolSlice is a slice off string with common methods.
type BoolSlice []bool

// Index retrives the first index [i] from the left, for the provided value [j] (bool).
// Returns -1 if index not found.
func (s *BoolSlice) Index(j bool) int {
	for i, v := range *s {
		if v == j {
			return i
		}
	}
	return -1
}

// IndexRight retrives the first index [i] from the right, for the provided value [j] (bool).
// Returns -1 if index not found.
func (s *BoolSlice) IndexRight(j bool) int {
	k := -1
	for i, v := range *s {
		if v == j {
			k = i
		}
	}
	if k == -1 {
		return k
	}
	return k
}

// Len returns the length of the slice
func (s BoolSlice) Len() int {
	return len(s)
}

// Sort inplace sorts the slice. Passing [true] will sort ascending, while passing [false] will sort
// descending.
//
// Implements sort.Sort() method.
func (s *BoolSlice) Sort(b bool) {
	if b {
		sort.Sort(blnSliceAsc{*s})
	} else {
		sort.Sort(blnSliceDsc{*s})
	}
}

// Swap implementation for general sort. This will swap the position of 2 items in the slice.
func (s BoolSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less returns true if value at index [i] is equal then value at index [j].
func (s BoolSlice) Less(i, j int) bool {
	return s[i] == s[j]
}

type blnSliceDsc struct{ BoolSlice }

// Less implementation for sort. Return true if value at index [i] is greater then value at index [j].
func (s blnSliceDsc) Less(i, j int) bool {
	return s.BoolSlice[i] != s.BoolSlice[j]
}

type blnSliceAsc struct{ BoolSlice }

// Less implementation for sort. Return true if value at index [i] is less then value at index [j].
func (s blnSliceAsc) Less(i, j int) bool {
	return s.BoolSlice[i] == s.BoolSlice[j]
}

// TruncateLeft shrinks the slice to [n] amount of bools starting from the left.
func (s *BoolSlice) TruncateLeft(n int) {
	if n != -1 && n <= len((*s)) {
		(*s) = (*s)[:n]
	}
}

// TruncateRight shrinks the slice to [n] amount of bools starting from the right.
func (s *BoolSlice) TruncateRight(n int) {
	if n != -1 && n <= len((*s)) {
		(*s) = (*s)[(*s).Len()-n:]
	}
}

// Reverse will swap the order of bools in the slice.
func (s *BoolSlice) Reverse() {
	for i, j := 0, (*s).Len()-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
