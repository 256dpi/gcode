# gcode

[![Test](https://github.com/256dpi/gcode/actions/workflows/test.yml/badge.svg)](https://github.com/256dpi/gcode/actions/workflows/test.yml)
[![GoDoc](https://godoc.org/github.com/256dpi/gcode?status.svg)](http://godoc.org/github.com/256dpi/gcode)
[![Release](https://img.shields.io/github/release/256dpi/gcode.svg)](https://github.com/256dpi/gcode/releases)

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
  gcode svg <input> <output>

Options:
  -h --help  Show this screen.
```

## License

The MIT License (MIT)

Copyright (c) 2016 Joël Gähwiler
