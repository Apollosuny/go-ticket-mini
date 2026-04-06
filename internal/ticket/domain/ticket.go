package domain

import "time"

type TicketStatus string

const (
	TicketStatusOpen	TicketStatus = "OPEN"
	TicketStatusInProgress TicketStatus = "IN_PROGRESS"
	TicketStatusResolved TicketStatus = "RESOLVED"
)

type Ticket struct {
	ID uint64
	Title string
	Description string
	RequesterEmail string
	Status TicketStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func IsValidTicketStatus(status TicketStatus) bool {
	switch status {
	case TicketStatusOpen, TicketStatusInProgress, TicketStatusResolved:
		return true
	default:
		return false
	}
}