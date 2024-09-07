package main

import (
	"log"
	"net"

	pb "Train_Ticket_Application/proto"

	"google.golang.org/grpc"
)

func main() {

	stablishDbConnection()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTrainTicketServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
