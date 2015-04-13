package main

import (
	// "errors"
	"fmt"
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
		// process, err = Start()
	case "restart":
		// process, err = Restart()
	case "stop":
		// process, err = Stop()
	default:
		// err = errors.New("Nothing to do, is this really an error?")
	}

	if err == nil {
		ProcessStack = append(ProcessStack, process)
	}

	return err
}

func Exec() {
	// Execute command
}

func Fork() {
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

func Start() {
	// Start Process
}

func Restart() {
	// Restart Process
}

func Stop() {
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

func MkNamedPipe() error {
	return syscall.Mkfifo(initSocket, syscall.S_IFIFO|0666)
}

func ReadFromPipe() (string, error) {
	recvData := make([]byte, 100)
	np, err := os.Open(initSocket)
	if err != nil {
		fmt.Printf("ReadFromPipe open: %v\n", err)
	}
	defer np.Close()

	count, err := np.Read(recvData)
	if err != nil {
		fmt.Printf("ReadFromPipe read: %v\n", err)
	}
	data := string(recvData[:count])
	return data, err
}

func ParsePipeData(data string) map[string]string {
	// data is the content of initSocket, make sure to only read the last line sent in
	// should be of structure: <command> <arg>
	splitData := strings.Split(data, "\n")
	lastLine := strings.Split(splitData[len(splitData)-1], " ")

	if len(lastLine) == 2 {
		cmdMap := map[string]string{
			"cmd": lastLine[0],
			"arg": lastLine[1],
		}
		return cmdMap
	}
	return nil
}

func init() {
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGTERM)
	signal.Notify(sigChan, syscall.SIGKILL)

	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal: %v\n", sig)
	}()

	err := MkNamedPipe()
	if err != nil {
		fmt.Printf("Init panic: %v\n", err) // TODO: Don't actually panic, try to drop to a shell or something
	}
}

func main() {

	for {
		// The following blocks, wrap in a goroutine if this becomes a problem
		// Read named pipe
		data, err := ReadFromPipe()
		if err != nil {
			fmt.Printf("Received error: %v and data: %v\n", err, data)
		}
		if len(data) > 0 {
			fmt.Printf("Received data: %v\n", data)
			cmdMap := ParsePipeData(data)

			if cmdMap != nil {
				// route
			}
		}
		// continue listening
	}
}
