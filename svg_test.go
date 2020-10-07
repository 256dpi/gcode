package gcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToSVG(t *testing.T) {
	f := &File{
		Lines: []Line{
			{
				Codes: []GCode{
					{Letter: "G", Value: 0},
					{Letter: "X", Value: 2},
				},
			},
			{
				Codes: []GCode{
					{Letter: "X", Value: 3},
					{Letter: "Y", Value: 4},
				},
			},
			{
				Codes: []GCode{
					{Letter: "G", Value: 1},
					{Letter: "X", Value: -1},
				},
			},
			{
				Codes: []GCode{
					{Letter: "X", Value: 4},
					{Letter: "Y", Value: 5},
				},
			},
		},
	}

	svg := ConvertToSVG(f)

	assert.Equal(t, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-1.000000 0.000000 5.000000 5.000000"><path d=" M0.000000,0.000000 L2.000000,0.000000 L3.000000,4.000000" fill="none" stroke="red" stroke-width="1" />
<path d="M3.000000,4.000000 L-1.000000,4.000000 L4.000000,5.000000" fill="none" stroke="black" stroke-width="1" /></svg>`, svg)
}
