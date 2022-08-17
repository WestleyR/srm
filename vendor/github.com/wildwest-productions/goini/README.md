# goini - A no-dependence builtin-like ini config file parser

## Install

```
go get -u github.com/wildwest-productions/goini
```

## Example

Short:

```go
type myConfig struct {
	Dir1 string `ini:"dir"`
	Runs int    `ini:"runs"`
}

iniBytes, err := ioutil.ReadFile("somefile.ini")

c := &myConfig{}
err := goini.Unmarshal(iniBytes, &c)
```

Full example:

```go
package main

import (
	"fmt"
	"os"

	"github.com/wildwest-productions/goini"
)

type testStruct struct {
	ConfigStr     string     `ini:"hello"`
	ConfigInt     int        `ini:"bar"`
	Val           string     `ini:"t"`
	Hello         subStruct  `ini:"sector"`
	SystemEnabled bool       `ini:"sys_enable"`
	Command       subStruct2 `ini:"command"`
}

type subStruct struct {
	Bar           string `ini:"hello"`
	SystemEnabled bool   `ini"sys_enable"`
}

type subStruct2 struct {
	Command string `ini:"command"`
	Runs    int    `ini:"runs"`
}

func main() {
	s := &testStruct{}

	iniData := []byte(
		`
hello=world
bar = 2
t=hi
sys_enable = true

[command]
command = echo hello
runs = 42

[sector]
hello=world-world
sys_enable = NO

`)

	err := goini.Unmarshal(iniData, &s)
	if err != nil {
		fmt.Printf("Failed to unmarshal ini: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Converted ini:\n%s\n", string(iniData))

	fmt.Printf("Into struct: %+v\n", s)
}
```

## Features

 - Supports comments, both `#` and `;`
 - Define with and without whitespace (`foo=bar` and `foo = bar`)
 - Supports common value types (bool, int, float32, float64, string)
 - No dependencies (other then builtins)

## Limitations

As of the first public commit:
 - sub-struct pointers will not work (willfix)
 - cannot unmarshal into a map (wontfix)

## License

This project is licensed under the terms of The Clear BSD License. See the
[LICENSE file](./LICENSE) for more details.

