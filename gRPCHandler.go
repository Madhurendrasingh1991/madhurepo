package main

import (
	pb "Train_Ticket_Application/proto"
	"context"
	"log"
)

type server struct {
	pb.UnimplementedTrainTicketServer
}

func (s *server) BookTicket(ctx context.Context, in *pb.BookTicketRequest) (*pb.BookTicketResponse, error) {
	log.Printf("Received: %v", in)
	for _, v := range in.User {
		log.Printf("User: %v", v)
	}
	return &pb.BookTicketResponse{IsTicketBook: "true"}, nil
}
