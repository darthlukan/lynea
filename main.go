package main

import (
	"fmt"
)

var (
    ProcessStack []Process
    ProcessChanStack []chan
)

type Process struct {
    Pid int
    State string
    Cli string
    Chan chan
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

func main() {
    for {
        // Don't die
    }
}
