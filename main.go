package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	ProcessStack []Process
	sigChan      = make(chan os.Signal, 1)
)

const (
	DELAY      = 10
	initSocket = "/run/lynea/init"
)

type Process struct {
	Pid        int
	State, Cli string
}

func RouteCommand(state, cmd string) error {
	// Which command are we?
	var err error
	var process Process

	switch strings.ToLower(state) {
	case "start":
		process, err = Start()
	case "restart":
		process, err = Restart()
	case "stop":
		process, err = Stop()
	default:
		err = errors.New("Nothing to do, is this really an error?")
	}

	if err == nil {
		ProcessStack = append(ProcessStack, process)
	}

	return err
}

func Exec() (Process, error) {
	// Execute command
}

func Fork() (Process, error) {
	// Make a channel for the supplied service arg
	// append it to ProcessChanStack
}

func Reboot() {
	// Reboot system
	// /sbin/reboot
}

func Shutdown() {
	// Shutdown system
	// /sbin/halt
}

func Poweroff() {
	// Poweroff system
	// /sbin/poweroff
}

func Start() (Process, error) {
	// Start Process
}

func Restart() (Process, error) {
	// Restart Process
}

func Stop() (Process, error) {
	// Stop Process
}

func GetBaseServices() {
	// Services required to have a minimally running system
	// Defined in /etc/lynea/services/base
}

func DesiredServices() {
	// Read from /etc/lynea/services/user_defined
	// and Fork
}

func StartupSystem() {

	// PID 1
	// Socket dir /run
	// Mount virtual filesystems
	// Mount real filesystems
	// Set $HOSTNAME (/proc/sys/kernel/hostname)
	// Create tmpfiles
	// Spawn TTYs
	// Exec (Fork) base services
}

func MkNamedPipe(name string) error {
	// syscall.Mkfifo(path string, mode uint32) (error)
	return nil
}

func main() {
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGTERM)
	signal.Notify(sigChan, syscall.SIGKILL)

	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal: %v\n", sig)
	}()

	for {
		// Read named pipe
		// route input to appropriate function
		// continue listening
	}
}
