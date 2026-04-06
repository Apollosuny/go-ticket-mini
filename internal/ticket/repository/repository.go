package repository

import (
	"context"
	"errors"

	"github.com/Apollosuny/go-ticket-mini/internal/ticket/domain"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateTicket(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	model := toTicketModel(ticket)

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	return model.toDomain(), nil
}

func (r *Repository) GetTicketByID(ctx context.Context, id uint64) (*domain.Ticket, error) {
	var model TicketModel

	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrTicketNotFound
		}
		return nil, err
	}
	return model.toDomain(), nil
}

func (r *Repository) ListTickets(ctx context.Context) ([]*domain.Ticket, error) {
	var models []TicketModel

	if err := r.db.WithContext(ctx).Order("id DESC").Find(&models).Error; err != nil {
		return nil, err
	}
	tickets := make([]*domain.Ticket, 0, len(models))
	for i := range models {
		tickets = append(tickets, models[i].toDomain())
	}
	return tickets, nil
}

func (r *Repository) UpdateTicketStatus(ctx context.Context, id uint64, status domain.TicketStatus) (*domain.Ticket, error) {
	var model TicketModel

	err := r.db.WithContext(ctx).
				First(&model, id).
				Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrTicketNotFound
		}
		return nil, err
	}

	model.Status = string(status)

	if err := r.db.WithContext(ctx).Save(&model).Error; err != nil {
		return nil, err
	}
	return model.toDomain(), nil
}

func (r *Repository) AddComment(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	var ticket TicketModel

	err := r.db.WithContext(ctx).
				First(&ticket, comment.TicketID).
				Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrTicketNotFound
		}
		return nil, err
	}

	model := toCommentModel(comment)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	return model.toDomain(), nil
}