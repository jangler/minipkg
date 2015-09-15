package ini

import (
	"bytes"
	"fmt"
	"io"
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
	windowsFile = "name1=value1\r\n" +
		"; comment followed by blank line\r\n" +
		"\r\n" +
		"name2=value2\r\n" +
		"[section1]\r\n" +
		"name1=value3\r\n" +
		"name2=value4\r\n" +
		"[section2]\r\n" +
		"name1=value5\r\n" +
		"name2=value6\r\n"
	invalidFile = `hello`
)

var testFile = File{
	"": Section{
		"name1": "value1",
		"name2": "value2",
	},
	"section1": Section{
		"name1": "value3",
		"name2": "value4",
	},
	"section2": Section{
		"name1": "value5",
		"name2": "value6",
	},
}

var sectionNames = []string{"", "section1", "section2"}

func TestReadEmpty(t *testing.T) {
	buf := bytes.NewBufferString(emptyFile)
	file, err := Read(buf)
	if err != nil {
		t.Errorf("Read() returned error: %v", err)
	} else if file == nil {
		t.Error("Read() returned nil file")
	}
}

func testReadValid(r io.Reader, t *testing.T) {
	file, err := Read(r)
	if err != nil {
		t.Fatalf("Read() returned error: %v", err)
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

func TestReadUnix(t *testing.T) {
	testReadValid(bytes.NewBufferString(unixFile), t)
	testReadValid(bytes.NewBufferString(windowsFile), t)
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

func TestWriteEmpty(t *testing.T) {
	file := make(File)
	buf := new(bytes.Buffer)
	if err := file.Write(buf); err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}
	if buf.String() != emptyFile {
		t.Error("Write() wrote %#v; expected %#v", buf.String(), emptyFile)
	}
}

func TestWriteValid(t *testing.T) {
	buf := new(bytes.Buffer)
	if err := testFile.Write(buf); err != nil {
		t.Fatalf("Write() returned error: %v", err)
	}
	testReadValid(buf, t)
}

type errorWriter struct{}

func (w errorWriter) Write(p []byte) (int, error) {
	return 0, fmt.Errorf("errorWriter.Write() called")
}

func TestWriteError(t *testing.T) {
	if err := testFile.Write(errorWriter{}); err == nil {
		t.Error("Write() did not return error")
	}
}
