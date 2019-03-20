package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	cwd := "."
CLOSE:
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		cmd := args[0]
		switch cmd {
		case "close":
			break CLOSE
		case "ls":
			if len(args) < 2 {
				ls(conn, cwd)
			} else {
				path := args[1]
				if err := ls(conn, path); err != nil {
					fmt.Fprint(conn, err)
				}
			}
		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(conn, "not enough argument")
			} else {
				cwd += "/" + args[1]
			}
		case "get":
			if len(args) < 2 {
				fmt.Fprintln(conn, "not enough argument")
			} else {
				filename := args[1]
				data, err := ioutil.ReadFile(filename)
				if err != nil {
					fmt.Fprint(conn, err)
				}
				fmt.Fprintf(conn, "%s\n", data)
			}
		case "send":
			filename := args[1]
			f, err := os.Create(filename)
			if err != nil {
				fmt.Fprint(conn, err)
			}
			defer f.Close()

			c, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Fprint(conn, err)
			}

			var texts string
			for i := 0; i < c && sc.Scan(); i++ {
				texts += sc.Text() + "\n"
			}
			texts = strings.TrimSuffix(texts, "\n")

			fmt.Fprint(f, texts)
		}
	}
}

func ls(w io.Writer, dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Fprintf(w, "%s\n", file.Name())
	}
	return nil
}
