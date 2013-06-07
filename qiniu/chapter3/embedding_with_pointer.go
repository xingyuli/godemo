package main

import (
	"log"
	"os"
)

type Job struct {
	Command string
	*log.Logger
}

func (j *Job) Start() {
	j.Println("starting now...")
	j.Println("executing:", j.Command)
	j.Println("started...")
}

func main() {
	j := &Job{"greeting to everyone", log.New(os.Stdout, "[Job] ", log.Ldate|log.Ltime)}
	j.Start()
}
