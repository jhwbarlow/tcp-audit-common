package socketstate

import "fmt"

// State represents the internal state of a socket in the Linux kernel
type State uint8

// Kernel socket states defined in <uapi/linux/net.h>, enum `socket_state`
const (
	StateFree          State = iota // not allocated
	StateUnconnected                // unconnected to any socket
	StateConnecting                 // in process of connecting
	StateConnected                  // connected to socket
	StateDisconnecting              // in process of disconnecting
)

func FromInt(state uint8) (State, error) {
	switch State(state) {
	case StateFree, StateUnconnected, StateConnecting, StateConnected, StateDisconnecting:
		return State(state), nil
	default:
		return State(0xFF), fmt.Errorf("illegal state integer: %d", state)
	}
}

func (s State) String() string {
	switch s {
	case StateFree:
		return "FREE"
	case StateUnconnected:
		return "UNCONNECTED"
	case StateConnecting:
		return "CONNECTING"
	case StateConnected:
		return "CONNECTED"
	case StateDisconnecting:
		return "DISCONNECTING"
	default:
		panic(fmt.Errorf("illegal socket state: %d", s))
	}
}
