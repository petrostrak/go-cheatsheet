package main

import (
	"context"
	"fmt"
	"log"
	pb "mongo-db/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreatePerson(ctx context.Context, in *pb.Person) (*pb.PersonId, error) {
	log.Printf("CreatePerson() invoked with %v\n", in)

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

func (s *Server) ReadPerson(ctx context.Context, in *pb.PersonId) (*pb.Person, error) {
	log.Printf("ReadPerson() invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("error parse ID: %v\n", err),
		)
	}

	data := &Person{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	err = res.Decode(data)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("cannot find person with the ID provided: %v\n", err),
		)
	}

	return documentToPerson(data), nil
}

func (s *Server) UpdatePerson(ctx context.Context, in *pb.Person) (*emptypb.Empty, error) {
	log.Printf("UpdatePerson() invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("error parse ID: %v\n", err),
		)
	}

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

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find person with Id",
		)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeletePerson(ctx context.Context, in *pb.PersonId) (*emptypb.Empty, error) {
	log.Printf("DeletePerson() invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("error parse ID: %v\n", err),
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Could not delete object in mongoDB: %v\n", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Person was not found",
		)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListPerson(in *emptypb.Empty, stream pb.PersonService_ListPersonServer) error {
	log.Println("ListPerson() invoked")

	cur, err := collection.Find(context.Background(), primitive.D{{}}) // empty filter
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &Person{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v\n", err),
			)
		}

		stream.Send(documentToPerson(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}

	return nil
}
