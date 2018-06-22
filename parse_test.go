package gcode

import (
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
G5 X0 Y0
G6 Z12.7
G7 X-0.4 Y0.8
S3000
X56.666
`

func TestParse(t *testing.T) {
	r := strings.NewReader(testGCode)

	file, err := ParseFile(r)
	assert.NoError(t, err)
	assert.Equal(t, &File{
		Lines: []Line{
			{
				Comment: " Line Comment",
			},
			{
				Comment: " After Line Comment",
				Codes: []GCode{
					{Letter: "G", Value: 1},
				},
			},
			{
				Codes: []GCode{
					{Comment: "Word Comment"},
				},
			},
			{
				Codes: []GCode{
					{Letter: "G", Value: 2},
					{Comment: "Word Comment"},
					{Letter: "M", Value: 1},
				},
			},
			{
				Codes: []GCode{
					{Letter: "G", Value: 3},
					{Comment: "Word Comment"},
					{Letter: "M", Value: 2},
					{Comment: "Word Comment"},
					{Letter: "M", Value: 3},
				},
			},
			{
				Codes: []GCode{
					{Letter: "G", Value: 4},
				},
			},
			{
				Codes: []GCode{
					{Letter: "G", Value: 5},
					{Letter: "X"},
					{Letter: "Y"},
				},
			},
			{
				Codes: []GCode{
					{Letter: "G", Value: 6},
					{Letter: "Z", Value: 12.7},
				},
			},
			{
				Codes: []GCode{
					{Letter: "G", Value: 7},
					{Letter: "X", Value: -0.4},
					{Letter: "Y", Value: 0.8},
				},
			},
			{
				Codes: []GCode{
					{Letter: "S", Value: 3000},
				},
			},
			{
				Codes: []GCode{
					{Letter: "X", Value: 56.666},
				},
			},
		},
	}, file)
}

func TestParseInvalid(t *testing.T) {
	gCodes := []string{
		"(Invalid Comment", // <- missing end brace
		"g1",               // <- not upper case
		"G",                // <- missing value
		"GF",               // <- invalid value
	}

	for _, gc := range gCodes {
		_, err := ParseFile(strings.NewReader(gc))
		assert.Error(t, err)
	}
}
