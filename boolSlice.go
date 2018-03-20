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

import (
	"sort"
)

// Code
//-------------------------------------------------------------------------------------------------- <-100

// BoolSlice is a slice off string with common methods.
type BoolSlice []bool

// Index retrives the first index [i] from the left, for the provided value [j] (bool).
// Returns -1 if index not found.
func (s *BoolSlice) Index(j interface{}) int {
	_, err := j.(bool)
	if !err {
		return -1
	}
	for i, v := range *s {
		if v == j {
			return i
		}
	}
	return -1
}

// IndexRight retrives the first index [i] from the right, for the provided value [j] (bool).
// Returns -1 if index not found.
func (s *BoolSlice) IndexRight(j interface{}) int {
	_, err := j.(bool)
	if !err {
		return -1
	}
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

// Swap implementation for general sort. This will swap the position of 2 items in the slice.
func (s BoolSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less implementation for sort. Return true if value at index [i] not equal to [j]. This will sort
// the slice so that all true booleans appear before false booleans.
func (s BoolSlice) Less(i, j int) bool {
	return s[i] != s[j]
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

// Mirror will swap the order of bools in the slice.
func (s *BoolSlice) Mirror() {
	for i, j := 0, len((*s))-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

// Append will add a value of type bool, into the BoolSlice.
func (s *BoolSlice) Append(j interface{}) int {
	_, err := j.(bool)
	if !err {
		return -1
	}
	(*s) = append((*s), j.(bool))
	return 0
}
