// Package ini implements functions for reading and writing INI files.
package ini

import (
	"bufio"
	"errors"
	"io"
	"regexp"
)

var (
	commentRegexp  = regexp.MustCompile(`^;.*`)
	propertyRegexp = regexp.MustCompile(`^(.+)=(.*)`)
	sectionRegexp  = regexp.MustCompile(`^\[(.+)\]$`)
)

// A Section is a map of property names to values in a section of an INI file.
type Section map[string]string

// A File is a map of section names to Sections. The "global" section
// (properties declared before a section declaration) can be accessed by using
// an empty string as the map key.
type File map[string]Section

// Read reads r as an INI file. Each line of r must either be a property
// ("name=value"), section ("[section]"), comment ("; comment"), or blank line.
func Read(r io.Reader) (File, error) {
	scanner := bufio.NewScanner(r)
	file := make(File)
	section := ""
	file[section] = make(Section)

	for scanner.Scan() {
		line := scanner.Text()
		if m := propertyRegexp.FindStringSubmatch(line); m != nil {
			file[section][m[1]] = m[2]
		} else if line == "" || commentRegexp.MatchString(line) {
			// do nothing with blank lines and comment lines
		} else if m := sectionRegexp.FindStringSubmatch(line); m != nil {
			section = m[1]
			file[section] = make(Section)
		} else {
			return nil, errors.New("ini.Read(): invalid line: " + line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return file, nil
}

// Write writes f to w in INI format.
func Write(f File, w io.Writer) error {
	buf := bufio.NewWriter(w)

	// write global properties
	if section := f[""]; section != nil {
		for name, value := range section {
			_, err := buf.WriteString(name + "=" + value + "\r\n")
			if err != nil {
				return err
			}
		}

		if _, err := buf.WriteString("\r\n"); err != nil {
			return err
		}
	}

	// write sections
	for sectionName, section := range f {
		if sectionName == "" {
			break // since we already wrote global properties
		}

		if _, err := buf.WriteString("[" + sectionName + "]\r\n"); err != nil {
			return err
		}

		for name, value := range section {
			_, err := buf.WriteString(name + "=" + value + "\r\n")
			if err != nil {
				return err
			}
		}

		if _, err := buf.WriteString("\r\n"); err != nil {
			return err
		}
	}

	return buf.Flush()
}
