package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"math"
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
type object struct{
	size float64
	x float64
	y float64
	x_speed float64
	y_speed float64
	x_acceleration float64
	y_acceleration float64
	enabled *bool
}
var G = 9.8
var time_step=0.22
func (obj *object) update(){
	fmt.Println(obj.y)
obj.x=obj.x+obj.x_speed*time_step+(obj.x_acceleration*math.Pow(time_step,2))/2
obj.y=obj.y-obj.y_speed*time_step+(obj.y_acceleration*math.Pow(time_step,2))/2
obj.x_speed=obj.x_speed+obj.x_acceleration*time_step
obj.y_speed=obj.y_speed+obj.y_acceleration*time_step
}
func (obj object) draw (frame *[512][512]uint8 ) (){
	// fmt.Println(obj.x)
	for  c_x := int(obj.x); c_x<=int(obj.x)+int(obj.size); c_x++{
		for c_y:=int(obj.y); c_y<=int(obj.y)+int(obj.size); c_y++{
			frame[c_y][c_x]=200
			
		}
	} 
}

func step(){
	for _,obj := range objects{
		if *obj.enabled{
		c=[512][512]uint8{}
		obj.update()
		obj.draw(&c)
		}
	}
}
var objects = make([]*object,0)
var c = [512][512]uint8{}
func get_data() ([]byte,error){
	step()
	jsonData, err := json.Marshal(c)
	return jsonData,err
}

func (obj object) deafault() *object{
	fmt.Println(obj.x)
	if obj.size==0{obj.size=16}
	if obj.enabled==nil{
		obj.enabled=new(bool)
		*obj.enabled=true
	}
	return &obj
}
func render() error{
	j,err:=get_data()
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
	objects = append(objects,object{x: 0, y:400,y_acceleration: G}.deafault())
	
	http.HandleFunc("/",handleRequest)
	fs := http.FileServer(http.Dir("./dynamic"))
	http.Handle("/dynamic/", http.StripPrefix("/dynamic/", fs)) 
    fmt.Println("http://127.0.0.1:8080")
	http.HandleFunc("/lock",lock)
    http.ListenAndServe("127.0.0.1:8080", nil)
}