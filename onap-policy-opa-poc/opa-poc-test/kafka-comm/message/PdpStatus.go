package main

import (
	"fmt"
)

// PdpMessageType represents the type of PDP message.
type PdpMessageType int

// Enumerate the possible PDP message types
const (
	PDP_STATUS PdpMessageType = iota
	PDP_UPDATE
	PDP_STATE_CHANGE
	PDP_HEALTH_CHECK
	PDP_TOPIC_CHECK
)

// String representation of PdpMessageType
func (msgType PdpMessageType) String() string {
	switch msgType {
	case PDP_STATUS:
		return "PDP_STATUS"
	case PDP_UPDATE:
		return "PDP_UPDATE"
	case PDP_STATE_CHANGE:
		return "PDP_STATE_CHANGE"
	case PDP_HEALTH_CHECK:
		return "PDP_HEALTH_CHECK"
	case PDP_TOPIC_CHECK:
		return "PDP_TOPIC_CHECK"
	default:
		return fmt.Sprintf("Unknown PdpMessageType: %d", msgType)
	}
}

// PdpStatus represents the PDP_STATUS message sent to PAP.
type PdpStatus struct {
	MessageType        PdpMessageType
	PdpType            string
	State              PdpState
	Healthy            PdpHealthStatus
	Description        string
	Policies           []ToscaConceptIdentifier
	DeploymentInstance string
	Properties         string
	Response           PdpResponseDetails
}

func main() {
	// Example usage
	fmt.Println(PDP_STATUS.String())
	fmt.Println(Healthy.String())

	var statusMsg = PdpStatus{
		MessageType: PDP_STATUS,
		PdpType:     "opa",
		State:       Passive,
		Healthy:     Healthy,
		Description: "opa pdp",
	}

	fmt.Println(statusMsg)
}
