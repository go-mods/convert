# convert

[![GoDoc](https://godoc.org/github.com/go-mods/convert?status.svg)](https://godoc.org/github.com/go-mods/convert)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-mods/convert)](https://goreportcard.com/report/github.com/go-mods/convert)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/go-mods/convert/blob/master/LICENSE)

Convert is a library for Go that provides easy and safe casting and conversion between different types. It's designed to be flexible, efficient, and easy to use in your Go projects.

## Features

- Convert between primitive types (string, int, float, bool, etc.)
- Convert to and from time.Time and time.Duration
- Convert to and from slices and maps
- Support for custom type conversions
- Comprehensive error handling
- Zero dependencies (except for testing)

## Installation

```shell
go get github.com/go-mods/convert
```

## Usage

```go
import "github.com/go-mods/convert"

// String conversions
str := convert.ToString(42)
num, err := convert.ToIntE("42")

// Time conversions
duration := convert.ToDuration("1h30m")
timeValue := convert.ToTime("2023-01-01T12:00:00Z")

// Slice conversions
intSlice := convert.ToIntSlice([]interface{}{1, "2", 3.0})

// Map conversions
stringMap := convert.ToStringMapString(map[interface{}]interface{}{"key": "value"})

// Custom conversions
customConverter := func(value interface{}) string {
    if v, ok := value.(CustomType); ok {
        s := v.CustomToString()
        return &s
    }
    return nil
}
result := convert.ToString(myCustomValue, customConverter)
```


## Documentation

For detailed documentation and examples, please refer to the [GoDoc](https://godoc.org/github.com/go-mods/convert).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
