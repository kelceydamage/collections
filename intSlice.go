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

# Doc
#---------------------------------------------------------------------------------------------------- <-100

// IntSlice is a slice on ints with common methods.
type IntSlice []int

// Index retrives the first index for the provided value [j] (int).
func (s *IntSlice) Iindex(j int) int {
	for i, v := range *s {
		if v == j {
			return i
		}
	}
	return -1
}

// Avg returns the average of all values in the slice.
func (s *IntSlice) Avg() int {
	n := 0
	for i := range *s {
		n += (*s)[i]
	}
	return n / len((*s))
}

// Min returns the smallest value in the slice.
func (s *IntSlice) Min() (int, int) {
	n := (*s)[0]
	k := 0
	for i := range *s {
		if (*s)[i] < n {
			n = (*s)[i]
			k = i
		}
	}
	return k, n
}

// MinNonZero returns the smallest non-zero value in the slice.
func (s *IntSlice) MinNonZero() int {
	n := s.Max()
	for i := range *s {
		if (*s)[i] != 0 {
			if (*s)[i] < n {
				n = (*s)[i]
			}
		}
	}
	return n
}

// Max returns the largest value in the slice.
func (s *IntSlice) Max() int {
	n := (*s)[0]
	for i := range *s {
		if (*s)[i] > n {
			n = (*s)[i]
		}
	}
	return n
}