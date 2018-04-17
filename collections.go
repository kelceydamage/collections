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

// Code
//---------------------------------------------------------------------------------------------------- <-100

// BaseSlice is a basic interface for implementing generic functions
type BaseSlice interface {
	All() []interface{}
	Overwrite(interface{})
	Append(interface{})
	Len() int
	Swap(int, int)
	TruncateLeft(int)
	TruncateRight(int)
	Mirror()
	Index(interface{}) int
	IndexRight(interface{}) int
	Get(int) interface{}
	Set(slice)
}

// Slice interface allows a small amount of generic application to all collections slices.
type Slice interface {
	BaseSlice
	Sort()
	Reverse()
}
