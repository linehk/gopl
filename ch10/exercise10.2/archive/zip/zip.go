package zip

import (
	"archive/zip"
	"os"

	"github.com/linehk/gopl/ch10/exercise10.2/archive"
)

func init() {
	archive.InitFormats(
		archive.Format{Name: "zip",
			Str: "PK\x03\x04", Offset: 0, List: list})
	archive.InitFormats(
		archive.Format{Name: "zip",
			Str: "PK\x05\x06", Offset: 0, List: list})
}

func list(f *os.File) ([]archive.FileHeader, error) {
	var headers []archive.FileHeader
	r, err := zip.OpenReader(f.Name())
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		headers = append(headers,
			archive.FileHeader{
				Name: f.Name, Size: f.UncompressedSize64})
	}
	return headers, nil
}
