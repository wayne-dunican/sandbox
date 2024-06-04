package main

import "time"

type HeartbeatMessage struct {
	Timestamp time.Time
	Status    string
}
