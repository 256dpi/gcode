package gcode

func StripComments(f File) File {
	cl := 0
	for i := range f {
		j := i - cl
		l := f[j]

		cd := 0
		for ii := range l.Codes {
			jj := ii - cd
			c := l.Codes[jj]

			// remove codes with comments
			if c.Comment != "" {
				l.Codes = append(l.Codes[:jj], l.Codes[jj+1:]...)
				cd++
				continue
			}
		}

		// remove lines with comments or no codes
		if l.Comment != "" || len(l.Codes) == 0 {
			f = append(f[:j], f[j+1:]...)
			cl++
			continue
		}
	}

	return f
}
