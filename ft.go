package main

import (
  "fmt"
  "gopkg.in/h2non/filetype.v1"
  "io/ioutil"
)

func main() {
  buf, _ := ioutil.ReadFile("jddj.png")

  kind, unknown := filetype.Match(buf)
  if unknown != nil {
    fmt.Printf("Unknown: %s", unknown)
    return
  }

  fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
}
