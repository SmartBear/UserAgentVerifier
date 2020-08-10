package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/teris-io/shortid"
)

type Data struct {
	ID string `json:"id"`
	Agent string `json:"agent"`
}

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		id, _ := shortid.Generate()
		data := Data{ID:id}
		c.Set(id, &data, cache.DefaultExpiration)
		json.NewEncoder(w).Encode(data)
	})

	router.HandleFunc("/agent/{id}", func(w http.ResponseWriter, r *http.Request){
		id := mux.Vars(r)["id"]
		ua := r.Header.Get("User-Agent")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s:%s", id, ua)
		data := Data{ID:id, Agent:ua}
		c.Set(id, &data, cache.DefaultExpiration)
	})

	router.HandleFunc("/verify/{id}", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		id := mux.Vars(r)["id"]

		if x, found := c.Get(id); found {
			data := x.(*Data)
			json.NewEncoder(w).Encode(data)
		} else {
			json.NewEncoder(w).Encode(Data{ID: id, Agent: "NA"})
		}
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}