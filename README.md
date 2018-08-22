# DEFUNCT (I am not working on Go at the moment and this project will not receive any updates in the forseable future)

# collections
[![Build Status](https://travis-ci.org/kelceydamage/collections.svg?branch=master)](https://travis-ci.org/kelceydamage/collections) [![CircleCI](https://circleci.com/gh/kelceydamage/collections/tree/master.svg?style=shield)](https://circleci.com/gh/kelceydamage/collections/tree/master)  [![Coverage Status](https://coveralls.io/repos/github/kelceydamage/collections/badge.svg?branch=master&service=github)](https://coveralls.io/github/kelceydamage/collections?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/kelceydamage/collections)](https://goreportcard.com/report/github.com/kelceydamage/collections) [![CodeFactor](https://www.codefactor.io/repository/github/kelceydamage/collections/badge)](https://www.codefactor.io/repository/github/kelceydamage/collections) [![Maintainability](https://api.codeclimate.com/v1/badges/41fec5ebe52b5258ee3b/maintainability)](https://codeclimate.com/github/kelceydamage/collections/maintainability) [![GoDoc](https://godoc.org/github.com/kelceydamage/collections?status.svg)](https://godoc.org/github.com/kelceydamage/collections) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 

import "github.com/kelceydamage/collections"

Collections is a library of types and methods that make manipulating slices a lot easier by providing some basic functionality. 

Included in collections are a couple interafaces:

### Slice 

this is a base interface and all slice-types in collections implement this interface.

```Go
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
}
```

```Go
type Slice interface {
	BaseSlice
	Sort()
	Reverse()
}
```

```Go
type NumSlice interface {
	Slice
	Avg() float64
	AvgNonZero() float64
	StdDev() float64
	Variance() float64
}
``` 

### Types

* IntSlice
* IntSliceM
* IntQueue
* IntStack
* Float64Slice
* Float64SliceM
* Float64Queue
* Float64Stack
* StringSlice
* BoolSlice 

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
