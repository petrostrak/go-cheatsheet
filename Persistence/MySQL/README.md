## Go persistence with MySQL

In this example, we persist data in MySQL.

To connect to MySQL:
```go
func main() {
    ...
    db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// We also defer a call to db.Close() so that the connection pool is closed
	// before the main() returns.
	defer db.Close()
}

// The openDB() function wraps sql.Open() and returns an sql.DB connection pool
// for a given DSN
func openDB(dsn string) (*sql.DB, error) {

	// The sql.Open() function doesn’t actually create any connections, all
	// it does is initialize the pool for future use. Actual connections to the
	// database are established lazily.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
```
