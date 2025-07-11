package domain

// EventRepository is an interface for event repositories
type EventRepository interface {
	GetEvent(eventPagePath string) (VlrEvent, error)
}
