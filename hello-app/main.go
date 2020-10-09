/**
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START container_hello_app]
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"html/template"
)
type User struct  {
	Id int 		 `json:"id"`
	Name string	 `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.HandleFunc("/template", templateHandler)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "<h3>Hello, world!</h3>\n")
	fmt.Fprintf(w, "Version: 2.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}

// [END container_hello_app]

func templateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}
	user := User{Id: 1, 
		Name: "John Doe", 
		Email: "johndoe@gmail.com", 
		Phone: "000099999"}
	t.Execute(w,user)
}
