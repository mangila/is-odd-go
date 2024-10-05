[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/mangila/is-odd-go/issues)
[![GoDoc](https://godoc.org/github.com/Pallinder/go-randomdata?status.svg)](https://godoc.org/github.com/mangila/is-odd-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Pallinder/go-randomdata)](https://goreportcard.com/report/github.com/mangila/is-odd-go)
# is-odd-go
When you are in trouble and can't decide if the number is odd or even. This go utility will solve it for you. 

* If the number ends in 2, 4, 6, 8, or 0, then the number is even.
* If the number ends in 1, 3, 5, 7, or 9, then the number is odd.

## Installation

```go get github.com/github.com/mangila/is-odd-go```

## Usage

```go

package main

import (
    "fmt"
    "github.com/mangila/is-odd-go"
)

func main() {
	fmt.Println(isodd.IsOdd(5)) // Output: true
	fmt.Println(isodd.IsOdd(4)) // Output: false
}

```
