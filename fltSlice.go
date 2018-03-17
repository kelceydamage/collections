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

// Struct Type:
//  FltSlice
//
// Struct Methods:
//  Avg()
//  Min()
//  MinNonZero()
//  Max()
//  MaxNonZero()
//  Index(int)

//---------------------------------------------------------------------------------------------------- <-100

package collections

// Code
//---------------------------------------------------------------------------------------------------- <-100

// FltSlice is a slice off float64 with common methods.
type FltSlice []float64

// Index retrives the first index [i] from the left, for the provided value [j] (int).
func (s *FltSlice) Index(j float64) int {
	for i, v := range *s {
		if v == j {
			return i
		}
	}
	return -1
}

// Sum add all the values in the slice and returns the result.
func (s *FltSlice) Sum() float64 {
	n := 0.0
	for _, v := range *s {
		n += v
	}
	return n
}

// Len returns the length of the slice
func (s *FltSlice) Len() int {
	return len((*s))
}

// Avg returns the average of all values in the slice.
func (s *FltSlice) Avg() float64 {
	return s.Sum() / float64(s.Len())
}

// Min returns the smallest value [n] in the slice along with its index [k].
func (s *FltSlice) Min() (float64, int) {
	n := (*s)[0]
	k := 0
	for i, v := range *s {
		if v < n {
			n = v
			k = i
		}
	}
	return n, k
}

// MinNonZero returns the smallest non-zero value [n] in the slice along with its index [k].
func (s *FltSlice) MinNonZero() (float64, int) {
	n, _ := s.Max()
	k := 0
	for i, v := range *s {
		if v != 0 {
			if v < n {
				n = v
				k = i
			}
		}
	}
	return n, k
}

// Max returns the largest value [n] in the slice along with its index [k].
func (s *FltSlice) Max() (float64, int) {
	n := (*s)[0]
	k := 0
	for i, v := range *s {
		if v > n {
			n = v
			k = i
		}
	}
	return n, k
}

// MaxNonZero returns the largest non-zero value [n] in the slice along with its index [k].
// Intended for use with negative integers.
func (s *FltSlice) MaxNonZero() (float64, int) {
	n, _ := s.Min()
	k := 0
	for i, v := range *s {
		if v != 0 {
			if v > n {
				n = v
				k = i
			}
		}
	}
	return n, k
}
