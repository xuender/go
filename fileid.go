package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"

	"./hd"
)

func main() {
	app := cli.NewApp()
	app.Name = "fileid"
	app.Usage = "生成文件的唯一ID"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			cli.ShowAppHelp(c)
		} else {
			for _, a := range c.Args() {
				b := hd.NewFileId()
				err := hd.BufRead(a, func(bs []byte) { b.Write(bs) })
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Printf("%s\t%s\n", b.String(), a)
				}
			}
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
