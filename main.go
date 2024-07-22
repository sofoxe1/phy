package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
)
type data struct {
	Kitten string
}
type User struct {
    Name       string
    Occupation string
}

func get_data() ([]byte,error){
	c:= [512][512]uint8{}
	for i := 0; i < len(c); i ++{
		for ii:= 0; ii<len(c[i]); ii++{
			c[i][ii]=uint8(rand.Uint32()) //unit32 -> unit 8, u can do that, what????
		}
	}
	jsonData, err := json.Marshal(c)
	return jsonData,err
}
func render() error{
	j,err:=get_data()
	// fmt.Println(j)
	err2 := os.WriteFile("/tmp/goipc.json",j,0644)
	if err !=nil{
		log.Fatal(err)
	}else if err2 !=nil{
		log.Fatal(err2)
	}
	os.Remove("img.png")
	cmd := exec.Command("python3","draw.py")
	cmd.Dir="./draw"
	
	if out, err := cmd.CombinedOutput(); err!=nil{
		return errors.New(string(out)+err.Error())
	}
	return nil
}

func lock(w http.ResponseWriter, r *http.Request) {
	err := render()

	if err!= nil{
	http.Error(w,strings.Replace(
		err.Error() ,"\n","<br>",-1)+err.Error(),500)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func handleRequest(w http.ResponseWriter, r *http.Request) {
	render()
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
	http.HandleFunc("/lock",lock)
    http.ListenAndServe("127.0.0.1:8080", nil)
}