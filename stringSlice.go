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
//  StringSlice
//
// Struct Methods:
//  len()
//  Index(string)
//  IndexRight(string)
//  Sort()
//  Swap(i, j)
//  Less(i, j)
//  TruncateLeft(int)
//  TruncateRight(int)
//  Reverse()
//  Mirror()

// Doc (90 char length for optimal godoc code-block parsing)                              | <- 90
//-------------------------------------------------------------------------------------------------- <-100

package collections

import "sort"

// Code
//-------------------------------------------------------------------------------------------------- <-100

// StringSlice is a slice off string with common methods.
type StringSlice []string

// Index retrives the first index [i] from the left, for the provided value [j] (string).
// Returns -1 if index not found.
func (s *StringSlice) Index(j string) int {
	for i, v := range *s {
		if v == j {
			return i
		}
	}
	return -1
}

// IndexRight retrives the first index [i] from the right, for the provided value [j] (string).
// Returns -1 if index not found.
func (s *StringSlice) IndexRight(j string) int {
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
func (s StringSlice) Len() int {
	return len(s)
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

// Swap implementation for general sort. This will swap the position of 2 items in the slice.
func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less returns true if value at index [i] is less then value at index [j].
func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

type strSliceDsc struct{ StringSlice }

// Less implementation for sort. Return true if value at index [i] is greater then value at index [j].
func (s strSliceDsc) Less(i, j int) bool {
	return s.StringSlice[i] > s.StringSlice[j]
}

type strSliceAsc struct{ StringSlice }

// Less implementation for sort. Return true if value at index [i] is less then value at index [j].
func (s strSliceAsc) Less(i, j int) bool {
	return s.StringSlice[i] < s.StringSlice[j]
}

// TruncateLeft shrinks the slice to [n] amount of float64s starting from the left.
func (s *StringSlice) TruncateLeft(n int) {
	if n != -1 && n <= len((*s)) {
		(*s) = (*s)[:n]
	}
}

// TruncateRight shrinks the slice to [n] amount of float64s starting from the right.
func (s *StringSlice) TruncateRight(n int) {
	if n != -1 && n <= len((*s)) {
		(*s) = (*s)[(*s).Len()-n:]
	}
}

// Mirror will swap the order of ints in the slice.
func (s *StringSlice) Mirror() {
	for i, j := 0, (*s).Len()-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
