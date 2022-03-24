package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
)

//this function is called from when you call the main function to render an .html page from it's templage page
func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

//when the user click's on the button in the template .hml page it calls the fetch command and passes in the argument
//fetch("/play?c=" + x)   (the fetch command is used for calling a URL)
//from the the fetch /play is called, it calls the module rps.go (myapp/rps) and that call analyzes the argument
//passed and returns the JSON result, the results are written to the .html page
func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

//below
//this is our main function: It creates an endpoint /play , /, with each url, it includes the method that is called
//this method is mapped to the /path, {function called when fetched/called}
//when you navigate to it
func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8085")
	//this starts the web server and starts it on PORT 8085 (you can use any port you like)
	http.ListenAndServe(":8085", nil)
}

//this function is used to render your html template page
func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
