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

type stack struct {
	Content Slice
	Depth   int
}

// Get returns the element at the given index.
func (q *stack) get(n int) interface{} {
	return (*q).Content.All()[n]
}

// Pop removes the newest element in the stack and returns it.
func (q *stack) pop() (interface{}, bool) {
	if q.Content.Len() == 0 {
		return -1, false
	}
	l := (*q).Content.Len()
	v := (*q).Content.All()[l]
	(*q).Content.Overwrite((*q).Content.All()[:l-1])
	return v, true
}

func (q *stack) popLeft() (interface{}, bool) {
	if q.Content.Len() == 0 {
		return -1, false
	}
	v := (*q).Content.All()[0]
	(*q).Content.Overwrite((*q).Content.All()[1:])
	return v, true
}

// Put inserts an element onto the top of the stack, provided there is room.
func (q *stack) put(x interface{}) {
	l := (*q).Content.Len()
	if l <= (*q).Depth {
		(*q).Content.Append(x)
	}
}

func (q *stack) force(x interface{}) {
	l := (*q).Content.Len()
	if l == (*q).Depth {
		q.popLeft()
	}
	(*q).Content.Append(x)
}

func (q *stack) all() interface{} {
	return q.Content.All()
}
