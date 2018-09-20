package main

import (
	"log"
	"os"
	"net/http"
	"time"
	"flag"
	"github.com/go-chi/chi"
	"go-rest/people"
	"github.com/go-chi/chi/middleware"
)

var opts struct{
	port string
}

func init() {
	flag.StringVar(&opts.port, "p", "8081", "web server port")
	flag.Parse()
	initDb()
}

func initDb() {
	people.CreatePerson(people.Person{ID: "1", Firstname: "paul", Lastname: "stanley",
		Address: &people.Address{City: "Queens", State: "NY"}})
	people.CreatePerson(people.Person{ID: "2", Firstname: "gene", Lastname: "simmons"})
	people.CreatePerson(people.Person{ID: "3", Firstname: "tommy", Lastname: "thayer",
		Address: &people.Address{City: "Portland",State:"OR"}})
	people.CreatePerson(people.Person{ID: "4", Firstname: "eric", Lastname: "singer"})
}


func main() {
	l := log.New(os.Stdout, "go-rest ", log.LstdFlags|log.Lshortfile)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Mount("/people",people.Routes())

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong\n"))
	})

	s := &http.Server{
		Addr:           ":" + opts.port,
		Handler:        r,
		ErrorLog:       l,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	host, _:= os.Hostname()
	l.Printf("%s - Starting server on port %v", host, opts.port)
	if err := s.ListenAndServe(); err != nil {
		l.Fatalf("Could not listen on %s: %v\n", opts.port, err)
	}
}



//TODO:
// - request validation
// - error json response
