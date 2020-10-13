package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"html/template"
)

type User struct  {
	Id int 		 `json:"id"`
	Name string	 `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/template", templateHandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
} 

//jsonHandler returns http respone in JSON format.
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := USER{
		Id: 1, 
		Name: "John Doe", 
		Email: "johndoe@gmail.com", 
		Phone: "000099999"
	}
	json.NewEncoder(w).Encode(user)
}

//templateHandler renders a template and returns as http response.
func templateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("template.html")
	fmt.Fprintf(w, "Success response HTML")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}

	user := User{Id: 1, 
		Name: "John Doe", 
		Email: "johndoe@gmail.com", 
		Phone: "000099999"}
	t.Execute(w, user)
}
