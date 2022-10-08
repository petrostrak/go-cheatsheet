package main

type Locations struct {
	Linkedin string `json:"linkedin"`
	Github   string `json:"github"`
	Personal string `json:"personal_website"`
}

type Metadata struct {
	Name                 string    `json:"name"`
	From                 string    `json:"from"`
	ProgrammingLanguages []string  `json:"programming_languages"`
	Tools                []string  `json:"tools"`
	Locations            Locations `json:"locations"`
	ForeignLanguages     []string  `json:"foreign_languages"`
}

type Favorites struct {
	Food           string `json:"fav_food"`
	Drink          string `json:"fav_drink"`
	ProgramingLang string `json:"fav_programming_language"`
}

type Person struct {
	Kind          string    `json:"kind"`
	Metadata      Metadata  `json:"metadata"`
	Favorites     Favorites `json:"favorites"`
	ThinkingAbout []string  `json:"thinking_about"`
	Hobbies       []string  `json:"hobbies"`
}

// petrosTrak := &AboutMe{
// 	kind: "Human",
// 	metadata: Metadata{
// 		name:    "Petros Trakadas",
// 		from:    "🇬🇷",
// 		programmingLanguages: []string{
// 			"Golang",
// 			"Java",
// 			"Javascript",
// 			"Rust",
// 		},
// 		tools: []string{
// 			"Debian Linux",
// 			"Docker",
// 			"!# Bash",
// 			"MySQL",
// 			"Postgresql",
// 			"Redis",
// 		},
// 		locations: Locations{
// 			github:   "https://github.com/petrostrak",
// 			linkedin: "https://www.linkedin.com/in/petrostrak/",
// 			personal: "https://petrostrak.netlify.app/",
// 		},
// 		foreignLanguages: []string{
// 			"🇬🇷",
// 			"🏴󠁧󠁢󠁥󠁮󠁧󠁿",
// 			"🇩🇪",
// 		},
// 	favorites: Favorites{
// 		food:           "🍣",
// 		drink:          "🍺",
// 		programingLang: "Golang",
// 	},
// 	thinkingAbout: []string{
// 		"gRPC",
// 		"Concurrency in Go",
// 		"русский язык",
// 	},
// 	hobbies: []string{
// 		"Coding",
// 		"Foreign Languages",
// 		"🎮",
// 	},
// }
