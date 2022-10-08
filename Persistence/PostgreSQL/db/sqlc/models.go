package db

type Person struct {
	ID                     int64    `json:"id"`
	Kind                   string   `json:"kind"`
	PersonsName            string   `json:"persons_name"`
	Origins                string   `json:"origins"`
	ProgrammingLanguages   []string `json:"programming_languages"`
	Tools                  []string `json:"tools"`
	Linkedin               string   `json:"linkedin"`
	Github                 string   `json:"github"`
	Personal               string   `json:"personal"`
	ForeignLanguages       []string `json:"foreign_languages"`
	FavFood                string   `json:"fav_food"`
	FavDrink               string   `json:"fav_drink"`
	FavProgrammingLanguage string   `json:"fav_programming_language"`
	ThinkingAbout          string   `json:"thinking_about"`
	Hobbies                string   `json:"hobbies"`
}
