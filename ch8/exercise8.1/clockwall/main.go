package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type server struct {
	name    string
	addr    string
	message string
}

func main() {
	servers := parse(os.Args[1:])
	for _, ser := range servers {
		conn, err := net.Dial("tcp", ser.addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go func(ser *server) {
			sc := bufio.NewScanner(conn)
			for sc.Scan() {
				ser.message = sc.Text()
			}
		}(ser)
	}

	for {
		fmt.Printf("\n")
		for _, server := range servers {
			fmt.Printf("%s: %s\n", server.name, server.message)
		}
		fmt.Print("--------")

		time.Sleep(time.Second)
	}
}

func parse(args []string) []*server {
	var servers []*server
	for _, arg := range args {
		s := strings.SplitN(arg, "=", 2)
		servers = append(servers, &server{s[0], s[1], ""})
	}
	return servers
}
