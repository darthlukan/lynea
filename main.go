package main

import (
	"fmt"
)

var (
    ProcessStack []Process
)

type Process struct {
    Pid int
    State string
    Cli string
}

func Fork() {
    // Make a channel for the supplied service arg
    // return that channel to the stack
}

func main() {
	fmt.Printf("Hello there!\n")
}
