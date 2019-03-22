package zip

import (
	"archive/zip"
	"github.com/linehk/gopl/ch10/exercise10.2/archive"
	"os"
)

func init() {
	archive.InitFormats(
		archive.Format{"zip", "PK\x03\x04", 0, list})
	archive.InitFormats(
		archive.Format{"zip", "PK\x05\x06", 0, list})
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
			archive.FileHeader{f.Name, f.UncompressedSize64})
	}
	return headers, nil
}
