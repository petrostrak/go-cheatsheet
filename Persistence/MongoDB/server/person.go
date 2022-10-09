package main

import (
	pb "mongo-db/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID                          primitive.ObjectID `bson:"_id,omitempty"`
	Kind                        string             `bson:"kind"`
	PersonsName                 string             `bson:"persons_name"`
	Origins                     string             `bson:"origins"`
	ProgrammingLanguages        []string           `bson:"programming_languages"`
	Tools                       []string           `bson:"tools"`
	Linkedin                    string             `bson:"linkedin"`
	Github                      string             `bson:"github"`
	Personal                    string             `bson:"personal"`
	ForeignLanguages            []string           `bson:"foreign_languages"`
	FavoriteFood                string             `bson:"favorite_food"`
	FavoriteDrink               string             `bson:"favorite_food"`
	FavoriteProgrammingLanguage string             `bson:"favorite_programming_language"`
	ThinkingAbout               []string           `bson:"thinking_about"`
	Hobbies                     []string           `bson:"hobbies"`
}

func documentToPerson(data *Person) *pb.Person {
	return &pb.Person{
		Id:                     data.ID.Hex(),
		Kind:                   data.Kind,
		PersonsName:            data.PersonsName,
		Origins:                data.Origins,
		ProgrammingLanguages:   data.ProgrammingLanguages,
		Tools:                  data.Tools,
		Linkedin:               data.Linkedin,
		Github:                 data.Github,
		Personal:               data.Personal,
		ForeignLanguages:       data.ForeignLanguages,
		FavFood:                data.FavoriteFood,
		FavDrink:               data.FavoriteDrink,
		FavProgrammingLanguage: data.FavoriteProgrammingLanguage,
		ThinkingAbout:          data.ThinkingAbout,
		Hobbies:                data.Hobbies,
	}
}
