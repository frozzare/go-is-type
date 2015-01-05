# go-is-type [![Build Status](https://travis-ci.org/frozzare/go-is-type.svg?branch=master)](https://travis-ci.org/frozzare/go-is-type)

Simple type checking.

View the [docs](http://godoc.org/github.com/frozzare/go-is-type).

## Installation

```
$ go get github.com/frozzare/go-is-type
```

## Example

```go
package main

import "github.com/frozzare/go-is-type"

func main() {
	is.String("hello")
	//=> true
	
	is.Bool(true)
	//=> true
	
	is.Array("string")
	//=> false 
}
```

See [is_test.go](is_test.go) for more examples.

# License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
