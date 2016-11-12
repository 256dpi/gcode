package gcode

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

// A GCode is either a single G-Code like "X12.3" or an in line comment in the
// form of "(Comment)".
type GCode struct {
	Letter  string
	Value   float64
	Comment string
}

// A Line consists of multiple G-Codes and a potential line comment.
type Line struct {
	Codes   []*GCode
	Comment string
}

// A File contains multiple lines of G-Codes.
type File struct {
	Lines []*Line
}

// ParseFile will parse a whole G-Code file from the passed reader.
func ParseFile(r io.Reader) (*File, error) {
	s := bufio.NewScanner(r)

	// prepare file
	file := &File{}

	// read line by line
	for s.Scan() {
		// parse lin
		line, err := ParseLine(s.Text())
		if err != nil {
			return file, err
		}

		// add line
		file.Lines = append(file.Lines, line)
	}

	// check error
	if err := s.Err(); err != nil {
		return file, err
	}

	return file, nil
}

// ParseLine will parse the specified string as a line of G-Codes.
func ParseLine(s string) (*Line, error) {
	// prepare line
	l := &Line{}

	// extract line comment
	if i := strings.Index(s, ";"); i >= 0 {
		// save comment
		l.Comment = s[i+1:]

		// reset string
		s = strings.TrimSpace(s[:i])
	}

	// check string
	if s == "" {
		return l, nil
	}

	// parse line
	for s != "" {
		// prepare code
		c := &GCode{}

		// check for word comment
		if strings.HasPrefix(s, "(") {
			if i := strings.Index(s, ")"); i >= 0 {
				// save comment
				c.Comment = s[1:i]

				// reset string
				s = strings.TrimSpace(s[i+1:])

				// add code
				l.Codes = append(l.Codes, c)

				// go on
				continue
			} else {
				return l, errors.New("missing ) for word comment")
			}
		}

		// check letter
		if !unicode.IsUpper(rune(s[0])) {
			return l, errors.New("expected uppercase letter to begin word")
		}

		// get word and reset string
		var w string
		if i := strings.Index(s, " "); i >= 0 {
			w = s[:i]
			s = strings.TrimSpace(s[i+1:])
		} else {
			w = s
			s = ""
		}

		// check length
		if len(w) < 2 {
			return l, errors.New("expected a word to have at least a length of two")
		}

		// extract letter
		c.Letter = string(w[0])
		w = w[1:]

		// parse value
		f, err := strconv.ParseFloat(w, 64)
		if err != nil {
			return l, err
		}

		c.Value = f

		// add code
		l.Codes = append(l.Codes, c)
	}

	return l, nil
}

// GenerateFile will write the specified G-Code file to the passed writer.
func GenerateFile(w io.Writer, f *File) error {
	// generate lines
	for _, l := range f.Lines {
		_, err := io.WriteString(w, GenerateLine(l))
		if err != nil {
			return err
		}
	}

	return nil
}

// GenerateLine will return the string representation of the passed line.
func GenerateLine(l *Line) string {
	// prepare string
	s := ""

	// write all codes
	for i, c := range l.Codes {
		// add space if any codes have been already added
		if i > 0 {
			s = s + " "
		}

		// add comment if present
		if c.Comment != "" {
			s = s + fmt.Sprintf("(%s)", c.Comment)
			continue
		}

		// write letter
		s = s + c.Letter

		// write value
		s = s + strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", c.Value), "0"), ".")
	}

	// write comment if existing
	if l.Comment != "" {
		// write space if any codes have been written
		if len(l.Codes) > 0 {
			s = s + " "
		}

		// add comment
		s = s + fmt.Sprintf(";%s", l.Comment)
	}

	// add line feed
	s = s + "\n"

	return s
}
