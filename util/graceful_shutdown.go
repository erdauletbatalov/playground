package util

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForShutdownSignal() os.Signal {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGSEGV,
		syscall.SIGQUIT,
	)

	signal := <-sigch

	return signal
}
