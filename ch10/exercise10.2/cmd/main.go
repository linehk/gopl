package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"
)

// 新建示例 tar 和 zip
func main() {
	var tbuf bytes.Buffer
	var zbuf bytes.Buffer
	tw := tar.NewWriter(&tbuf)
	zw := zip.NewWriter(&zbuf)

	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		// tar
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}

		// zip
		f, err := zw.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Close
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	// tar
	tf, err := os.Create("tar.tar")
	if err != nil {
		log.Fatal(err)
	}
	defer tf.Close()
	_, err = io.Copy(tf, &tbuf)
	if err != nil {
		log.Fatal(err)
	}

	// zip
	zf, err := os.Create("zip.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zf.Close()
	_, err = io.Copy(zf, &zbuf)
	if err != nil {
		log.Fatal(err)
	}
}
