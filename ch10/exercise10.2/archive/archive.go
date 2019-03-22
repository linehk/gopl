package archive

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

type FileHeader struct {
	Name string
	Size uint64
}

func List(f *os.File) ([]FileHeader, error) {
	format, err := match(f)
	if err != nil {
		return nil, err
	}
	return format.List(f)
}

type Format struct {
	Name   string
	Str    string
	Offset int
	List   func(*os.File) ([]FileHeader, error)
}

var formats []Format

func InitFormats(format Format) {
	formats = append(formats, format)
}

func match(f *os.File) (*Format, error) {
	for _, format := range formats {
		f.Seek(0, io.SeekStart)
		r := bufio.NewReader(f)
		b, err := r.Peek(format.Offset + len(format.Str))
		if err == nil && bytes.Equal([]byte(format.Str), b[format.Offset:]) {
			f.Seek(0, io.SeekStart)
			return &format, nil
		}
	}
	return nil, errors.New("unknown format")
}
