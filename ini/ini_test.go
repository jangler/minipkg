package ini

import (
	"bytes"
	"fmt"
	"testing"
)

const (
	emptyFile = ``
	unixFile  = `name1=value1
name2=value2
; comment followed by blank line

[section1]
name1=value3
name2=value4
[section2]
name1=value5
name2=value6
`
	invalidFile = `hello`
)

var sectionNames = []string{"", "section1", "section2"}

func TestReadEmpty(t *testing.T) {
	buf := bytes.NewBufferString(emptyFile)
	file, err := Read(buf)
	if err != nil {
		t.Error(err)
	} else if file == nil {
		t.Error("Read() returned nil file")
	}
}

func TestReadUnix(t *testing.T) {
	buf := bytes.NewBufferString(unixFile)
	file, err := Read(buf)
	if err != nil {
		t.Fatal(err)
	} else if file == nil {
		t.Fatal("Read() returned nil file")
	}

	for i, sectionName := range sectionNames {
		if section := file[sectionName]; section != nil {
			for j := 0; j < 2; j++ {
				name := fmt.Sprintf("name%d", j+1)
				value := fmt.Sprintf("value%d", i*2+j+1)
				if section[name] != value {
					t.Errorf("[%s]: %s=%s; expected %s", sectionName, name,
						section[name], value)
				}
			}
		} else {
			t.Fatal("missing section: " + sectionName)
		}
	}
}

func TestReadInvalid(t *testing.T) {
	buf := bytes.NewBufferString(invalidFile)
	file, err := Read(buf)
	if err == nil {
		t.Error("Read() returned no error")
	}
	if file != nil {
		t.Error("Read() returned non-nil file")
	}
}

type errorReader struct{}

func (r errorReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("errorReader.Read() called")
}

func TestReadError(t *testing.T) {
	file, err := Read(errorReader{})
	if err == nil {
		t.Error("Read() returned no error")
	}
	if file != nil {
		t.Error("Read() returned non-nil file")
	}
}

func TestWrite(t *testing.T) {

}
