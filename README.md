[![GoDoc](https://godoc.org/github.com/kelceydamage/collections?status.svg)](https://godoc.org/github.com/kelceydamage/collections)  [![Build Status](https://travis-ci.org/kelceydamage/collections.svg?branch=master)](https://travis-ci.org/kelceydamage/collections) 

# collections
import "github.com/kelceydamage/collections"

Collections is a library of types and methods that make manipulating slices a lot easier by providing some basic functionality.

Collections is compatible with the built in Sort types, and should be familiar to use. 

#### example: 
you can use a collections.IntSlice as a direct replacement for sort.IntSlice, while using all the features of Sort. This goes for all types included in Sort.

```go
sort.Sort(sort.Reverse(sort.IntSlice(s)))
```

Is the same as:
 
```
sort.Sort(sort.Reverse(collections.IntSlice(s)))
```

# Documentation
https://godoc.org/github.com/kelceydamage/collections