package main

import (
	// "errors"
	"bufio"
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
	DELAY = 10
	// initSocket = "/run/lynea/init"
	initSocket = "/home/darthlukan/tmp/lynea"
	confDir    = "/etc/init"
)

type Process struct {
	Pid int
	Cli string
}

type Command struct {
	Type, Arg string
}

func (c Command) IsEmpty() bool {
	if len(c.Type) == 0 && len(c.Arg) == 0 {
		return true
	}
	return false
}

func RouteCommand(cmd Command) error {
	// Which command are we?
	var err error
	var process Process

	switch strings.ToLower(cmd.Type) {
	case "enable":
		// Set service to start with system (copy service.json to /etc/init/services/
	case "disable":
		// remove service.json from /etc/init/services/
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

func ParsePipeData(data string) Command {
	// data is the content of initSocket, make sure to only read the last line sent in
	// should be of structure: <command> <arg>
	splitData := strings.Split(data, " ")
	var cmd Command

	if len(splitData) == 2 {
		cmd.Type = splitData[0]
		cmd.Arg = splitData[1]
		fmt.Printf("returning cmdMap: %v\n", cmd)
	}
	return cmd
}

func PIDOneCheck() bool {
	// Are we PID 1?
	// look in /proc/1/status => Name (first line of file: "Name: $NAME")
	// $NAME == "lynea" => true || false

	pfile, err := os.Open("/proc/1/status")
	if err != nil {
		fmt.Printf("PIDOneCheck caught error: %v\n", err)
		return false
	}
	defer pfile.Close()

	var lines []string

	scanner := bufio.NewScanner(pfile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if strings.Contains(lines[0], "lynea") {
		return true
	}
	return false
}

func init() {
	fmt.Printf("Lynea initializing...\n")
	err := MkNamedPipe()
	if err != nil && err.Error() != "file exists" {
		fmt.Printf("Init panic: %v\n", err) // TODO: Don't actually panic, try to drop to a shell or something
	}

	if pid1 := PIDOneCheck(); pid1 == true {
		// execute bootup, base services, etc
		fmt.Printf("Booting the system...\n")
		StartupSystem() // TODO: Fill this out
	}
}

func main() {
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGTERM)
	signal.Notify(sigChan, syscall.SIGKILL)

	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal: %v\n", sig)
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			os.Exit(1)
		}
	}()

	fmt.Printf("Running...\n")
	for {
		// TODO: Determine if this can be handled safely by a goroutine, also, profile to see how performant it's not.
		data, err := ReadFromPipe()
		if err != nil {
			if err.Error() != "EOF" && len(data) == 0 {
				// We'll get EOF as a normal occurrence but we only care if data is not populated.
				fmt.Printf("Received error: '%v' and data: '%v'\n", err, data)
			}
		}
		if len(data) > 0 {
			fmt.Printf("Received data: %v\n", data)
			cmd := ParsePipeData(data)

			if !cmd.IsEmpty() {
				// route
			}
		}
		// continue listening
	}
}
