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

// Code
//---------------------------------------------------------------------------------------------------- <-100

// IntQueue ...
type IntQueue struct {
	queue
}

// New returna a new IntQueue of [w] depth
func (q IntQueue) New(w int) IntQueue {
	newQ := IntQueue{}
	newQ.Depth = w
	return newQ
}

// Get returns the oldest element in the queue.
func (q *IntQueue) Get(n int) int {
	return (*q).get(n).(int)
}

// Pop removes the oldest element in the list and returns it.
func (q *IntQueue) Pop() (int, bool) {
	v, ok := q.pop()
	return v.(int), ok
}

// Put inserts an element into the queue provided there is room.
func (q *IntQueue) Put(n int) {
	q.put(n)
}

// Cycle inserts an element while popping and returning the oldest element at the same time.
func (q *IntQueue) Cycle(n int) (int, bool) {
	v, ok := q.cycle(n)
	return v.(int), ok
}

// Buffer will fill upt the queue returning nil. once the queue is full, it will start
// popping the oldest element and return it.
func (q *IntQueue) Buffer(n int) (int, bool) {
	v, ok := q.buffer(n)
	return v.(int), ok
}

func (q *IntQueue) All() IntSlice {
	return q.all().(IntSlice)
}
