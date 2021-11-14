package main

// https://tutorialedge.net/golang/creating-restful-api-with-golang/
// https://golang.org/doc/tutorial/web-service-gin

// JPMens November 2021 for Ansible training

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Cow struct {
	Id     int    `json:"id"`
	Breed  string `json:"breed"`
	Origin string `json:"origin"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Try the /cow endpoint\n")
}

func cow(w http.ResponseWriter, r *http.Request) {
	// https://en.wikipedia.org/wiki/List_of_dairy_cattle_breeds
	var cows = []Cow{
		{Breed: "Abondance", Origin: "France"},
		{Breed: "Alderney", Origin: "England"},
		{Breed: "Australian Braford", Origin: "Australia"},
		{Breed: "Australian Friesian Sahiwal", Origin: "Australia"},
		{Breed: "Australian Lowline", Origin: "Australia"},
		{Breed: "Australian Milking Zebu", Origin: "Australia"},
		{Breed: "Ayrshire", Origin: "Scotland"},
		{Breed: "Belgian Red", Origin: "Belgium"},
		{Breed: "Bianca Modenese", Origin: "Italy"},
		{Breed: "Brown Swiss", Origin: "Switzerland"},
		{Breed: "Burlina", Origin: "Italy"},
		{Breed: "Buša", Origin: "Balkan"},
		{Breed: "Canadienne", Origin: "Canada"},
		{Breed: "Carora", Origin: "Venezuela"},
		{Breed: "Dairy Shorthorn", Origin: "England"},
		{Breed: "Danish Jersey", Origin: "Denmark"},
		{Breed: "Danish Red", Origin: "Denmark"},
		{Breed: "Dexter", Origin: "Ireland"},
		{Breed: "Estonian Red", Origin: "Estonia"},
		{Breed: "Simmental", Origin: "Austria"},
		{Breed: "French Simmental", Origin: "France"},
		{Breed: "German Black Pied Dairy", Origin: "Germany"},
		{Breed: "Girolando", Origin: "Brazil"},
		{Breed: "Gloucester", Origin: "England"},
		{Breed: "Guernsey", Origin: "Guernsey"},
		{Breed: "Harzer Rotvieh", Origin: "Germany"},
		{Breed: "Hays Converter", Origin: "Canada"},
		{Breed: "Hérens", Origin: "Switzerland"},
		{Breed: "Holstein-Friesian", Origin: "Netherlands"},
		{Breed: "Illawarra", Origin: "Australia"},
		{Breed: "Irish Moiled", Origin: "Ireland"},
		{Breed: "Jamaica Hope", Origin: "Jamaica"},
		{Breed: "Jersey", Origin: "Jersey"},
		{Breed: "Aandu", Origin: "India"},
		{Breed: "Lakenvelder", Origin: "Netherlands"},
		{Breed: "Lineback", Origin: "United States"},
		{Breed: "Meuse-Rhine-Issel", Origin: "Germany"},
		{Breed: "Milking Devon", Origin: "United States"},
		{Breed: "Montbéliarde", Origin: "France"},
		{Breed: "Normande", Origin: "France"},
		{Breed: "Norwegian Red", Origin: "Norway"},
		{Breed: "Randall", Origin: "United States"},
		{Breed: "Pie Rouge des Plaines", Origin: "France"},
		{Breed: "Pinzgauer", Origin: "Austria"},
		{Breed: "Red Chittagong", Origin: "Bangladesh"},
		{Breed: "Red Poll", Origin: "England"},
		{Breed: "Red Sindhi", Origin: "Pakistan"},
		{Breed: "Sahiwal", Origin: "Pakistan"},
		{Breed: "Square Meater", Origin: "Australia"},
		{Breed: "Sussex", Origin: "England"},
		{Breed: "Tyrol Grey", Origin: "Austria"},
		{Breed: "Vorderwald", Origin: "Germany"},
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	n := r1.Intn(len(cows))
	cows[n].Id = n

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cows[n])
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/cow", cow)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
