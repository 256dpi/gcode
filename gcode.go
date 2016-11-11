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

type Code struct {
	Letter  string
	Integer int64
	Float   float64
	Comment string
}

type Line struct {
	Codes   []Code
	Comment string
}

type File []Line

func ParseFile(r io.Reader) (File, error) {
	s := bufio.NewScanner(r)

	// prepare file
	var file File

	// read line by line
	for s.Scan() {
		// parse lin
		line, err := ParseLine(s.Text())
		if err != nil {
			return file, err
		}

		// add line
		file = append(file, line)
	}

	// check error
	if err := s.Err(); err != nil {
		return file, err
	}

	return file, nil
}

func ParseLine(s string) (Line, error) {
	// prepare line
	l := Line{}

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
		c := Code{}

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

		// parse integer
		n, err := strconv.ParseInt(w, 10, 64)
		if err == nil {
			c.Integer = n
		} else {
			f, err := strconv.ParseFloat(w, 64)
			if err != nil {
				return l, err
			}

			c.Float = f
		}

		// add code
		l.Codes = append(l.Codes, c)
	}

	return l, nil
}

func GenerateFile(w io.Writer, f File) error {
	// generate lines
	for _, l := range f {
		err := GenerateLine(w, l)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateLine(w io.Writer, l Line) error {
	// write all codes
	for i, c := range l.Codes {
		// write space if any codes have been before
		if i > 0 {
			_, err := io.WriteString(w, " ")
			if err != nil {
				return err
			}
		}

		// check comment
		if c.Comment != "" {
			_, err := fmt.Fprintf(w, "(%s)", c.Comment)
			if err != nil {
				return err
			}

			continue
		}

		// write letter
		_, err := io.WriteString(w, c.Letter)
		if err != nil {
			return err
		}

		// write integer if set
		if c.Integer != 0 {
			_, err := fmt.Fprintf(w, "%d", c.Integer)
			if err != nil {
				return err
			}

			continue
		}

		// otherwise write float
		_, err = fmt.Fprintf(w, "%f", c.Float)
		if err != nil {
			return err
		}
	}

	// write comment if existing
	if l.Comment != "" {
		// write space if any codes have been written
		if len(l.Codes) > 0 {
			_, err := io.WriteString(w, " ")
			if err != nil {
				return err
			}
		}

		_, err := fmt.Fprintf(w, ";%s", l.Comment)
		if err != nil {
			return err
		}
	}

	// write line feed
	_, err := io.WriteString(w, "\n")
	if err != nil {
		return err
	}

	return nil
}
