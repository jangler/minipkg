package wave

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"testing"
)

func TestFileWrite(t *testing.T) {
	buf := new(bytes.Buffer)

	// test writing a file
	file := File{
		Channels:       1,
		SampleRate:     8363,
		BytesPerSample: 1,
		Data:           []byte{0, 0, 0, 0, 255, 255, 255, 255},
	}
	if err := file.Write(buf); err != nil {
		t.Fatalf("File.Write() returned error: %v", err)
	}

	// test reading back the data into a raw struct
	raw := new(rawFile)
	if err := binary.Read(buf, binary.LittleEndian, raw); err != nil {
		t.Fatalf("binary.Read() returned error: %v", err)
	}
	if got, want := string(raw.RIFFID[:]), "RIFF"; got != want {
		t.Errorf("rawFile.RIFFID == %#v; want %#v", got, want)
	}
	if got, want := raw.RIFFSize, uint32(44); got != want {
		t.Errorf("rawFile.RIFFSize == %v; want %v", got, want)
	}
	if got, want := string(raw.WaveID[:]), "WAVE"; got != want {
		t.Errorf("rawFile.WaveID == %#v; want %#v", got, want)
	}
	if got, want := string(raw.FmtID[:]), "fmt "; got != want {
		t.Errorf("rawFile.FmtID == %#v; want %#v", got, want)
	}
	if got, want := raw.FmtSize, uint32(16); got != want {
		t.Errorf("rawFile.FmtSize == %v; want %v", got, want)
	}
	if got, want := raw.FormatTag, uint16(WAVE_FORMAT_PCM); got != want {
		t.Errorf("rawFile.FormatTag == %v; want %v", got, want)
	}
	if got, want := raw.Channels, uint16(1); got != want {
		t.Errorf("rawFile.Channels == %v; want %v", got, want)
	}
	if got, want := raw.SamplesPerSecond, uint32(8363); got != want {
		t.Errorf("rawFile.SamplesPerSecond == %v; want %v", got, want)
	}
	if got, want := raw.BytesPerSecond, uint32(8363); got != want {
		t.Errorf("rawFile.BytesPerSecond == %v; want %v", got, want)
	}
	if got, want := raw.BlockAlign, uint16(1); got != want {
		t.Errorf("rawFile.BlockAlign == %v; want %v", got, want)
	}
	if got, want := raw.BitsPerSample, uint16(8); got != want {
		t.Errorf("rawFile.BitsPerSample == %v; want %v", got, want)
	}
	if got, want := string(raw.DataID[:]), "data"; got != want {
		t.Errorf("rawFile.DataID == %#v; want %#v", got, want)
	}
	if got, want := raw.DataSize, uint32(8); got != want {
		t.Errorf("rawFile.DataSize == %v; want %v", got, want)
	}
	got, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Fatalf("ioutil.ReadAll() returned error: %v", err)
	}
	if want := file.Data; bytes.Compare(got, want) != 0 {
		t.Errorf("data == %v; want %v", got, want)
	}
}
