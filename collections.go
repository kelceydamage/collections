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
// providing some basic functionality
package collections

// Slice interface allows a small amount of generic application to all collections slices.
type Slice interface {
	Len() int
	Swap(int int)
	Sort(bool)
	TruncateLft(int)
	TruncateRgt(int)
}

// Queue interface allows a small amount of generic application to all collections Queues.
type Queue interface {
}
