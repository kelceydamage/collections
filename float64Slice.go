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
//  Float64Slice
//
// Struct Methods:
//  Avg()
//  Min()
//  MinNonZero()
//  Max()
//  MaxNonZero()
//  Index(float64)
//  IndexRight(float64)
//  Sum()
//  Variance()
//  StdDev()
//  Sort(bool)
//  Swap(i, j)
//  Less(i, j)
//  TruncateLeft(int)
//  TruncateRight(int)
//  Reverse()

//---------------------------------------------------------------------------------------------------- <-100

package collections

import (
	"math"
	"sort"
)

// Code
//---------------------------------------------------------------------------------------------------- <-100

// Float64Slice is a slice off float64 with common methods.
type Float64Slice []float64

// Index retrives the first index [i] from the left, for the provided value [j] (int).
// Returns -1 if index not found.
func (s *Float64Slice) Index(j float64) int {
	for i, v := range *s {
		if v == j {
			return i
		}
	}
	return -1
}

// IndexRight retrives the first index [i] from the right, for the provided value [j] (float64).
// Returns -1 if index not found.
func (s *Float64Slice) IndexRight(j float64) int {
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

// Sum add all the values in the slice and returns the result.
func (s *Float64Slice) Sum() float64 {
	n := 0.0
	for _, v := range *s {
		n += v
	}
	return n
}

// Len returns the length of the slice
func (s Float64Slice) Len() int {
	return len(s)
}

// Avg returns the average of all values in the slice.
func (s *Float64Slice) Avg() float64 {
	return s.Sum() / float64(s.Len())
}

// Min returns the smallest value [n] in the slice along with its index [k].
func (s *Float64Slice) Min() (float64, int) {
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
func (s *Float64Slice) MinNonZero() (float64, int) {
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
func (s *Float64Slice) Max() (float64, int) {
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
func (s *Float64Slice) MaxNonZero() (float64, int) {
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

// Variance returns the variance of the integers in the slice. Expressed as a float64.
func (s *Float64Slice) Variance() float64 {
	n := 0.0
	mapR := func(n float64) float64 {
		for _, v := range *s {
			n += math.Pow(v-(*s).Avg(), 2.0)
		}
		return n
	}
	variance := mapR(n) / float64((*s).Len())
	return variance
}

// StdDev returns the standard deviation of the integers in the slice. Expressed as a float64.
func (s *Float64Slice) StdDev() float64 {
	return math.Sqrt((*s).Variance())
}

// Sort inplace sorts the slice. Passing [true] will sort ascending, while passing [false] will sort
// descending
//
// Implements sort.Sort() method.
func (s *Float64Slice) Sort(b bool) {
	if b {
		sort.Sort(fltSliceAsc{*s})
	} else {
		sort.Sort(fltSliceDsc{*s})
	}
}

// Swap implementation for general sort. This will swap the position of 2 items in the slice.
func (s Float64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less returns true if value at index [i] is less then value at index [j].
func (s Float64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

type fltSliceDsc struct{ Float64Slice }

// Less implementation for sort. Return true if value at index [i] is greater then value at index [j].
func (s fltSliceDsc) Less(i, j int) bool {
	return s.Float64Slice[i] > s.Float64Slice[j]
}

type fltSliceAsc struct{ Float64Slice }

// Less implementation for sort. Return true if value at index [i] is less then value at index [j].
func (s fltSliceAsc) Less(i, j int) bool {
	return s.Float64Slice[i] < s.Float64Slice[j]
}

// TruncateLeft shrinks the slice to [n] amount of float64s starting from the left.
func (s *Float64Slice) TruncateLeft(n int) {
	if n != -1 && n <= len((*s)) {
		(*s) = (*s)[:n]
	}
}

// TruncateRight shrinks the slice to [n] amount of float64s starting from the right.
func (s *Float64Slice) TruncateRight(n int) {
	if n != -1 && n <= len((*s)) {
		(*s) = (*s)[(*s).Len()-n:]
	}
}

// Reverse will swap the order of float64s in the slice.
func (s *Float64Slice) Reverse() {
	for i, j := 0, (*s).Len()-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
