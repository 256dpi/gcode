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
			Codes: []*Code{
				{Comment: "bar"},
			},
		},
		{
			Codes: []*Code{
				{Letter: "A"},
				{Comment: "baz"},
			},
		},
	}

	f = StripComments(f)

	assert.Equal(t, File{
		{
			Codes: []*Code{
				{Letter: "A"},
			},
		},
	}, f)
}
