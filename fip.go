package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/ip"
	"io"
	"os"
	"strings"
)

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

var ips map[string]string

func read(line string) {
	_, ok := ips[line]
	if !ok {
		ips[line] = ""
	}
}

func main() {
	ips = make(map[string]string)
	for _, a := range os.Args[1:] {
		ReadLine(a, read)
	}

	ip.Load("./17monipdb.dat")
	for k, _ := range ips {
		address := ip.Find(k)
		ips[k] = address
		fmt.Printf("%s = %s\n", k, address)
	}
}
