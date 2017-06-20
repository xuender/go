package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/xuender/goutils"
)

func main() {
	file, err := os.Open("/home/ender/work/groom/ml-latest-small/movies.csv")
	goutils.CheckError(err)
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else {
			goutils.CheckError(err)
		}
		log.Println(record)
	}
}
