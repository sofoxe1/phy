package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func render(){
	step()
	fb.SavePNG("dynamic/img.png")
}

func lock(w http.ResponseWriter, r *http.Request) {
	render()
	w.WriteHeader(http.StatusOK)
}
func handleRequest(w http.ResponseWriter, r *http.Request) {
	render()
	t, err := template.ParseFiles("static/main.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, data{})
}

func main() {
	objects = append(objects, object{x: 120, y: 0, y_acceleration: G}.Initialize())
	fs := http.FileServer(http.Dir("./dynamic"))
	http.Handle("/dynamic/", http.StripPrefix("/dynamic/", fs))
	fs2 := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs2))

	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/lock", lock)
	fmt.Println("http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)

	
}
