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
// This file is to pretty up the godoc.

// Package collections is a library of types and methods that make manipulating slices a lot easier by
// providing some basic functionality.
//
// Collections is compatible with the built in Sort types, and should be familiar to use. As an example:
// you can use a collections.IntSlice as a direct replacement for sort.IntSlice, while using all the
// features of Sort. This goes for all types included in Sort.
//
//   sort.Sort(sort.Reverse(sort.IntSlice(s)))
// Is the same as:
//   sort.Sort(sort.Reverse(collections.IntSlice(s)))
// However this is also the same as:
//   collections.IntSlice.Reverse()
package collections

// Slice interface allows a small amount of generic application to all collections slices.
type Slice interface {
	Len() int
	Swap(int, int)
	Less(int, int) bool
	Sort()
	Reverse()
	Mirror()
	Index(interface{}) int
	IndexRight(interface{}) int
	TruncateLeft(int)
	TruncateRight(int)
	Append(interface{}) int
}

// NumSlice interface extends Slice with basic math functions.
type NumSlice interface {
	Slice
	Sum() interface{}
	Avg() float64
	Max() (interface{}, int)
	MaxNonZero() (interface{}, int)
	Min() (interface{}, int)
	MinNonZero() (interface{}, int)
	StdDev() float64
	Variance() float64
}

// Swap is a function interface for the Swap method on collections slices.
func Swap(s Slice, i, j int) {
	s.Swap(i, j)
}

// Less is a function interface for the Less method on collections slices.
func Less(s Slice, i, j int) bool {
	return s.Less(i, j)
}

// Sort is a function interface for the Sort method on collections slices.
func Sort(s Slice) {
	s.Sort()
}

// Reverse is a function interface for the Reverse method on collections slices.
func Reverse(s Slice) {
	s.Reverse()
}

// Mirror is a function interface for the Mirror method on collections slices.
func Mirror(s Slice) {
	s.Mirror()
}

// TruncateLeft is a function interface for the TruncateLeft method on collections slices.
func TruncateLeft(s Slice, n int) {
	s.TruncateLeft(n)
}

// TruncateRight is a function interface for the TruncateRight method on collections slices.
func TruncateRight(s Slice, n int) {
	s.TruncateRight(n)
}

// Len is a function interface for the Len method on collections slices.
//
// This is only included for consistency and one should use the default len() function.
func Len(s Slice) int {
	return s.Len()
}

// Functions receiving interfaces
//---------------------------------------------------------------------------------------------------- <-100

// Index is a function interface for the Index method on collections slices.
func Index(s Slice, j interface{}) int {
	return s.Index(j)
}

// IndexRight is a function interface for the IndexRight method on collections slices.
func IndexRight(s Slice, j interface{}) int {
	return s.IndexRight(j)
}

// Append is a function interface for the append method on collections slices.
func Append(s Slice, j interface{}) int {
	return s.Append(j)
}
