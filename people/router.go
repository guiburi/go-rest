package people

import (
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi"
)

func getPerson(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	err,person := GetPerson(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func getPeople(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GetPeople())
}

func createPerson(w http.ResponseWriter, req *http.Request) {
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = chi.URLParam(req, "id")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreatePerson(person))
}

func modifyPerson(w http.ResponseWriter, req *http.Request) {
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = chi.URLParam(req, "id")
	err, people := ModifyPerson(person)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(people)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people)
}

func deletePerson(w http.ResponseWriter, req *http.Request) {
	err, people := DeletePerson(chi.URLParam(req, "id"))
	if err!= nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(people)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people)
}

// A completely separate router for administrator routes
func Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getPeople)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", getPerson)       // GET /articles/123
		r.Post("/",createPerson)       // GET /articles/123
		r.Put("/", modifyPerson)    // PUT /articles/123
		r.Delete("/", deletePerson) // DELETE /articles/123
	})
	return r
}

