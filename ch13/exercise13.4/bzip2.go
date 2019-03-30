package bzip

import (
	"io"
	"os/exec"
)

type writer struct {
	in  io.WriteCloser
	cmd *exec.Cmd
}

func NewWriter(w io.Writer) io.WriteCloser {
	// macOS /usr/bin/bzip2
	// Linux /bin/bzip2
	// Not for Windows
	cmd := exec.Command("bzip2")

	// 把原本发到 Stdout 的发到 w
	cmd.Stdout = w

	// 向 Stdin 写入
	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	// 开始
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	wc := &writer{in, cmd}
	return wc
}

func (w *writer) Write(data []byte) (int, error) {
	return w.in.Write(data)
}

func (w *writer) Close() error {
	err := w.in.Close()
	cmdErr := w.cmd.Wait()
	if err != nil {
		return err
	}
	if cmdErr != nil {
		return cmdErr
	}
	return nil
}
