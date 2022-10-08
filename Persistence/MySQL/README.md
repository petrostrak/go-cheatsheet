## Go persistence with MySQL

In this example, we persist data in MySQL.

To connect to MySQL:
```go
var (
	DB_DRIVER      = "mysql"
	DB_SOURCE      = "mariadb:password@/cheatsheet"
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
Insert an entity to MySQL:
```go
// DB.Query() is used for SELECT queries which return multiple rows.
// DB.QueryRow() is used for SELECT queries which return a single row.
// DB.Exec() is used for statements which don’t return rows (like INSERT and DELETE).

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

func (s *Server) createPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("createPerson() invoked!")

	arg := db.CreatePersonParams{
		Kind:                   "Human",
		PersonsName:            "Petros Trak",
		Origins:                "Athens, Greece",
		ProgrammingLanguages:   "Golang, Java, Javascript ,Rust",
		Tools:                  "Debian Linux, Docker, !# Bash, MySQL, Postgresql, Redis",
		Github:                 "https://github.com/petrostrak",
		Linkedin:               "https://www.linkedin.com/in/petrostrak/",
		Personal:               "https://petrostrak.netlify.app/",
		ForeignLanguages:       "Greek, English, German",
		FavFood:                "Ramen",
		FavDrink:               "Gin",
		FavProgrammingLanguage: "Golang",
		ThinkingAbout:          "gRPC, Concurrency in Go, русский язык",
		Hobbies:                "Coding, Foreign Languages, Video Games",
	}

	_, err := s.store.CreatePerson(r.Context(), arg)
	if err != nil {
		log.Println(err)
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, arg)
}
```
Get all entities in MySQL:
```go
func (s *Server) readAllPersons(w http.ResponseWriter, r *http.Request) {
	log.Println("readAllPersons() invoked!")

	persons, err := s.store.ListPersons(r.Context())
	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, persons)
}
```
Update an entity in MySQL:
```go
func (s *Server) updatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("updatePerson() invoked!")

	arg := db.UpdatePersonParams{
		ID:                     1,
		Kind:                   "Alien",
		PersonsName:            "Petros Trak",
		Origins:                "Athens, Greece",
		ProgrammingLanguages:   "Golang, Java, Javascript ,Rust",
		Tools:                  "Debian Linux, Docker, !# Bash, MySQL, Postgresql, Redis",
		Github:                 "https://github.com/petrostrak",
		Linkedin:               "https://www.linkedin.com/in/petrostrak/",
		Personal:               "https://petrostrak.netlify.app/",
		ForeignLanguages:       "Greek, English, German",
		FavFood:                "Ramen",
		FavDrink:               "Gin",
		FavProgrammingLanguage: "Golang",
		ThinkingAbout:          "gRPC, Concurrency in Go, русский язык",
		Hobbies:                "Coding, Foreign Languages, Video Games",
	}

	_, err := s.store.UpdatePerson(r.Context(), arg)

	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, arg)
}
```
Delete an entity in MySQL:
```go
func (s *Server) deletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("deletePerson() invoked!")

	var id int64 = 1
	err := s.store.DeletePersonById(r.Context(), id)
	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, fmt.Sprintf("Deleted person with id of %d!", id))
}
```