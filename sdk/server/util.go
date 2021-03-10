package server

import (
	"os"
	"os/signal"
	"syscall"
)

const (
	// for convenience
	SIGHUP  = syscall.SIGHUP
	SIGINT  = syscall.SIGINT
	SIGTERM = syscall.SIGTERM
	SIGUSR1 = syscall.SIGUSR1
	SIGUSR2 = syscall.SIGUSR2
)

// wait for terminate signals
// this method can only be used at one place at then same time
func WaitForSignals(sigs ...os.Signal) <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGPIPE)
	signal.Notify(c, sigs...)
	defer signal.Stop(c)
	return c
}
