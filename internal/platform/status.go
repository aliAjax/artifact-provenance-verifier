package platform

import "fmt"

type Status string

const (
	StatusActive    Status = "active"
	StatusIsolated  Status = "isolated"
	StatusWithdrawn Status = "withdrawn"
	StatusRestored  Status = "restored"
)

func (s Status) Valid() bool {
	switch s {
	case StatusActive, StatusIsolated, StatusWithdrawn, StatusRestored:
		return true
	}
	return false
}
func Transition(from, to Status) error {
	if !to.Valid() {
		return fmt.Errorf("invalid target status")
	}
	if from == StatusWithdrawn && to != StatusRestored {
		return fmt.Errorf("withdrawn resource requires restore")
	}
	return nil
}
