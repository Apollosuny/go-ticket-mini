package grpc

import (
	"context"

	ticketpb "github.com/Apollosuny/go-ticket-mini/api/proto"
	"github.com/Apollosuny/go-ticket-mini/internal/ticket/domain"
	"github.com/Apollosuny/go-ticket-mini/internal/ticket/service"
)

type Server struct {
	ticketpb.UnimplementedTicketServiceServer
	service service.Service
}

func NewServer(svc service.Service) *Server {
	return &Server{
		service: svc,
	}
}

func (s *Server) CreateTicket(ctx context.Context, req *ticketpb.CreateTicketRequest) (*ticketpb.CreateTicketResponse, error) {
	ticket, err := s.service.CreateTicket(ctx, req.Title, req.Description, req.RequesterEmail)
	if err != nil {
		return nil, err
	}

	return &ticketpb.CreateTicketResponse{
		Ticket: &ticketpb.Ticket{
			Id:             ticket.ID,
			Title:          ticket.Title,
			Description:    ticket.Description,
			RequesterEmail: ticket.RequesterEmail,
			Status:         string(ticket.Status),
			CreatedAt:      ticket.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:      ticket.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

func (s *Server) GetTicket(ctx context.Context, req *ticketpb.GetTicketRequest) (*ticketpb.GetTicketResponse, error) {
	ticket, err := s.service.GetTicket(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &ticketpb.GetTicketResponse{
		Ticket: &ticketpb.Ticket{
			Id:             ticket.ID,
			Title:          ticket.Title,
			Description:    ticket.Description,
			RequesterEmail: ticket.RequesterEmail,
			Status:         string(ticket.Status),
			CreatedAt:      ticket.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:      ticket.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

func (s *Server) ListTickets(ctx context.Context, req *ticketpb.ListTicketsRequest) (*ticketpb.ListTicketsResponse, error) {
	tickets, err := s.service.ListTickets(ctx)
	if err != nil {
		return nil, err
	}

	pbTickets := make([]*ticketpb.Ticket, 0, len(tickets))
	for _, ticket := range tickets {
		pbTickets = append(pbTickets, &ticketpb.Ticket{
			Id: ticket.ID,
			Title:          ticket.Title,
			Description:    ticket.Description,
			RequesterEmail: ticket.RequesterEmail,
			Status:         string(ticket.Status),
			CreatedAt:      ticket.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:      ticket.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &ticketpb.ListTicketsResponse{
		Tickets: pbTickets,
	}, nil
}

func (s *Server) UpdateTicketStatus(ctx context.Context, req *ticketpb.UpdateTicketStatusRequest) (*ticketpb.UpdateTicketStatusResponse, error) {
	ticket, err := s.service.UpdateTicketStatus(ctx, req.Id, domain.TicketStatus(req.Status))
	if err != nil {
		return nil, err
	}
	return &ticketpb.UpdateTicketStatusResponse{
		Ticket: &ticketpb.Ticket{
			Id:             ticket.ID,
			Title:          ticket.Title,
			Description:    ticket.Description,
			RequesterEmail: ticket.RequesterEmail,
			Status:         string(ticket.Status),
			CreatedAt:      ticket.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:      ticket.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

func (s *Server) AddComment(ctx context.Context, req *ticketpb.AddCommentRequest) (*ticketpb.AddCommentResponse, error) {
	comment, err := s.service.AddComment(ctx, req.TicketId, req.AuthorName, req.Message)
	if err != nil {
		return nil, err
	}

	return &ticketpb.AddCommentResponse{
		Comment: &ticketpb.Comment{
			Id:         comment.ID,
			TicketId:   comment.TicketID,
			AuthorName: comment.AuthorName,
			Message:    comment.Message,
			CreatedAt:  comment.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}