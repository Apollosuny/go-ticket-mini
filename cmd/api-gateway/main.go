package main

import (
	"log"
	"net/http"
	"strconv"

	ticketpb "github.com/Apollosuny/go-ticket-mini/api/proto"
	"github.com/Apollosuny/go-ticket-mini/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type createTicketBody struct {
	Title string `json:"title"`
	Description string `json:"description"`
	RequesterEmail string `json:"requester_email"`
}

type updateTicketStatusBody struct {
	Status string `json:"status"`
}

type addCommentBody struct {
	AuthorName string `json:"author_name"`
	Message string `json:"message"`
}

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	conn, err := grpc.NewClient(
		cfg.TicketServiceGRPCAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("connect ticket service grpc: %v", err)
	}
	defer conn.Close()

	ticketClient := ticketpb.NewTicketServiceClient(conn)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	api := router.Group("/api/v1")
	{
		tickets := api.Group("/tickets")
		{
			tickets.POST("", func(c *gin.Context) {
				var body createTicketBody
				if err := c.ShouldBindJSON(&body); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					return
				}

				resp, err := ticketClient.CreateTicket(c.Request.Context(), &ticketpb.CreateTicketRequest{
					Title:          body.Title,
					Description:    body.Description,
					RequesterEmail: body.RequesterEmail,
				})
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				}

				c.JSON(http.StatusCreated, resp)
			})

			tickets.GET("", func(c *gin.Context) {
				resp, err := ticketClient.ListTickets(c.Request.Context(), &ticketpb.ListTicketsRequest{})
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				}

				c.JSON(http.StatusOK, resp)
			})

			tickets.GET("/:id", func(c *gin.Context) {
				id, err := strconv.ParseUint(c.Param("id"), 10, 64)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "invalid ticket id",
					})
					return
				}

				resp, err := ticketClient.GetTicket(c.Request.Context(), &ticketpb.GetTicketRequest{
					Id: id,
				})
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				}

				c.JSON(http.StatusOK, resp)
			})

			tickets.PATCH("/:id/status", func(c *gin.Context) {
				id, err := strconv.ParseUint(c.Param("id"), 10, 64)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "invalid ticket id",
					})
					return
				}

				var body updateTicketStatusBody
				if err := c.ShouldBindJSON(&body); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					return
				}

				resp, err := ticketClient.UpdateTicketStatus(c.Request.Context(), &ticketpb.UpdateTicketStatusRequest{
					Id:     id,
					Status: body.Status,
				})
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				}

				c.JSON(http.StatusOK, resp)
			})

			tickets.POST("/:id/comments", func(c *gin.Context) {
				id, err := strconv.ParseUint(c.Param("id"), 10, 64)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "invalid ticket id",
					})
					return
				}

				var body addCommentBody
				if err := c.ShouldBindJSON(&body); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					return
				}

				resp, err := ticketClient.AddComment(c.Request.Context(), &ticketpb.AddCommentRequest{
					TicketId:   id,
					AuthorName: body.AuthorName,
					Message:    body.Message,
				})
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				}

				c.JSON(http.StatusCreated, resp)
			})
		}
	}

	log.Printf("api gateway listening on port %s", cfg.APIGatewayPort)

	if err := router.Run(":" + cfg.APIGatewayPort); err != nil {
		log.Fatalf("run api gateway: %v", err)
	}
}