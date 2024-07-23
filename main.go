package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func get_data() ([]byte, error) {
	step()
	jsonData, err := json.Marshal(fb)
	return jsonData, err
}

func render() error { //generate .png
	j, err := get_data()
	err2 := os.WriteFile("/tmp/goipc.json", j, 0644) //state of the art ipc
	if err != nil {
		log.Fatal(err)
	} else if err2 != nil {
		log.Fatal(err2)
	}
	os.Remove("img.png")
	cmd := exec.Command("python3", "draw.py")
	cmd.Dir = "./draw"

	if out, err := cmd.CombinedOutput(); err != nil {
		return errors.New(string(out) + err.Error())
	}
	return nil
}

func lock(w http.ResponseWriter, r *http.Request) {
	err := render()

	if err != nil {
		http.Error(w, strings.Replace(
			err.Error(), "\n", "<br>", -1)+err.Error(), 500)
		return
	}
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
	objects = append(objects, object{x: 0, y: 400, y_acceleration: G}.Initialize())
	fs := http.FileServer(http.Dir("./dynamic"))
	http.Handle("/dynamic/", http.StripPrefix("/dynamic/", fs))
	fs2 := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs2))

	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/lock", lock)
	fmt.Println("http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)

	
}
