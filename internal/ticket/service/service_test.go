package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Apollosuny/go-ticket-mini/internal/ticket/domain"
)

type mockRepository struct {
	createTicketFn       func(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error)
	getTicketByIDFn      func(ctx context.Context, id uint64) (*domain.Ticket, error)
	listTicketsFn        func(ctx context.Context) ([]*domain.Ticket, error)
	updateTicketStatusFn func(ctx context.Context, id uint64, status domain.TicketStatus) (*domain.Ticket, error)
	addCommentFn         func(ctx context.Context, comment *domain.Comment) (*domain.Comment, error)
}

func (m *mockRepository) CreateTicket(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	if m.createTicketFn == nil {
		return nil, errors.New("createTicketFn is not implemented")
	}
	return m.createTicketFn(ctx, ticket)
}

func (m *mockRepository) GetTicketByID(ctx context.Context, id uint64) (*domain.Ticket, error) {
	if m.getTicketByIDFn == nil {
		return nil, errors.New("getTicketByIDFn is not implemented")
	}
	return m.getTicketByIDFn(ctx, id)
}

func (m *mockRepository) ListTickets(ctx context.Context) ([]*domain.Ticket, error) {
	if m.listTicketsFn == nil {
		return nil, errors.New("listTicketsFn is not implemented")
	}
	return m.listTicketsFn(ctx)
}

func (m *mockRepository) UpdateTicketStatus(ctx context.Context, id uint64, status domain.TicketStatus) (*domain.Ticket, error) {
	if m.updateTicketStatusFn == nil {
		return nil, errors.New("updateTicketStatusFn is not implemented")
	}
	return m.updateTicketStatusFn(ctx, id, status)
}

func (m *mockRepository) AddComment(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	if m.addCommentFn == nil {
		return nil, errors.New("addCommentFn is not implemented")
	}
	return m.addCommentFn(ctx, comment)
}

func TestCreateTicket_Success(t *testing.T) {
	repo := &mockRepository{
		createTicketFn: func(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
			if ticket.Title != "Cannot login" {
				t.Fatalf("expected title %q, got %q", "Cannot login", ticket.Title)
			}
			if ticket.RequesterEmail != "user@example.com" {
				t.Fatalf("expected requester email %q, got %q", "user@example.com", ticket.RequesterEmail)
			}
			if ticket.Status != domain.TicketStatusOpen {
				t.Fatalf("expected status %q, got %q", domain.TicketStatusOpen, ticket.Status)
			}

			ticket.ID = 1
			return ticket, nil
		},
	}

	svc := New(repo)

	got, err := svc.CreateTicket(context.Background(), "Cannot login", "I cannot login to the system", "user@example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got == nil {
		t.Fatal("expected ticket, got nil")
	}
	if got.ID != 1 {
		t.Fatalf("expected id %d, got %d", 1, got.ID)
	}
	if got.Status != domain.TicketStatusOpen {
		t.Fatalf("expected status %q, got %q", domain.TicketStatusOpen, got.Status)
	}
}

func TestCreateTicket_TitleRequired(t *testing.T) {
	repo := &mockRepository{
		createTicketFn: func(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
			t.Fatal("repository should not be called when title is empty")
			return nil, nil
		},
	}

	svc := New(repo)

	got, err := svc.CreateTicket(context.Background(), "   ", "desc", "user@example.com")
	if !errors.Is(err, domain.ErrTitleRequired) {
		t.Fatalf("expected ErrTitleRequired, got %v", err)
	}
	if got != nil {
		t.Fatalf("expected nil ticket, got %#v", got)
	}
}

func TestUpdateTicketStatus_InvalidStatus(t *testing.T) {
	repo := &mockRepository{
		updateTicketStatusFn: func(ctx context.Context, id uint64, status domain.TicketStatus) (*domain.Ticket, error) {
			t.Fatal("repository should not be called when status is invalid")
			return nil, nil
		},
	}

	svc := New(repo)

	got, err := svc.UpdateTicketStatus(context.Background(), 1, domain.TicketStatus("INVALID"))
	if !errors.Is(err, domain.ErrInvalidTicketStatus) {
		t.Fatalf("expected ErrInvalidTicketStatus, got %v", err)
	}
	if got != nil {
		t.Fatalf("expected nil ticket, got %#v", got)
	}
}

func TestAddComment_Success(t *testing.T) {
	repo := &mockRepository{
		addCommentFn: func(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
			if comment.TicketID != 1 {
				t.Fatalf("expected ticket id %d, got %d", 1, comment.TicketID)
			}
			if comment.AuthorName != "Support Agent" {
				t.Fatalf("expected author name %q, got %q", "Support Agent", comment.AuthorName)
			}
			if comment.Message != "We are checking this issue" {
				t.Fatalf("expected message %q, got %q", "We are checking this issue", comment.Message)
			}

			comment.ID = 10
			return comment, nil
		},
	}

	svc := New(repo)

	got, err := svc.AddComment(context.Background(), 1, "Support Agent", "We are checking this issue")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got == nil {
		t.Fatal("expected comment, got nil")
	}
	if got.ID != 10 {
		t.Fatalf("expected comment id %d, got %d", 10, got.ID)
	}
}

func TestAddComment_MessageRequired(t *testing.T) {
	repo := &mockRepository{
		addCommentFn: func(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
			t.Fatal("repository should not be called when message is empty")
			return nil, nil
		},
	}

	svc := New(repo)

	got, err := svc.AddComment(context.Background(), 1, "Support Agent", "   ")
	if !errors.Is(err, domain.ErrCommentMessageEmpty) {
		t.Fatalf("expected ErrCommentMessageEmpty, got %v", err)
	}
	if got != nil {
		t.Fatalf("expected nil comment, got %#v", got)
	}
}