package wave

import (
	"encoding/binary"
	"io"
)

const (
	WAVE_FORMAT_PCM        = 0x0001
	WAVE_FORMAT_IEEE_FLOAT = 0x0003
	WAVE_FORMAT_ALAW       = 0x0006
	WAVE_FORMAT_MULAW      = 0x0007
	WAVE_FORMAT_EXTENSIBLE = 0xfffe
)

type rawFile struct {
	RIFFID           [4]byte
	RIFFSize         uint32
	WaveID           [4]byte
	FmtID            [4]byte
	FmtSize          uint32
	FormatTag        uint16
	Channels         uint16
	SamplesPerSecond uint32
	BytesPerSecond   uint32
	BlockAlign       uint16
	BitsPerSample    uint16
	DataID           [4]byte
	DataSize         uint32
}

type File struct {
	Channels       int
	SampleRate     int
	BytesPerSample int
	Data           []byte
}

func (f *File) Write(w io.Writer) error {
	raw := &rawFile{
		RIFFID:           [4]byte{'R', 'I', 'F', 'F'},
		RIFFSize:         4 + 24 + 8 + uint32(len(f.Data)),
		WaveID:           [4]byte{'W', 'A', 'V', 'E'},
		FmtID:            [4]byte{'f', 'm', 't', ' '},
		FmtSize:          16,
		FormatTag:        WAVE_FORMAT_PCM,
		Channels:         uint16(f.Channels),
		SamplesPerSecond: uint32(f.SampleRate),
		BytesPerSecond:   uint32(f.SampleRate * f.BytesPerSample * f.Channels),
		BlockAlign:       uint16(f.BytesPerSample * f.Channels),
		BitsPerSample:    8 * uint16(f.BytesPerSample),
		DataID:           [4]byte{'d', 'a', 't', 'a'},
		DataSize:         uint32(len(f.Data)),
	}
	padByte := raw.RIFFSize%2 == 1
	if padByte {
		raw.RIFFSize += 1
		raw.DataSize += 1
	}

	if err := binary.Write(w, binary.LittleEndian, raw); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, f.Data); err != nil {
		return err
	}
	if padByte {
		if _, err := w.Write([]byte{0}); err != nil {
			return err
		}
	}

	return nil
}
