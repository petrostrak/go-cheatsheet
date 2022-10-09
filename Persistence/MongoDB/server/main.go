package main

import (
	"log"
	"net"

	pb "mongo-db/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	addr       = "0.0.0.0:50051"
	collection *mongo.Collection
)

type Server struct {
	pb.PersonServiceServer
}

func main() {
	collection = connectToMongoDB()

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s...\n", addr)

	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &Server{})
	reflection.Register(s)

	err = s.Serve(l)
	if err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
