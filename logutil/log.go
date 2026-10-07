package logutil

import (
	"fmt"
	"log"
	"os"
)

func NewLogger(service string) *log.Logger {
	pid := os.Getpid()
	prefix := fmt.Sprintf("%d [%s] :", pid, service)
	return log.New(os.Stderr, prefix, log.LstdFlags|log.Lmsgprefix)
}
