# zw [![Build Status](https://travis-ci.org/trubitsyn/zw.svg?branch=master)](https://travis-ci.org/trubitsyn/zw) [![GoDoc](https://godoc.org/github.com/trubitsyn/zw?status.png)](https://godoc.org/github.com/trubitsyn/zw)
Zero-width character detection and removal for Go.

## Installation
`go get github.com/trubitsyn/zw`

## Usage
<pre>
package main

import (
	"github.com/trubitsyn/zw"
	"fmt"
)

func main() {
	s := "abc​def"				 // zero-width space between "c" and "d"
	fmt.Printf("% x\n", s)			 // 61 62 63 <b>e2 80 8b</b> 64 65 66
	clean := zw.RemoveZeroWidthCharacters(s) // zero-width space is in bold
	fmt.Printf("% x\n", clean)		 // 61 62 63 64 65 66
}
</pre>

## Testing
`go test github.com/trubitsyn/zw`

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
