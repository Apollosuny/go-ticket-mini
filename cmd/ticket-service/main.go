package main

import (
	"log"
	"net"

	ticketpb "github.com/Apollosuny/go-ticket-mini/api/proto"
	"github.com/Apollosuny/go-ticket-mini/internal/ticket/repository"
	"github.com/Apollosuny/go-ticket-mini/internal/ticket/service"
	ticketgrpc "github.com/Apollosuny/go-ticket-mini/internal/ticket/transport/grpc"
	"github.com/Apollosuny/go-ticket-mini/pkg/config"
	"github.com/Apollosuny/go-ticket-mini/pkg/database"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := database.NewPostgres(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("connection postgres: %v", err)
	}

	repo := repository.New(db)
	svc := service.New(repo)
	grpcServer := ticketgrpc.NewServer(svc)

	lis, err := net.Listen("tcp", ":"+cfg.TicketServicePort)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	server := grpc.NewServer()
	ticketpb.RegisterTicketServiceServer(server, grpcServer)

	log.Printf("ticket service listening on port %s", cfg.TicketServicePort)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("serve grpc: %v", err)
	}
}