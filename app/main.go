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
	"github.com/DanMHammer/user_agent"
)

type Data struct {
	ID string `json:"id"`

	Agent string `json:"agent"`
	OS string `json:"os"`
	Browser string `json:"browser"`
	Version string `json:"version"`

	ExpectedOS string `json:"expected_os"`
	ExpectedBrowser string `json:"expected_browser"`
	ExpectedVersion string `json:"expected_version"`

	Result bool `json:"result"`
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
		c.Set(id, data, cache.DefaultExpiration)
		json.NewEncoder(w).Encode(data)
	})

	router.HandleFunc("/expect/create", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		id, _ := shortid.Generate()
		var data Data
		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		data.ID = id
		c.Set(id, data, cache.DefaultExpiration)
		json.NewEncoder(w).Encode(data)
	})

	router.HandleFunc("/agent/{id}", func(w http.ResponseWriter, r *http.Request){
		id := mux.Vars(r)["id"]
		ua := r.Header.Get("User-Agent")
		fmt.Fprintf(w, "%s:%s", id, ua)

		if x, found := c.Get(id); found {
			data := x.(Data)
			data.Agent = ua
			ua := user_agent.New(data.Agent)

			os := ua.OS()

			if os == "Windows 8.1" {
				data.OS = "Windows 8"
			} else {
				data.OS = os
			}
			
			browser, version := ua.Browser()
			data.Browser = browser
			data.Version = version

			c.Set(id, data, cache.DefaultExpiration)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	router.HandleFunc("/verify/{id}", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		id := mux.Vars(r)["id"]

		if x, found := c.Get(id); found {
			data := x.(Data)
			json.NewEncoder(w).Encode(data)
		} else {
			json.NewEncoder(w).Encode(Data{ID: id, Agent: "NA"})
		}
	})

	router.HandleFunc("/expect/verify/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := mux.Vars(r)["id"]
		
		if x, found := c.Get(id); found {
			data := x.(Data)
			
			if (data.OS == data.ExpectedOS && data.Browser == data.ExpectedBrowser && data.Version == data.ExpectedVersion) {
				data.Result = true
			} else {
				data.Result = false
			}
			json.NewEncoder(w).Encode(data)
		} else {
			json.NewEncoder(w).Encode(Data{ID: id, Agent: "NA"})
		}
	})

	log.Fatal(http.ListenAndServe(":3000", router))
}