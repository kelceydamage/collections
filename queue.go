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

// queue is a FIFO type of queue.
type queue struct {
	Content Slice
	Depth   int
}

// Get returns the element at the given index.
func (q *queue) get(n int) interface{} {
	return (*q).Content.All()[n]
}

// Pop removes the oldest element in the list and returns it.
func (q *queue) pop() (interface{}, bool) {
	if q.Content.Len() == 0 {
		return -1, false
	}
	v := (*q).Content.All()[0]
	(*q).Content.Overwrite((*q).Content.All()[1:])
	return v, true
}

// Put inserts an element into the queue provided there is room.
func (q *queue) put(x interface{}) {
	l := (*q).Content.Len()
	if l <= (*q).Depth {
		(*q).Content.Append(x)
	}
}

// Cycle inserts an element while popping and returning the oldest element at the same time.
func (q *queue) cycle(x interface{}) (interface{}, bool) {
	(*q).put(x)
	return (*q).pop()
}

// Buffer will fill upt the queue returning nil. once the queue is full, it will start
// popping the oldest element and return it.
func (q *queue) buffer(x interface{}) (interface{}, bool) {
	l := (*q).Content.Len()
	w := (*q).Depth
	(*q).put(x)
	if l == w+1 {
		return (*q).pop()
	}
	return -1, false
}

// All returns the objects within the Stack as a slice
func (q *queue) all() interface{} {
	return q.Content.All()
}
