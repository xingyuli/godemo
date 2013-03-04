package main

import "fmt"
import "os"

type PathError struct {
	Op string
	Path string
	Err error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func main() {
	err := PathError{"open", "/etc/passex", fmt.Errorf("no such file or directory")}
	fmt.Fprintln(os.Stderr, err.Error())
}
