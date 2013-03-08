package main

import (
	"fmt"
	"io"
	"os"
)

func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		return "", err
	}

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf)
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}

func main() {
	contents, _ := Contents("/etc/hosts")
	fmt.Println(contents)
}
