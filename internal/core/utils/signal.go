package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitSignal(fn func(os.Signal)) {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)

	sig := <-termChan

	fn(sig)

	/*
		ch := make(chan os.Signal, 2)
		signal.Notify(
			ch,
			syscall.SIGINT,
			syscall.SIGQUIT,
			syscall.SIGTERM,
		)
		for {
			sig := <-ch
			switch sig {
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				return sig
			}
		}
	*/
}
