# zero-width [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go) [![Build Status](https://travis-ci.org/trubitsyn/go-zero-width.svg?branch=master)](https://travis-ci.org/trubitsyn/go-zero-width) [![GoDoc](https://godoc.org/github.com/trubitsyn/go-zero-width?status.svg)](https://godoc.org/github.com/trubitsyn/go-zero-width) [![Go Report Card](https://goreportcard.com/badge/github.com/trubitsyn/go-zero-width)](https://goreportcard.com/report/github.com/trubitsyn/go-zero-width)
Zero-width character detection and removal for Go. Inspired by [this Medium article](https://medium.com/@umpox/be-careful-what-you-copy-invisibly-inserting-usernames-into-text-with-zero-width-characters-18b4e6f17b66).

## Installation
`go get github.com/trubitsyn/go-zero-width`

## Usage
<pre>
package main

import (
	"github.com/trubitsyn/go-zero-width"
	"fmt"
)

func main() {
	login := "abc​def"					// zero-width space between "c" and "d"
	clean := zerowidth.RemoveZeroWidthCharacters(login)	// a  b  c           d  e  f
	fmt.Printf("% x\n", login)				// 61 62 63 <b>e2 80 8b</b> 64 65 66
	fmt.Printf("% x\n", clean)				// 61 62 63          64 65 66
}
</pre>

## Testing
```
go get -t github.com/trubitsyn/go-zero-width
go test github.com/trubitsyn/go-zero-width
```

## LICENSE
```
Copyright 2018 Nikola Trubitsyn

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
