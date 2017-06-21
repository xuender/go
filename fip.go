package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/axgle/ip"
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
}

var ips map[string]string

func read(line string) {
	_, ok := ips[line]
	if !ok {
		ips[line] = ""
	}
}

func main() {
	ip.Load("./17monipdb.dat")
	fmt.Println(ip.Find("8.8.8.8"))
	fmt.Println(ip.Find("202.106.46.151"))
	fmt.Println(ip.Find("202.115.128.64"))

}

func old() {
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
