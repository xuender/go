package main

import (
	"log"
	"os"
)

func main() {
	log.Println("xxx")
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)
	log.SetPrefix("测试 ")
}
