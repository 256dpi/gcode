package gcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripComments(t *testing.T) {
	f := File{
		{
			Comment: "foo",
		},
		{
			Codes: []*GCode{
				{Comment: "bar"},
			},
		},
		{
			Codes: []*GCode{
				{Letter: "A"},
				{Comment: "baz"},
			},
		},
	}

	f = StripComments(f)

	assert.Equal(t, File{
		{
			Codes: []*GCode{
				{Letter: "A"},
			},
		},
	}, f)
}

func TestOffsetXYZ(t *testing.T) {
	f := File{
		{
			Codes: []*GCode{
				{Letter: "G", Value: 1},
			},
		},
		{
			Codes: []*GCode{
				{Letter: "X", Value: 2},
			},
		},
		{
			Codes: []*GCode{
				{Letter: "X", Value: 3},
				{Letter: "Y", Value: 4},
				{Letter: "Z", Value: 5},
			},
		},
	}

	OffsetXYZ(f, 1, 2, 3)

	assert.Equal(t, File{
		{
			Codes: []*GCode{
				{Letter: "G", Value: 1},
			},
		},
		{
			Codes: []*GCode{
				{Letter: "X", Value: 3},
			},
		},
		{
			Codes: []*GCode{
				{Letter: "X", Value: 4},
				{Letter: "Y", Value: 6},
				{Letter: "Z", Value: 8},
			},
		},
	}, f)
}
