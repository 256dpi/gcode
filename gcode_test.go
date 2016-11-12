package gcode

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {
	r := strings.NewReader(testGCode)

	file, err := ParseFile(r)
	assert.NoError(t, err)

	b := new(bytes.Buffer)

	err = WriteFile(b, file)
	assert.NoError(t, err)
	assert.Equal(t, testGCode, b.String())
}
