package main

import "fmt"

// PdpState represents the possible values for the state of PDP.
type PdpState int

// Enumerate the possible PDP states
const (
	Passive PdpState = iota
	Safe
	Test
	Active
	Terminated
)

// String representation of PdpState
func (state PdpState) String() string {
	switch state {
	case Passive:
		return "PASSIVE"
	case Safe:
		return "SAFE"
	case Test:
		return "TEST"
	case Active:
		return "ACTIVE"
	case Terminated:
		return "TERMINATED"
	default:
		return fmt.Sprintf("Unknown PdpState: %d", state)
	}
}
