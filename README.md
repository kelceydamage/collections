# collections
[![Build Status](https://travis-ci.org/kelceydamage/collections.svg?branch=master)](https://travis-ci.org/kelceydamage/collections) [![CircleCI](https://circleci.com/gh/kelceydamage/collections/tree/master.svg?style=shield)](https://circleci.com/gh/kelceydamage/collections/tree/master)  [![Coverage Status](https://coveralls.io/repos/github/kelceydamage/collections/badge.svg?branch=master&service=github)](https://coveralls.io/github/kelceydamage/collections?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/kelceydamage/collections)](https://goreportcard.com/report/github.com/kelceydamage/collections) [![CodeFactor](https://www.codefactor.io/repository/github/kelceydamage/collections/badge)](https://www.codefactor.io/repository/github/kelceydamage/collections) [![Maintainability](https://api.codeclimate.com/v1/badges/41fec5ebe52b5258ee3b/maintainability)](https://codeclimate.com/github/kelceydamage/collections/maintainability) [![GoDoc](https://godoc.org/github.com/kelceydamage/collections?status.svg)](https://godoc.org/github.com/kelceydamage/collections) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 

import "github.com/kelceydamage/collections"

Collections is a library of types and methods that make manipulating slices a lot easier by providing some basic functionality. 

Included in collections are a couple interafaces:

### Slice 

this is a base interface and all slice-types in collections implement this interface.

```Go
type Slice interface {
	Len() int                           //value receiver
	Swap(int, int)                      //value receiver
	Less(int, int) bool                 //value receiver
	Sort()                              //pointer receiver
	Reverse()                           //pointer receiver
	Mirror()                            //pointer receiver
	Index(interface{}) int              //pointer receiver
	IndexRight(interface{}) int         //pointer receiver
	TruncateLeft(int)                   //pointer receiver
	TruncateRight(int)                  //pointer receiver
	Append(interface{}) int             //pointer receiver
}
```

Each of these methods also has a function API:

#### Len
Len is purely included for consistency. 

99% of the time you will want to use the builin len() function. However there are times where the Len() API can make dealing with the pointer receivers a lot easier.
```Go
Len(s Slice) int {}
```

#### Swap
Swap switches the position of items at the provided indexes [i] and [j].
```Go
Swap(s Slice, i, j int) {}
```

#### Less
Less returns true if the item at index [i] is less then the item at index [j].
```Go
Less(s Slice, i, j bool) {}
```

#### Sort
Sort wraps the sort.Sort() function using the type specific Len(), Less(), Swap() methods.
```Go
Sort(s Slice) {}
```

#### Reverse
Reverse wraps the sort.Sort(sort.Reverse()) function using either the type specific inverted Less() method, or the Mirror() method.
```Go
Reverse(s Slice) {}
```

#### Mirror
Mirror swaps the positions of all items in the slice providing essentially a mirror image.
```Go
Mirror(s Slice) {}
```

#### TruncateLeft
TruncateLeft resizes the slice to [n] items starting from the left-side. 
```Go
TruncateLeft(s Slice, n int) {}
```

#### TruncateRight
TruncateLeft resizes the slice to [n] items starting from the right-side. 
```Go
TruncateRight(s Slice, n int) {}
```

#### Index
Index returns the first index of a value [j] starting from the left. If the value is not found, or the type of [j] does not match the slice-type of [s], a -1 int is returned. 
```Go
Index(s Slice, j interface{}) int {}
```

#### IndexRight
Index returns the first index of a value [j] starting from the right. If the value is not found, or the type of [j] does not match the slice-type of [s], a -1 int is returned. 
```Go
IndexRight(s Slice, j interface{}) int {}
```

#### Append
Append adds item [j] to slice [s] inplace, provided [j] is of the type required by [s]. Otherwise a -1 will be returned

99% of the time you will want to use the builin append() function. However there are times where the Append() API can make dealing with the pointer receivers a lot easier.
```Go
Append(s Slice, j interface{}) int {}
```

All of the above function APIs have matching struct.method() syntax if prefered.

### NumSlice

this is a utility interface for IntSlice and FloatSlice, that extends Slice with basic mathematical functions.

```Go
type NumSlice interface {
	Slice
	Sum() float64
	Avg() float64
	Max() interface{}
	MaxNonZero() interface{}
	Min() interface{}
	MinNonZero() interface{}
	StdDev() float64
	Varaince() float64
}
```

These methods are for making number streams and metrics easier to manipulate in a clean fashion.

### Compatibility

Collections is compatible with the built in Sort types, and should be familiar to use. 

#### example: 
you can use a collections.IntSlice as a direct replacement for sort.IntSlice, while using all the features of Sort. This goes for all types included in Sort.

```Go
sort.Sort(sort.Reverse(sort.IntSlice(s))) 
```

Is the same as:
 
```Go
sort.Sort(sort.Reverse(collections.IntSlice(s)))
```

However this is also the same as:

```Go
collections.IntSlice.Reverse()
```

And

```Go
collections.Reverse(IntSlice)
```

# Documentation
https://godoc.org/github.com/kelceydamage/collections