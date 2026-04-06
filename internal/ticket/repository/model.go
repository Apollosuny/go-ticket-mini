package repository

import (
	"time"

	"github.com/Apollosuny/go-ticket-mini/internal/ticket/domain"
)

type TicketModel struct {
	ID uint64 `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	RequesterEmail string `gorm:"type:varchar(255);not null;index"`
	Status string `gorm:"type:varchar(50);not null;index"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type CommentModel struct {
	ID uint64 `gorm:"primaryKey;autoIncrement"`
	TicketID uint64 `gorm:"not null, index"`
	AuthorName string `gorm:"type:varchar(255);not null"`
	Message string `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null"`
}

func (TicketModel) TableName() string {
	return "tickets"
}

func (CommentModel) TableName() string {
	return "comments"
}

func toTicketModel(ticket *domain.Ticket) *TicketModel {
	if ticket == nil {
		return nil
	}

	return &TicketModel{
		ID: ticket.ID,
		Title: ticket.Title,
		Description: ticket.Description,
		RequesterEmail: ticket.RequesterEmail,
		Status: string(ticket.Status),
		CreatedAt: ticket.CreatedAt,
		UpdatedAt: ticket.UpdatedAt,
	}
}

func (m *TicketModel) toDomain() *domain.Ticket {
	if m == nil {
		return nil
	}

	return &domain.Ticket{
		ID: m.ID,
		Title: m.Title,
		Description: m.Description,
		RequesterEmail: m.RequesterEmail,
		Status: domain.TicketStatus(m.Status),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func toCommentModel(comment *domain.Comment) *CommentModel {
	if comment == nil {
		return nil
	}

	return &CommentModel{
		ID:         comment.ID,
		TicketID:   comment.TicketID,
		AuthorName: comment.AuthorName,
		Message:    comment.Message,
		CreatedAt:  comment.CreatedAt,
	}
}

func (m *CommentModel) toDomain() *domain.Comment {
	if m == nil {
		return nil
	}

	return &domain.Comment{
		ID:         m.ID,
		TicketID:   m.TicketID,
		AuthorName: m.AuthorName,
		Message:    m.Message,
		CreatedAt:  m.CreatedAt,
	}
}