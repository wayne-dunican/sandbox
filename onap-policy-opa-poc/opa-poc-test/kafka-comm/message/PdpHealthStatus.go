package main

import "fmt"

// PdpHealthStatus represents the possible values for the health status of PDP.
type PdpHealthStatus int

// Enumerate the possible PDP health statuses
const (
	Healthy PdpHealthStatus = iota
	NotHealthy
	TestInProgress
	Unknown
)

// String representation of PdpHealthStatus
func (status PdpHealthStatus) String() string {
	switch status {
	case Healthy:
		return "HEALTHY"
	case NotHealthy:
		return "NOT_HEALTHY"
	case TestInProgress:
		return "TEST_IN_PROGRESS"
	case Unknown:
		return "UNKNOWN"
	default:
		return fmt.Sprintf("Unknown PdpHealthStatus: %d", status)
	}
}
