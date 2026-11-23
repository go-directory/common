//go:build godir_debug_common || godir_debug_all

package common

import (
	"log"
)

/*
ConsoleLogger implements a basic console-based logger
using Go's [log] package.
*/
type ConsoleLogger struct{}

func (_ ConsoleLogger) Printf(s string, args ...any) { log.Printf(s, args...) }
func (_ ConsoleLogger) Fatalf(s string, args ...any) { log.Fatalf(s, args...) }

func emitW(m string, v ...any) {
	if LogLevel&logWarning != 0 {
		emit(m, v...)
	}
}

func emitE(m string, v ...any) {
	if LogLevel&logError != 0 {
		emit(m, v...)
	}
}

func emitT(m string, v ...any) {
	if LogLevel&logTrace != 0 {
		emit(m, v...)
	}
}

func emitI(m string, v ...any) {
	if LogLevel&logInfo != 0 {
		emit(m, v...)
	}
}

var emit func(string, ...any)

func init() {
	log.SetFlags(0)
	EventLogger = ConsoleLogger{}
	emit = NewEmitter(EventLogger)
}
