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
//  IndexRgt(int)
//  Sum()
//  Variance()
//  StdDev()
//  Sort(bool)
//  Swap(i, j)
//  Less(i, j)
//  TruncateLft(int)
//  TruncateRgt(int)
//  Reverse()

//---------------------------------------------------------------------------------------------------- <-100

package collections

import (
	"math"
	"sort"
)

// Code
//---------------------------------------------------------------------------------------------------- <-100

// IntSlice is a slice of ints with common methods.
type IntSlice []int

// Index retrives the first index [i] from the left, for the provided value [j] (int).
// Returns -1 if index not found.
func (s *IntSlice) Index(j int) int {
	for i, v := range *s {
		if v == j {
			return i
		}
	}
	return -1
}

// IndexRgt retrives the first index [i] from the right, for the provided value [j] (int).
// Returns -1 if index not found.
func (s *IntSlice) IndexRgt(j int) int {
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
func (s *IntSlice) Sum() int {
	n := 0
	for _, v := range *s {
		n += v
	}
	return n
}

// Len returns the length of the slice
func (s IntSlice) Len() int {
	return len(s)
}

// Avg returns the average of all values in the slice. Expressed as a float64
func (s *IntSlice) Avg() float64 {
	return float64(s.Sum()) / float64(s.Len())
}

// Min returns the smallest value [n] in the slice along with its index [k].
func (s *IntSlice) Min() (int, int) {
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
func (s *IntSlice) MinNonZero() (int, int) {
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
func (s *IntSlice) Max() (int, int) {
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
func (s *IntSlice) MaxNonZero() (int, int) {
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
func (s *IntSlice) Variance() float64 {
	n := 0.0
	mapR := func(n float64) float64 {
		for _, v := range *s {
			n += math.Pow(float64(v)-(*s).Avg(), 2.0)
		}
		return n
	}
	variance := mapR(n) / float64((*s).Len())
	return variance
}

// StdDev returns the standard deviation of the integers in the slice. Expressed as a float64.
func (s *IntSlice) StdDev() float64 {
	return math.Sqrt((*s).Variance())
}

// Sort inplace sorts the slice. Passing [true] will sort ascending, while passing [false] will sort
// descending
func (s *IntSlice) Sort(b bool) {
	if b {
		sort.Sort(intSliceAsc{*s})
	} else {
		sort.Sort(intSliceDsc{*s})
	}
}

// Swap implementation for general sort. This will swap the position of 2 items in the slice.
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type intSliceDsc struct{ IntSlice }

// Less implementation for sort. Returnd true if value at index [i] is greater then value at index [j].
func (s intSliceDsc) Less(i, j int) bool {
	return s.IntSlice[i] > s.IntSlice[j]
}

type intSliceAsc struct{ IntSlice }

// Less implementation for sort. Returnd true if value at index [i] is less then value at index [j].
func (s intSliceAsc) Less(i, j int) bool {
	return s.IntSlice[i] < s.IntSlice[j]
}

// TruncateLft shrinks the slice to [n] amount of ints starting from the left.
func (s *IntSlice) TruncateLft(n int) {
	if n != -1 && n <= len((*s)) {
		(*s) = (*s)[:n]
	}
}

// TruncateRgt shrinks the slice to [n] amount of ints starting from the right.
func (s *IntSlice) TruncateRgt(n int) {
	if n != -1 && n <= len((*s)) {
		(*s) = (*s)[(*s).Len()-n:]
	}
}

// Reverse will swap the order of ints in the slice.
func (s *IntSlice) Reverse() {
	for i, j := 0, (*s).Len()-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
