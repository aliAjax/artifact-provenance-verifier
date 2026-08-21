package domain

type JobState struct {
	ID, Status, LeaseOwner string
	Attempts               int
}
