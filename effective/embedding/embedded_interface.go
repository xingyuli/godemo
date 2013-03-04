package main

import "fmt"

/* interfaces */
type Opener interface {
	Open() (data string, err error)
}

type Closer interface {
	Close() bool
}

type OpenCloser interface {
	Opener
	Closer
}

/* implementation */
type Trial struct {
	source string
}

func (t *Trial) Open() (data string, err error) {
	return "dummy " + t.source, nil
}

func (t *Trial) Close() bool {
	fmt.Println("Trial.close() ...")
	return true
}

/* test */
func main() {
	oc := &Trial{"web.xml"}
	if data, err := oc.Open(); err == nil {
		fmt.Println(data)
	} else {
		fmt.Println(err)
	}
	oc.Close()
}
