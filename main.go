package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func accountPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the accountPage!")
	fmt.Println("Endpoint Hit: accountPage")
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the aboutPage!")
	fmt.Println("Endpoint Hit: aboutPage")
}
func contactUsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the contactUsPage!")
	fmt.Println("Endpoint Hit: contactUsPage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/accounts", accountPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactUsPage)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequests()
}
