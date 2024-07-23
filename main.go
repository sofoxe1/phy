package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"git.sr.ht/~sbinet/gg"
)

var fb = gg.NewContext(512, 512)
var fb_y = 512
var fb_x = 512
type _fb struct {
	Width int
	Height int
  }
func render(){
	step()
	fb.SavePNG("dynamic/img.png")
}

func lock(w http.ResponseWriter, r *http.Request) {
	render()
	w.WriteHeader(http.StatusOK)
}

func config(w http.ResponseWriter, r *http.Request) {
	var data _fb
	json.NewDecoder(r.Body).Decode(&data)
	w.WriteHeader(http.StatusOK)
	fb_x=data.Width-25
	fb_y=data.Height-25 //for some reason actual size is slightly smaller
	if fb_x<100 || fb_y<100{
		panic("screen too small")
	}
	fb=gg.NewContext(fb_x, fb_y)
	fmt.Println(fb_x)



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
	http.HandleFunc("/config", config)
	fmt.Println("http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)

	
}
