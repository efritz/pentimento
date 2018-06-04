# Pentimento

[![GoDoc](https://godoc.org/github.com/efritz/pentimento?status.svg)](https://godoc.org/github.com/efritz/pentimento)
[![Build Status](https://secure.travis-ci.org/efritz/pentimento.png)](http://travis-ci.org/efritz/pentimento)
[![Maintainability](https://api.codeclimate.com/v1/badges/7adb5ac9e13bd0583f49/maintainability)](https://codeclimate.com/github/efritz/pentimento/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/7adb5ac9e13bd0583f49/test_coverage)](https://codeclimate.com/github/efritz/pentimento/test_coverage)

Pentimento is a utility to help show progress events in the terminal.

## Example

Here is a complete working example. Each iteration of the loop overwrites the
previous output from the printer so only the current information stays on-screen.

This pattern can be adapted to show outstanding tasks, current external event progress,
or multiple progress bars (à la `docker pull`). The implementation uses ANSI codes, so
beware of backing writers that do not support them.

See the file `main.go` in the example directory of this project for a working example.

The output of this example can be viewed here:

[![asciicast](https://asciinema.org/a/7oF6KUK4Z8nvEoGMNvbrcPUxO.png)](https://asciinema.org/a/7oF6KUK4Z8nvEoGMNvbrcPUxO)

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
