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

// Code
//-------------------------------------------------------------------------------------------------- <-100

type slice []interface{}

func (e *slice) Append(x interface{}) {
	(*e) = append((*e), x)
}

func (e *slice) All() []interface{} {
	return (*e)
}

func (e *slice) Overwrite(x interface{}) {
	(*e) = x.(slice)
}

// Len implementation for general sort.
func (e slice) Len() int {
	return len(e)
}

// Swap implementation for general sort.
func (e slice) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// TruncateLeft shrinks the slice to [n] amount of float64s starting from the left.
func (e *slice) TruncateLeft(n int) {
	if n != -1 && n <= len((*e)) {
		(*e) = (*e)[:n]
	}
}

// TruncateRight shrinks the slice to [n] amount of float64s starting from the right.
func (e *slice) TruncateRight(n int) {
	if n != -1 && n <= len((*e)) {
		(*e) = (*e)[(*e).Len()-n:]
	}
}

// Mirror will swap the order of ints in the slice.
func (e *slice) Mirror() {
	for i, j := 0, len((*e))-1; i < j; i, j = i+1, j-1 {
		(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
	}
}

// Index retrives the first index [i] from the left, for the provided value [j] (int).
// Returns -1 if index not found.
func (e *slice) Index(j interface{}) int {
	for i, v := range *e {
		if v == j {
			return i
		}
	}
	return -1
}

// IndexRight retrives the first index [i] from the right, for the provided value [j] (int).
// Returns -1 if index not found.
func (e *slice) IndexRight(j interface{}) int {
	k := -1
	for i, v := range *e {
		if v == j {
			k = i
		}
	}
	if k == -1 {
		return k
	}
	return k
}
