package main

import (
	"os"
	"log"
)

type Job struct {
	Command string
	*log.Logger
}

func main() {
	job := &Job{"greeting", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Println("starting now")
}
