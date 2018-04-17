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

package stats

import ( 
	"math"
	"fmt"

	"github.com/kelceydamage/collections"
)

// Code
//-------------------------------------------------------------------------------------------------- <-100

// Sum with high precision
func sum(s collections.Slice) int {
	n := 0
	for _, v := range s.All() {
		v, ok := v.(int)
		if !ok {
			n += ifaceToInt(v)
		} else {
			n += v
		}
	}
	return n
}

func sumFloat64(s collections.Slice) float64 {
	n := 0.0
	for _, v := range s.All() {
		v, ok := v.(float64)
		if !ok {
			n += ifaceToFloat64(v)
		} else {
			n += v
		}
	}
	return n
}

func ifaceToFloat64(x interface{}) float64 {
	v, ok := x.(float64)
	if !ok {
		return float64(x.(int))
	}
	return v
}

func ifaceToInt(x interface{}) int {
	v, ok := x.(int)
	if !ok {
		return int(x.(float64))
	}
	return v
}

func removeZeroes(s collections.Slice) collections.Slice {
	var l collections.Slice
	for _, v := range s.All() {
		if ifaceToFloat64(v) == 0.0 {
			l.Append(v) 
		}
	}
	return l
}

func less(n, v float64) bool {
	if v < n {
		return true
	}
	return false
}

func lessNonZero(n, v float64) bool {
	if v < n && v != 0 {
		return true
	}
	return false
}

func greater(n, v float64) bool {
	if v > n {
		return true
	}
	return false
}

func greaterNonZero(n, v float64) bool {
	if v > n && v != 0 {
		return true
	}
	return false
}

func compare(s collections.Slice, n float64, f func(float64, float64) bool) int {
	k := 0
	for i, v := range s.All() {
		if f(n, ifaceToFloat64(v)) {
			k = i
		}
	}
	return k
}

func varPowSum(s collections.Slice) float64 {
	n := 0.0
	for _, v := range s.All() {
		n += math.Pow(ifaceToFloat64(v)-Avg(s), 2.0)
	}
	return n
}

// Variance returns the variance of the values in the slice. Expressed as a float64.
func Variance(s collections.Slice) float64 {
	return varPowSum(s) / ifaceToFloat64(s.Len())
}

// StdDev returns the standard deviation of the integers in the slice. Expressed as a float64.
func stdDev(s collections.Slice) float64 {
	return math.Sqrt(Variance(s))
}

// Sum returns the sum of all objects in the slice as if they were ints.
func Sum(s collections.Slice) int {
	return sum(s)
}

// SumFloat64 returns the sum of all objects in the slice as if they were float64.
func SumFloat64(s collections.Slice) float64 {
	return sumFloat64(s)
}

// Avg returns the average of all objects in the slice.
func Avg(s collections.Slice) float64 {
	return SumFloat64(s) / float64(s.Len())
}

// AvgNonZero returns the average of all objects in the slice while skipping sero values for both sum and len.
func AvgNonZero(s collections.Slice) float64 {
	fmt.Println(s)
	return Avg(removeZeroes(s))
}

// Min returns the index of the rightmost lowest value in the slice, including zero.
func Min(s collections.Slice) int {
	return compare(s, ifaceToFloat64(s.All()[Max(s)]), less)
}

// MinNonZero returns the index of the rightmost lowest value in the slice, excluding zero.
func MinNonZero(s collections.Slice) int {
	return compare(s, ifaceToFloat64(s.All()[Max(s)]), lessNonZero)
}

// Max returns the index of the rightmost highest value in the slice, including zero.
func Max(s collections.Slice) int {
	return compare(s, ifaceToFloat64(s.All()[0]), greater)
}

// MaxNonZero returns the index of the rightmost highest value in the slice, excluding zero.
func MaxNonZero(s collections.Slice) int {
	return compare(s, ifaceToFloat64(s.All()[0]), greaterNonZero)
}
