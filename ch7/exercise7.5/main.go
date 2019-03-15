package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lr := LimitReader(strings.NewReader("12345"), 1)
	b, err := ioutil.ReadAll(lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}
	fmt.Printf("%s\n", b)
}

type LimitedReader struct {
	underlyingReader io.Reader
	remainBytes      int64
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.remainBytes <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.remainBytes {
		p = p[:r.remainBytes]
	}
	n, err = r.underlyingReader.Read(p)
	r.remainBytes -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
