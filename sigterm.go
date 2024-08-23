package sigterm

import (
	"encoding"
	"fmt"
	"strings"
	"syscall"
)

// Signal is a termination signal that is used for initiating the shutdown procedure.
// This type implements the encoding.TextMarshaller and encoding.TextUnmarshaler interfaces
// for text-based serialization and deserialization. So, it can be used for configuration
// management using environment variables or flag.TextVar().
//
// List of termination signals can be found in the GNU libc manual:
// - https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
type Signal syscall.Signal

// Supported shutdown signals.
const (
	SIGINT  = Signal(syscall.SIGINT)
	SIGHUP  = Signal(syscall.SIGHUP)
	SIGTERM = Signal(syscall.SIGTERM)
	SIGQUIT = Signal(syscall.SIGQUIT)
	SIGKILL = Signal(syscall.SIGKILL)
)

var _signalNames = map[Signal]string{
	SIGINT:  "SIGINT",
	SIGHUP:  "SIGHUP",
	SIGTERM: "SIGTERM",
	SIGQUIT: "SIGQUIT",
	SIGKILL: "SIGKILL",
}

var (
	_ encoding.TextMarshaler   = (*Signal)(nil)
	_ encoding.TextUnmarshaler = (*Signal)(nil)
)

// Unwrap returns the syscall.Signal value of the sigterm.Signal.
func (s Signal) Unwrap() syscall.Signal {
	return syscall.Signal(s)
}

// String returns the string representation of the signal.
func (s Signal) String() string {
	if name, ok := _signalNames[s]; ok {
		return name
	}
	return fmt.Sprintf("sigterm.Signal(%d): unknown signal", s)
}

func (s Signal) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *Signal) UnmarshalText(b []byte) error {
	text := string(b)
	for sig, name := range _signalNames {
		if strings.EqualFold(text, name) {
			*s = sig
			return nil
		}
	}
	return fmt.Errorf("unknown termination signal: %s", text)
}

type signalValue interface {
	syscall.Signal | Signal
}

// IsTermination returns true if the given sigterm.Signal or syscall.Signal is a termination signal.
func IsTermination[T signalValue](sig T) bool {
	switch t := any(sig).(type) {
	case Signal:
		return IsTermination(t.Unwrap())
	case syscall.Signal:
		switch t {
		case
			syscall.SIGINT,
			syscall.SIGHUP,
			syscall.SIGTERM,
			syscall.SIGQUIT,
			syscall.SIGKILL:
			return true
		}
	}
	return false
}
