package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 在后台把 conn 的内容发送到标准输出
	go mustCopy(os.Stdout, conn)

	sc := bufio.NewScanner(os.Stdin)
CLOSE:
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		cmd := args[0]
		switch cmd {
		case "close":
			fmt.Fprintln(conn, sc.Text())
			break CLOSE
		case "ls", "cd", "get":
			fmt.Fprintln(conn, sc.Text())
		case "send":
			if len(args) < 2 {
				log.Println("not enough argument")
			} else {
				filename := args[1]
				data, err := ioutil.ReadFile(filename)
				if err != nil {
					log.Printf("read file err: %v", err)
				}
				fmt.Fprintf(conn, "%s %d\n", sc.Text(), countLines(data))
				fmt.Fprintf(conn, "%s", data)
			}
		}
	}
}

func countLines(data []byte) int {
	c := 0
	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		c++
	}
	return c
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
