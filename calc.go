package main

import (
	"math"
)


var G = 9.8
var time_step = 0.22
var objects = make([]*object, 0)
var fb = [512][512]uint8{}

type data struct{

}

type object struct {
	size           float64
	x              float64
	y              float64
	x_speed        float64
	y_speed        float64
	x_acceleration float64
	y_acceleration float64
	enabled        *bool
}

func (obj object) Initialize() *object {
	if obj.size == 0 {
		obj.size = 16
	}
	if obj.enabled == nil {
		obj.enabled = new(bool)
		*obj.enabled = true
	}
	return &obj
}

func (obj *object) update() { //do physics
	obj.x = obj.x + obj.x_speed*time_step + (obj.x_acceleration*math.Pow(time_step, 2))/2
	obj.y = obj.y - obj.y_speed*time_step + (obj.y_acceleration*math.Pow(time_step, 2))/2
	obj.x_speed = obj.x_speed + obj.x_acceleration*time_step
	obj.y_speed = obj.y_speed + obj.y_acceleration*time_step
}

func (obj object) draw(frame *[512][512]uint8) {//draw an object
	for c_x := int(obj.x); c_x <= int(obj.x)+int(obj.size); c_x++ {
		for c_y := int(obj.y); c_y <= int(obj.y)+int(obj.size); c_y++ {
			frame[c_y][c_x] = 200

		}
	}
}



func step() { //self explanatory 
	for _, obj := range objects {
		if *obj.enabled {
			fb = [512][512]uint8{}
			obj.update()
			obj.draw(&fb)
		}
	}
}
