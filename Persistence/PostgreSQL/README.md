## Go persistence with PostgreSQL

In this example, we persist data in PostgreSQL.

To connect to PostgreSQL:
```go
var (
	DB_DRIVER      = "postgres"
	DB_SOURCE      = "postgresql://postgres:secret@localhost:5432/cheatsheet?sslmode=disable"
	SERVER_ADDRESS = "0.0.0.0:8001"
)

type Server struct {
	store db.Querier
}

func (s *Server) setupRouter() {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", s.createPerson)
	mux.HandleFunc("/read-all", s.readAllPersons)
	mux.HandleFunc("/update", s.updatePerson)
	mux.HandleFunc("/delete", s.deletePerson)

	err := http.ListenAndServe(SERVER_ADDRESS, mux)
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.New(conn)
	server := &Server{
		store: store,
	}

	server.setupRouter()
}
```
Create an entry in PostgreSQL:
```go
func (s *Server) createPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("createPerson() invoked!")

	arg := db.CreatePersonParams{
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

	person, err := s.store.CreatePerson(r.Context(), arg)
	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, person)
}
```
Read all entries in PostgreSQL:
```go
func (s *Server) readAllPersons(w http.ResponseWriter, r *http.Request) {
	log.Println("readAllPersons() invoked!")

	persons, err := s.store.ListPersons(r.Context(), 5)
	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, persons)
}
```
Update an entry in PostgreSQL:
```go
func (s *Server) updatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("updatePerson() invoked!")

	arg := db.UpdatePersonParams{
		ID:                     0,
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

	person, err := s.store.UpdatePerson(r.Context(), arg)

	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, person)
}
```
Delete an entry in PostgreSQL:
```go
func (s *Server) deletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("deletePerson() invoked!")

	var id int64 = 0
	err := s.store.DeletePersonById(r.Context(), id)
	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, fmt.Sprintf("Deleted person with id of %d!", id))
}
```