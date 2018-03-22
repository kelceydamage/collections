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
	"math"
)

// Code
//---------------------------------------------------------------------------------------------------- <-100

// Float64SliceM is a higher level slice for float64 that includes math functions.
type Float64SliceM struct {
	Float64Slice
}

// Sum add all the values in the slice and returns the result.
func (s *Float64SliceM) Sum() float64 {
	n := 0.0
	for _, v := range (*s).All() {
		n += v.(float64)
	}
	return n
}

// Avg returns the average of all values in the slice.
func (s *Float64SliceM) Avg() float64 {
	return s.Sum() / float64(s.Len())
}

// AvgNonZero the average of all values in the slice not accounting for zero values.
// Expressed as a float64
func (s *Float64SliceM) AvgNonZero() float64 {
	c := 0.0
	l := 1
	for _, v := range (*s).All() {
		if v.(float64) != 0 {
			c += v.(float64)
			l++
		}
	}
	return c / float64(l)
}

// Min returns the smallest value [n] in the slice along with its index [k].
func (s *Float64SliceM) Min() (float64, int) {
	n := (*s).All()[0].(float64)
	k := 0
	for i, v := range (*s).All() {
		if v.(float64) < n {
			n = v.(float64)
			k = i
		}
	}
	return n, k
}

// MinNonZero returns the smallest non-zero value [n] in the slice along with its index [k].
func (s *Float64SliceM) MinNonZero() (float64, int) {
	n, _ := s.Max()
	k := 0
	for i, v := range (*s).All() {
		if v.(float64) < n && v.(float64) != 0 {
			n = v.(float64)
			k = i
		}
	}
	return n, k
}

// Max returns the largest value [n] in the slice along with its index [k].
func (s *Float64SliceM) Max() (float64, int) {
	n := (*s).All()[0].(float64)
	k := 0
	for i, v := range (*s).All() {
		if v.(float64) > n {
			n = v.(float64)
			k = i
		}
	}
	return n, k
}

// MaxNonZero returns the largest non-zero value [n] in the slice along with its index [k].
// Intended for use with negative integers.
func (s *Float64SliceM) MaxNonZero() (float64, int) {
	n, _ := s.Min()
	k := 0
	for i, v := range (*s).All() {
		if v.(float64) > n && v.(float64) != 0 {
			n = v.(float64)
			k = i
		}
	}
	return n, k
}

// Variance returns the variance of the integers in the slice. Expressed as a float64.
func (s *Float64SliceM) Variance() float64 {
	n := 0.0
	mapR := func(n float64) float64 {
		for _, v := range (*s).All() {
			n += math.Pow(v.(float64)-(*s).Avg(), 2.0)
		}
		return n
	}
	variance := mapR(n) / float64((*s).Len())
	return variance
}

// StdDev returns the standard deviation of the integers in the slice. Expressed as a float64.
func (s *Float64SliceM) StdDev() float64 {
	return math.Sqrt((*s).Variance())
}
