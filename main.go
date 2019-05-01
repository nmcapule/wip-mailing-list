package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	// "google.golang.org/appengine"

	_ "github.com/lib/pq"
)

var port = flag.Int("port", 8080, "Port to serve the app")

var (
	once sync.Once
	db   *sql.DB
)

func connectDB() (*sql.DB, error) {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", os.Getenv("POSTGRES_CONNECTION"))
		if err != nil {
			panic(err)
		}
	})
	return db, nil
}

type mailingList struct {
	email     string    `db:"email"`
	timestamp time.Time `db:"timestamp"`
}

func mailingListHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	if email != "" {
		db, err := connectDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = db.Exec("INSERT INTO mailing_list (email) VALUES ($1)", email)
		if err != nil {
			// If error happens on insert, log error and carry on.
			fmt.Println("Failed to insert into mailing_list:", err)
		}
	}
	http.ServeFile(w, r, "static/mailing-success.html")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, fmt.Sprintf("static%s", r.URL.Path))
	}
}

func main() {
	fmt.Printf("Running on port %d...", *port)

	http.HandleFunc("/_ah/health", healthCheckHandler)
	http.HandleFunc("/mailing", mailingListHandler)
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
	// appengine.Main()
}
