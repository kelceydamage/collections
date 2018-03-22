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

// IntStack is a stack type specifically for ints.
type IntStack struct {
	stack
}

// New returna a new IntStack of [w] depth
func (q IntStack) New(w int) IntStack {
	newQ := IntStack{}
	newQ.Depth = w
	return newQ
}

// Get returns the element at the given index.
func (q *IntStack) Get(n int) int {
	return (*q).get(n).(int)
}

// Pop removes the newest element in the stack and returns it.
func (q *IntStack) Pop() (int, bool) {
	v, ok := q.pop()
	return v.(int), ok
}

// Put inserts an element onto the top of the stack, provided there is room.
func (q *IntStack) Put(n int) {
	q.put(n)
}

// Force inserts an element into the queue and poplefts an element if there is no space.
func (q *IntStack) Force(n int) {
	q.force(n)
}

// All returns the objects within the Stack as a slice
func (q *IntStack) All() IntSlice {
	return q.all().(IntSlice)
}
