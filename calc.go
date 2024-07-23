package main

import (
	"math"

	"git.sr.ht/~sbinet/gg"
)


var G = 9.8
var time_step = 0.01
var objects = make([]*object, 0)
var fb = gg.NewContext(512, 512)

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
	obj.y = obj.y + obj.y_speed*time_step + (obj.y_acceleration*math.Pow(time_step, 2))/2
	obj.x_speed = obj.x_speed + obj.x_acceleration*time_step
	obj.y_speed = obj.y_speed + obj.y_acceleration*time_step
}

func (obj object) draw(frame *gg.Context) {//draw an object
	for c_x := int(obj.x); c_x <= int(obj.x)+int(obj.size); c_x++ {
		for c_y := int(obj.y); c_y <= int(obj.y)+int(obj.size); c_y++ {
			if !inRange(c_y,0,512) || !inRange(c_x,0,512){
				return
			}
			frame.SetPixel(c_x,c_y)

		}
	}
}



func step() { //self explanatory 
	for _, obj := range objects {
		fb.SetHexColor("#023e49")
		fb.Clear()
		fb.SetHexColor("#ffff00")
		obj.update()
		if !inRange(int(obj.y),0,512) || !inRange(int(obj.x),0,512){
			*obj.enabled=false
		}
		if *obj.enabled {
			obj.draw(fb)
		}
	}
}
