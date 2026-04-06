package service

import (
	"context"
	"strings"
	"time"

	"github.com/Apollosuny/go-ticket-mini/internal/ticket/domain"
)

type Repository interface {
	CreateTicket(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error)
	GetTicketByID(ctx context.Context, id uint64) (*domain.Ticket, error)
	ListTickets(ctx context.Context) ([]*domain.Ticket, error)
	UpdateTicketStatus(ctx context.Context, id uint64, status domain.TicketStatus) (*domain.Ticket, error)
	AddComment(ctx context.Context, comment *domain.Comment) (*domain.Comment, error)
}

type Service interface {
	CreateTicket(ctx context.Context, title, description, requesterEmail string) (*domain.Ticket, error)
	GetTicket(ctx context.Context, id uint64) (*domain.Ticket, error)
	ListTickets(ctx context.Context) ([]*domain.Ticket, error)
	UpdateTicketStatus(ctx context.Context, id uint64, status domain.TicketStatus) (*domain.Ticket, error)
	AddComment(ctx context.Context, ticketID uint64, authorName, message string) (*domain.Comment, error)
}

type service struct {
	repo Repository
}

func New(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTicket(ctx context.Context, title, description, requesterEmail string) (*domain.Ticket, error) {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)
	requesterEmail = strings.TrimSpace(requesterEmail)

	if title == "" {
		return nil, domain.ErrTitleRequired
	}
	if requesterEmail == "" {
		return nil, domain.ErrRequesterEmailEmpty
	}

	now := time.Now()

	ticket := &domain.Ticket{
		Title: title,
		Description: description,
		RequesterEmail: requesterEmail,
		Status: domain.TicketStatusOpen,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.repo.CreateTicket(ctx, ticket)
}

func (s *service) GetTicket(ctx context.Context, id uint64) (*domain.Ticket, error) {
	return s.repo.GetTicketByID(ctx, id)
}

func (s *service) ListTickets(ctx context.Context) ([]*domain.Ticket, error) {
	return s.repo.ListTickets(ctx)
}

func (s *service) UpdateTicketStatus(ctx context.Context, id uint64, status domain.TicketStatus) (*domain.Ticket, error) {
	if !domain.IsValidTicketStatus(status) {
		return nil, domain.ErrInvalidTicketStatus
	}
	return s.repo.UpdateTicketStatus(ctx, id, status)
}

func (s *service) AddComment(ctx context.Context, ticketID uint64, authorName, message string) (*domain.Comment, error) {
	authorName = strings.TrimSpace(authorName)
	message = strings.TrimSpace(message)

	if authorName == "" {
		return nil, domain.ErrAuthorNameEmpty
	}
	if message == "" {
		return nil, domain.ErrCommentMessageEmpty
	}

	comment := &domain.Comment{
		TicketID: ticketID,
		AuthorName: authorName,
		Message: message,
		CreatedAt: time.Now(),
	}

	return s.repo.AddComment(ctx, comment)
}