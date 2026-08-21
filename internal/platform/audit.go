package platform

import (
	"fmt"
	"time"
)

type AuditEvent struct {
	ID        string
	Type      string
	Actor     string
	Resource  string
	Summary   string
	At        time.Time
	RequestID string
}

func NewAuditEvent(typ, actor, resource, summary, requestID string) AuditEvent {
	return AuditEvent{ID: fmt.Sprintf("evt-%d", time.Now().UnixNano()), Type: typ, Actor: actor, Resource: resource, Summary: SafeMessage(summary), At: time.Now().UTC(), RequestID: requestID}
}
func (e AuditEvent) Valid() bool {
	return e.ID != "" && e.Type != "" && e.Resource != "" && !e.At.IsZero()
}
