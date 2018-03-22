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

import "math"

// Code
//-------------------------------------------------------------------------------------------------- <-100

// IntSliceM is a higher level slice for ints that includes math functions.
type IntSliceM struct {
	IntSlice
}

// Sum add all the values in the slice and returns the result.
func (s *IntSliceM) Sum() int {
	n := 0
	for _, v := range (*s).All() {
		n += v.(int)
	}
	return n
}

// Avg returns the average of all values in the slice. Expressed as a float64
func (s *IntSliceM) Avg() float64 {
	return float64(s.Sum()) / float64(s.Len())
}

// AvgNonZero the average of all values in the slice not accounting for zero values.
// Expressed as a float64
func (s *IntSliceM) AvgNonZero() float64 {
	c := 0
	l := 1
	for _, v := range (*s).All() {
		if v.(int) != 0 {
			c += v.(int)
			l++
		}
	}
	return float64(c) / float64(l)
}

// Min returns the smallest value [n] in the slice along with its index [k].
func (s *IntSliceM) Min() (int, int) {
	n := (*s).All()[0].(int)
	k := 0
	for i, v := range (*s).All() {
		if v.(int) < n {
			n = v.(int)
			k = i
		}
	}
	return n, k
}

// MinNonZero returns the smallest non-zero value [n] in the slice along with its index [k].
func (s *IntSliceM) MinNonZero() (int, int) {
	n, _ := s.Max()
	k := 0
	for i, v := range (*s).All() {
		if v.(int) < n && v.(int) != 0 {
			n = v.(int)
			k = i
		}
	}
	return n, k
}

// Max returns the largest value [n] in the slice along with its index [k].
func (s *IntSliceM) Max() (int, int) {
	n := (*s).All()[0].(int)
	k := 0
	for i, v := range (*s).All() {
		if v.(int) > n {
			n = v.(int)
			k = i
		}
	}
	return n, k
}

// MaxNonZero returns the largest non-zero value [n] in the slice along with its index [k].
// Intended for use with negative integers.
func (s *IntSliceM) MaxNonZero() (int, int) {
	n, _ := s.Min()
	k := 0
	for i, v := range (*s).All() {
		if v.(int) > n && v.(int) != 0 {
			n = v.(int)
			k = i
		}
	}
	return n, k
}

// Variance returns the variance of the integers in the slice. Expressed as a float64.
func (s *IntSliceM) Variance() float64 {
	n := 0.0
	mapR := func(n float64) float64 {
		for _, v := range (*s).All() {
			n += math.Pow(float64(v.(int))-(*s).Avg(), 2.0)
		}
		return n
	}
	variance := mapR(n) / float64((*s).Len())
	return variance
}

// StdDev returns the standard deviation of the integers in the slice. Expressed as a float64.
func (s *IntSliceM) StdDev() float64 {
	return math.Sqrt((*s).Variance())
}
