package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
)
type data struct {
	Kitten string
}
type User struct {
    Name       string
    Occupation string
}
func handleRequest(w http.ResponseWriter, r *http.Request) {
	os.Remove("img.png")
	exec.Command("go","run", "./draw/draw.go").Run()
    // buf, err := os.ReadFile("img.png")

    // if err != nil {
        // return
    // }
	// title := r.URL.Path[len("/"):]
	p := data{"meow:3"}
    t, err := template.ParseFiles("static/main.tmpl")
	if err != nil{
		log.Fatal(err)
	}
    t.Execute(w, p)
}



func main() {
	
	http.HandleFunc("/",handleRequest)
	fs := http.FileServer(http.Dir("./dynamic"))
	http.Handle("/dynamic/", http.StripPrefix("/dynamic/", fs)) 
    fmt.Println("http://127.0.0.1:8080")
    http.ListenAndServe("127.0.0.1:8080", nil)
}