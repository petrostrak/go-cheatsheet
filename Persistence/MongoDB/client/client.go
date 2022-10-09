package main

import (
	"context"
	"log"
	pb "mongo-db/proto"
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
