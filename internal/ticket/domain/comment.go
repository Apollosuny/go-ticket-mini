package domain

import "time"

type Comment struct {
	ID uint64
	TicketID uint64
	AuthorName string
	Message string
	CreatedAt time.Time
}