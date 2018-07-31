package hd

import (
	"bufio"
	"io"
	"os"
)

func BufRead(path string, hookfn func([]byte)) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 102400)
	bfRd := bufio.NewReader(f)
	for {
		n, err := bfRd.Read(buf)
		hookfn(buf[:n])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}
