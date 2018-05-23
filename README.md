# Pentimento

[![GoDoc](https://godoc.org/github.com/efritz/pentimento?status.svg)](https://godoc.org/github.com/efritz/pentimento)
[![Build Status](https://secure.travis-ci.org/efritz/pentimento.png)](http://travis-ci.org/efritz/pentimento)
[![Code Coverage](http://codecov.io/github/efritz/pentimento/coverage.svg?branch=master)](http://codecov.io/github/efritz/pentimento?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/efritz/pentimento)](https://goreportcard.com/report/github.com/efritz/pentimento)

Pentimento is a utility to help show progress events in the terminal.

## Example

Here is a complete working example. Each iteration of the loop overwrites the
previous output from the printer so only the current information stays on-screen.

This pattern can be adapted to show outstanding tasks, current external event progress,
or multiple progress bars (à la `docker pull`). The implementation uses ANSI codes, so
beware of backing writers that do not support them.

```go
package main

import (
	"fmt"
	"time"

	"github.com/efritz/pentimento"
)

func main() {
	fmt.Printf("Before\n")

	pentimento.PrintProgress(func(p *pentimento.Printer) {
		for i := 1; i <= 10; i++ {
			content := pentimento.NewContent()
			content.AddLine("[%s] a %d", pentimento.Throbber, i)
			content.AddLine("[%s] b %d", pentimento.Spinner, i)
			content.AddLine("[%s] c %d", pentimento.Throbber, i)
			p.WriteContent(content)

			<-time.After(time.Second / 2)
		}

		p.Reset()
		p.WriteString("Done")
	})

	fmt.Printf("After\n")
}
```

The output of this example can be viewed here:

[![asciicast](https://asciinema.org/a/nan3axRuk8QFSgABdFmGENoGE.png)](https://asciinema.org/a/nan3axRuk8QFSgABdFmGENoGE)

## License

Copyright (c) 2018 Eric Fritz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
