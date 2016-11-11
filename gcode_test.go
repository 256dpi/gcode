package gcode

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testGCode = `; Line Comment
G1 ; After Line Comment
(Word Comment)
G2 (Word Comment) M1
G3 (Word Comment) M2 (Word Comment) M3
G4
G5 X0.000000 Y0.000000
G6 Z12.700000
G7 X-0.400000 Y0.800000
S3000
X56.666000
`

func TestParse(t *testing.T) {
	r := strings.NewReader(testGCode)

	file, err := ParseFile(r)
	assert.NoError(t, err)
	assert.Equal(t, File{
		Line{
			Comment: " Line Comment",
		},
		Line{
			Comment: " After Line Comment",
			Codes: []Code{
				{Letter: "G", Integer: 1},
			},
		},
		Line{
			Codes: []Code{
				{Comment: "Word Comment"},
			},
		},
		Line{
			Codes: []Code{
				{Letter: "G", Integer: 2},
				{Comment: "Word Comment"},
				{Letter: "M", Integer: 1},
			},
		},
		Line{
			Codes: []Code{
				{Letter: "G", Integer: 3},
				{Comment: "Word Comment"},
				{Letter: "M", Integer: 2},
				{Comment: "Word Comment"},
				{Letter: "M", Integer: 3},
			},
		},
		Line{
			Codes: []Code{
				{Letter: "G", Integer: 4},
			},
		},
		Line{
			Codes: []Code{
				{Letter: "G", Integer: 5},
				{Letter: "X"},
				{Letter: "Y"},
			},
		},
		Line{
			Codes: []Code{
				{Letter: "G", Integer: 6},
				{Letter: "Z", Float: 12.7},
			},
		},
		Line{
			Codes: []Code{
				{Letter: "G", Integer: 7},
				{Letter: "X", Float: -0.4},
				{Letter: "Y", Float: 0.8},
			},
		},
		Line{
			Codes: []Code{
				{Letter: "S", Integer: 3000},
			},
		},
		Line{
			Codes: []Code{
				{Letter: "X", Float: 56.666},
			},
		},
	}, file)
}

func TestGenerate(t *testing.T) {
	r := strings.NewReader(testGCode)

	file, err := ParseFile(r)
	assert.NoError(t, err)

	b := new(bytes.Buffer)

	err = GenerateFile(b, file)
	assert.NoError(t, err)
	assert.Equal(t, testGCode, b.String())
}
