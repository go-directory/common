package common

/*
Logger is qualified through facilities bearing the methods below.
A given qualifier of this interface can be made active via the
[EventLogger] global variable.

By default, a "null" [Logger] is set as the [EventLogger]. This
means that no events will be emitted to the console. To enable
logging:

  - go build/run with -tags "godir_debug_<pkg_namespace>", or ...
  - go build/run with -tags "godir_debug_all"

Take care to limit use of logging on performance-critical systems.
*/
type Logger interface {
	Printf(string, ...any)
	Fatalf(string, ...any)
}

/*
EventLogger stores the [Logger] instance meant to handle or discard
log events. Users may supply their own concrete [Logger] type if
desired.
*/
var EventLogger Logger

/*
[LogLevel] defines the desired degree of verbosity.

The value should be set in an additive or subtractive manner
according to the following available "verbosity bit values":

  - 1 (error) -- intended for very serious events
  - 2 (warning) -- intended for potentially serious events
  - 4 (trace) -- intended for deep-level events
  - 8 (info) -- intended for succinct informational events

For instance, to enable errors and info, set the value to nine
(9).

To test whether a particular logging level is set, evaluate via:

	LogLevel&uint16(8) != 0 // is "info" logging enabled?

A return value of true indicates the specified uint16 bit is set.
*/
var LogLevel uint16

const (
	// not to exceed uint16
	logError   uint16 = 1 << iota // 1
	logWarning                    // 2
	logTrace                      // 4
	logInfo                       // 8
)

/*
NewEmitter returns an instance of func(string, ...any).

The return value is meant to be used when a particular log
entry is to be written to the specified logging facility l,
and likely only when a particular logging level is set.

The return value internals are as follows:

	func(string, ...any) {
		l.Printf(<timestamp> SPACE <formatted-message>)
	}

If a different format is desired, callers may simply craft
their own constructor(s).

By default, the return value will emit messages whenever it
is called. Callers SHOULD wrap any call of return value within
a logging level conditional statement if emissions are to be
controlled based on a desired verbosity.
*/
func NewEmitter(l Logger) func(string, ...any) {
	return func(m string, v ...any) {
                ts := string(GenTimeNow(true))
                l.Printf(ts + " " + m, v...)
	}
}

