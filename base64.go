package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"

	"github.com/xuender/goutils/u32"
)

func main() {
	for i := 1; i < 110; i++ {
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, uint32(i))
		end := 4
		for i := 3; i > 0; i-- {
			if b[i] == 0 {
				end = i
			} else {
				break
			}
		}
		// binary.BigEndian.PutUint32(b,uint32(i))
		fmt.Println(base64.StdEncoding.EncodeToString(b[0:end]))
	}
	for i := 0; i < 1000; i++ {
		fmt.Printf("%d -> %s -> %d\n", i, u32.Uint32ToStr(uint32(i)), u32.StrToUint32(u32.Uint32ToStr(uint32(i))))
	}
}
