package gcode

import (
	"fmt"
	"strings"
)

// ConvertToSVG returns the SVG code for the passed G-Code file.
func ConvertToSVG(f *File) string {
	paths := make(map[int][]string)

	var x, y float64
	var g int
	var path []string

	// range over all codes
	for _, l := range f.Lines {
		var ok bool

		for _, c := range l.Codes {
			if c.Letter == "G" && c.Value == 0 {
				// finish previous path
				paths[g] = append(paths[g], strings.Join(path, " "))

				// set state
				path = nil
				g = 0

				// starting pos
				path = append(path, fmt.Sprintf("M%f,%f", x, y))
			} else if c.Letter == "G" && c.Value == 1 {
				// finish previous path
				paths[g] = append(paths[g], strings.Join(path, " "))

				// set state
				path = nil
				g = 1

				// starting pos
				path = append(path, fmt.Sprintf("M%f,%f", x, y))

			} else if c.Letter == "X" {
				// set state
				x = c.Value
				ok = true
			} else if c.Letter == "Y" {
				// set state
				y = c.Value
				ok = true
			}
		}

		if ok {
			path = append(path, fmt.Sprintf("L%f,%f", x, y))
		}
	}

	// finish previous path
	paths[g] = append(paths[g], strings.Join(path, " "))

	var els []string

	// range over all levels
	for i, gpath := range paths {
		stroke := "black"
		if i == 0 {
			stroke = "red"
		}
		els = append(els, fmt.Sprintf(`<path d="%s" fill="none" stroke="%s" stroke-width="1" />`, strings.Join(gpath, " "), stroke))
	}

	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg">%s</svg>`, strings.Join(els, "\n"))
}
