package main

import (
	"context"
	"io"
	"log"
	pb "mongo-db/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

func createPerson(c pb.PersonServiceClient) string {
	log.Println("createPerson() invoked!")

	person := &pb.Person{
		Kind:                   "Human",
		PersonsName:            "Petros Trak",
		Origins:                "Athens, Greece",
		ProgrammingLanguages:   []string{"Golang", "Java", "Javascript", "Rust"},
		Tools:                  []string{"Debian Linux", "Docker", "!# Bash", "MySQL", "Postgresql", "Redis"},
		Github:                 "https://github.com/petrostrak",
		Linkedin:               "https://www.linkedin.com/in/petrostrak/",
		Personal:               "https://petrostrak.netlify.app/",
		ForeignLanguages:       []string{"Greek", "English", "German"},
		FavFood:                "Ramen",
		FavDrink:               "Gin",
		FavProgrammingLanguage: "Golang",
		ThinkingAbout:          []string{"gRPC", "Concurrency in Go", "русский язык"},
		Hobbies:                []string{"Coding", "Foreign Languages", "Video Games"},
	}

	res, err := c.CreatePerson(context.Background(), person)
	if err != nil {
		log.Printf("unexpected err %v\n", err)
	}

	log.Printf("Person has been created with a new id of %s.\n", res.Id)

	return res.Id

}

func readPerson(c pb.PersonServiceClient, id string) *pb.Person {
	log.Println("readPerson() invoked!")

	req := &pb.PersonId{Id: id}
	res, err := c.ReadPerson(context.Background(), req)
	if err != nil {
		log.Printf("err while reading %v\n", err)
	}

	log.Printf("Person was read: %v\n", res)

	return res
}

func updatePerson(c pb.PersonServiceClient, id string) {
	log.Println("updatePerson() invoked!")

	person := &pb.Person{
		Kind:                   "Alien",
		PersonsName:            "Petros Trak",
		Origins:                "Athens, Greece",
		ProgrammingLanguages:   []string{"Golang", "Java", "Javascript", "Rust"},
		Tools:                  []string{"Debian Linux", "Docker", "!# Bash", "MySQL", "Postgresql", "Redis"},
		Github:                 "https://github.com/petrostrak",
		Linkedin:               "https://www.linkedin.com/in/petrostrak/",
		Personal:               "https://petrostrak.netlify.app/",
		ForeignLanguages:       []string{"Greek", "English", "German"},
		FavFood:                "Ramen",
		FavDrink:               "Gin",
		FavProgrammingLanguage: "Golang",
		ThinkingAbout:          []string{"gRPC", "Concurrency in Go", "русский язык"},
		Hobbies:                []string{"Coding", "Foreign Languages", "Video Games"},
	}

	_, err := c.UpdatePerson(context.Background(), person)
	if err != nil {
		log.Printf("err while updating %v\n", err)
	}

	log.Println("Person was updated!")
}

func listPerson(c pb.PersonServiceClient) {
	log.Println("listPerson() invoked!")

	stream, err := c.ListPerson(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Error while calling listPerson: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Error while receiving from listPerson: %v\n", err)
		}

		log.Println(res)
	}
}

func deletePerson(c pb.PersonServiceClient, id string) {
	log.Println("deleteBlog() invoked!")

	_, err := c.DeletePerson(context.Background(), &pb.PersonId{Id: id})
	if err != nil {
		log.Printf("Error while deleting: %v\n", err)
	}

	log.Printf("Person with id of %s deleted!\n", id)
}
