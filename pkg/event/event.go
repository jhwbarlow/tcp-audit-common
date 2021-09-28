package event

import (
	"fmt"
	"net"
	"time"

	"github.com/jhwbarlow/tcp-audit-common/pkg/socketstate"
	"github.com/jhwbarlow/tcp-audit-common/pkg/tcpstate"
)

// Eventer is an interface which describes objects which emit TCP state change events
type Eventer interface {
	Event() (*Event, error)
}

// EventerCloser is an interface which describes Eventers which must be closed when no
// longer needed in order to free resources they have acquired.
type EventerCloser interface {
	Eventer
	Close() error
}

// Event is a TCP state change event
type Event struct {
	Time                 time.Time
	PIDOnCPU             int
	CommandOnCPU         string
	SourceIP, DestIP     net.IP
	SourcePort, DestPort uint16
	OldState, NewState   tcpstate.State
	SocketInfo           *SocketInfo // nil if eventer is not able to provide socket info
}

func (e *Event) String() string {
	socketInfo := "<not available>"
	if e.SocketInfo != nil {
		socketInfo = e.SocketInfo.String()
	}

	return fmt.Sprintf("PID (on CPU): %d, Command (on CPU): %s, Source Port: %v:%d, Destination Port: %v:%d, Old State: %v, New State: %v, Socket Info: [%s]",
		e.PIDOnCPU,
		e.CommandOnCPU,
		e.SourceIP,
		e.SourcePort,
		e.DestIP,
		e.DestPort,
		e.OldState,
		e.NewState,
		socketInfo)
}

func (e *Event) Equal(event *Event) bool {
	if e == event {
		return true
	}

	equal := event.Time == e.Time &&
		event.PIDOnCPU == e.PIDOnCPU &&
		event.CommandOnCPU == e.CommandOnCPU &&
		event.SourceIP.Equal(e.SourceIP) &&
		event.DestIP.Equal(e.DestIP) &&
		event.SourcePort == e.SourcePort &&
		event.DestPort == e.DestPort &&
		event.OldState == e.OldState &&
		event.NewState == e.NewState

	switch {
	case e.SocketInfo != nil && event.SocketInfo != nil:
		return equal && event.SocketInfo.Equal(e.SocketInfo)
	case e.SocketInfo == nil && event.SocketInfo == nil:
		return equal
	default: // One is nil but the other is not
		return false
	}
}

// SocketInfo contains internal kernel implementation detail of a socket
type SocketInfo struct {
	ID          string
	INode       uint32
	UID, GID    uint32
	SocketState socketstate.State
}

func (si *SocketInfo) String() string {
	return fmt.Sprintf("ID: %s, INode: %d, UID: %d, GID: %d, State: %s",
		si.ID,
		si.INode,
		si.UID,
		si.GID,
		si.SocketState)
}

func (si *SocketInfo) Equal(socketInfo *SocketInfo) bool {
	if si == socketInfo {
		return true
	}

	return *si == *socketInfo
}
