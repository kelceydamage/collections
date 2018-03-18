# collections
[![Build Status](https://travis-ci.org/kelceydamage/collections.svg?branch=master)](https://travis-ci.org/kelceydamage/collections) [![CircleCI](https://circleci.com/gh/kelceydamage/collections/tree/master.svg?style=shield)](https://circleci.com/gh/kelceydamage/collections/tree/master)  [![Coverage Status](https://coveralls.io/repos/github/kelceydamage/collections/badge.svg?branch=master&service=github)](https://coveralls.io/github/kelceydamage/collections?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/kelceydamage/collections)](https://goreportcard.com/report/github.com/kelceydamage/collections) [![CodeFactor](https://www.codefactor.io/repository/github/kelceydamage/collections/badge)](https://www.codefactor.io/repository/github/kelceydamage/collections) [![GoDoc](https://godoc.org/github.com/kelceydamage/collections?status.svg)](https://godoc.org/github.com/kelceydamage/collections) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 

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

However this is also the same as:

```
collections.IntSlice.Reverse()
```

# Documentation
https://godoc.org/github.com/kelceydamage/collections