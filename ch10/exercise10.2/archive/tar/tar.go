package tar

import (
	"archive/tar"
	"github.com/linehk/gopl/ch10/exercise10.2/archive"
	"io"
	"os"
)

func init() {
	archive.InitFormats(
		archive.Format{"tar", "ustar\x0000", 257, list})
}

func list(f *os.File) ([]archive.FileHeader, error) {
	var headers []archive.FileHeader
	tr := tar.NewReader(f)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		headers = append(headers,
			archive.FileHeader{hdr.Name, uint64(hdr.Size)})
	}
	return headers, nil
}
