package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Home"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	type User struct {
		Id   int
		Name string
		Age  int
	}

	data := []User{
		{Id: 1, Name: "Vlad", Age: 21},
		{Id: 2, Name: "Roma", Age: 20},
		{Id: 3, Name: "Kate", Age: 20},
	}

	var targetUser User

	for _, user := range data {
		if user.Id == id {
			targetUser = user
		}
	}

	if targetUser == (User{}) {
		http.NotFound(w, r)
		return
	} else {
		jsonResponse, err := json.Marshal(targetUser)

		if err != nil {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Snippet create"))
}
