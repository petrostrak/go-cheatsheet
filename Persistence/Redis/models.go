package main

type Locations struct {
	linkedin string
	github   string
	personal string
}

type Metadata struct {
	name                 string
	from                 string
	programmingLanguages []string
	tools                []string
	locations            Locations
	foreignLanguages     []string
}

type Favorites struct {
	food           string
	drink          string
	programingLang string
}

type Person struct {
	kind          string
	metadata      Metadata
	favorites     Favorites
	thinkingAbout []string
	hobbies       []string
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
