# GoCSSUnits
A lightweight Go library for parsing CSS font size units and dimensions with JSON marshaler/unmarshaler.

## Installation
```
go get github.com/EgorSheff/gocssunits
```

## Usage
```go
package main

import (
    "fmt"
    "log"

    "github.com/EgorSheff/gocssunits"
)

func main() {
    // Parse font size
    size, err := gocssunits.ParseFontSize("1.5rem")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Value: %.1f%s\n", size.Value, size.Unit) // Value: 1.5rem
}
```
