package main

import (
    "os"
    "os/signal"
    "net"
    "fmt"
    "syscall"
)

var (
    ProcessStack []Process
    sigChan = make(chan os.Signal, 1)
)

const (
    DELAY = 10
    initSocket = "/run/lynea/init"
)

type Process struct {
    Pid int
    State, Cli string
}

func RouteCommand() {
    // Which command are we?
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

func Dispatcher(socketConn *net.UnixConn) {
    buffer := make([]byte, 1024)
    outOfBand := make([]byte, 1024)

    size, oob, flags, addr, err := socketConn.ReadMsgUnix(buffer, outOfBand)
    if err != nil {
        fmt.Printf("Caught error '%v' while reading from socket\n", err)
    }
    // TODO: Switch...case -> functions
    fmt.Printf("received: '%s' from buffer\n", string(buffer[:size]))
    fmt.Printf("received: '%s' from oob\n", string(outOfBand[:oob]))
    fmt.Printf("received flags: %v, from addr: %v.\n", flags, addr)
}

func main() {
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, syscall.SIGTERM)
    signal.Notify(sigChan, syscall.SIGKILL)

    go func() {
        sig := <-sigChan
        fmt.Printf("Received signal: %v\n", sig)
    }()

    sockAddr, err := net.ResolveUnixAddr("unix", initSocket)
    if err != nil {
        fmt.Printf("Caught error '%v' resolving socket address.\n", err)
    }

    listener, err := net.ListenUnix("unix", sockAddr)
    if err != nil {
        fmt.Printf("Caught error '%v' trying to listen on '%v'\n", err, initSocket)
    }

    for {
        sockConn, err := listener.AcceptUnix()
        if err != nil {
            fmt.Printf("Caught error '%v' listening on '%v'\n", err, initSocket)
        }
        go Dispatcher(sockConn)
    }
}
