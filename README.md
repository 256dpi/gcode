# gcode

[![Build Status](https://travis-ci.org/256dpi/gcode.svg?branch=master)](https://travis-ci.org/256dpi/gcode)
[![Coverage Status](https://coveralls.io/repos/github/256dpi/gcode/badge.svg?branch=master)](https://coveralls.io/github/256dpi/gcode?branch=master)
[![GoDoc](https://godoc.org/github.com/256dpi/gcode?status.svg)](http://godoc.org/github.com/256dpi/gcode)
[![Release](https://img.shields.io/github/release/256dpi/gcode.svg)](https://github.com/256dpi/gcode/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/256dpi/gcode)](http://goreportcard.com/report/256dpi/gcode)

**A G-Code parser and generator for Go.**
 
## Installation

Get the package using the go tool:

```bash
$ go get -u github.com/256dpi/gcode
```

## Usage

The included command line application can be used to post-process g-code files:

```
gcode.

Usage:
  gcode info <input>
  gcode strip <input> <output>
  gcode offset <input> <output> <x> <y> <z>

Options:
  -h --help  Show this screen.
```

## License

The MIT License (MIT)

Copyright (c) 2016 Joël Gähwiler
