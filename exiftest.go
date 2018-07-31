package main

import (
	"fmt"
	"github.com/xiam/exif"
)

func main() {
	//data, err:= exif.Read("i.jpg")
	data, err := exif.Read("jddj.png")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for key, val := range data.Tags {
			fmt.Printf("%s = %s\n", key, val)
		}
	}
}
