package main

// https://tutorialedge.net/golang/creating-restful-api-with-golang/
// https://golang.org/doc/tutorial/web-service-gin

// JPMens November 2021 for Ansible training

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	// yes, password is hard-coded
	// yes, I know
	// yes, this is for toying with the Ansible uri module only
	// no, I will never use this in production
	username = "cow"
	password = "milk"
)

type Cow struct {
	Id     int    `json:"id"`
	Breed  string `json:"breed"`
	Origin string `json:"origin"`
}

type Cowrhyme struct {
	Url      string `json:"url"`
	Rhyme    string `json:"rhyme"`
	Instance string `json:"instance"`
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

	js, err := json.Marshal(cows[n])
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	fmt.Fprintf(w, "\n")
}

// copied from https://stackoverflow.com/questions/21936332/
// BasicAuth wraps a handler requiring HTTP basic auth for it using the given
// username and password and the specified realm, which shouldn't contain quotes.
//
// Most web browser display a dialog with something like:
//
//    The website says: "<realm>"
//
// Which is really stupid so you may want to set the realm to a message rather than
// an actual realm.

func BasicAuth(handler http.HandlerFunc, username, password, realm string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}

		handler(w, r)
	}
}

func secret(w http.ResponseWriter, r *http.Request) {
	myhostname, _ := os.Hostname()
	var cowrhyme = Cowrhyme{
		Url:      "https://en.wikipedia.org/wiki/Hey_Diddle_Diddle",
		Rhyme:    "Hey diddle diddle, The cat and the fiddle, The cow jumped over the moon",
		Instance: myhostname,
	}

	js, _ := json.Marshal(cowrhyme)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	fmt.Fprintf(w, "\n")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/cow", cow)
	http.HandleFunc("/secret",
		BasicAuth(secret, username, password, "Moo are you?"))

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
