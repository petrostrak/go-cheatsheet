package main

import (
	"context"
	"fmt"
	"log"
	pb "mongo-db/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePerson(ctx context.Context, in *pb.Person) (*pb.PersonId, error) {
	log.Printf("CreateBLog() invoked with %v\n", in)

	data := &pb.Person{
		Id:                     in.Id,
		Kind:                   in.Kind,
		PersonsName:            in.PersonsName,
		Origins:                in.Origins,
		ProgrammingLanguages:   in.ProgrammingLanguages,
		Tools:                  in.Tools,
		Linkedin:               in.Linkedin,
		Github:                 in.Github,
		Personal:               in.Personal,
		ForeignLanguages:       in.ForeignLanguages,
		FavFood:                in.FavFood,
		FavDrink:               in.FavDrink,
		FavProgrammingLanguage: in.FavProgrammingLanguage,
		ThinkingAbout:          in.ThinkingAbout,
		Hobbies:                in.Hobbies,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal err: %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("error casting to oid: %v\n", err),
		)
	}

	return &pb.PersonId{Id: oid.Hex()}, nil
}
