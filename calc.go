package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"

	"git.sr.ht/~sbinet/gg"
)

var G = 9.8
var time_step = 0.02
var objects = make([]*object, 0)

type data struct {
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
	mass           float64
	color          string
	pending        bool
}

func (obj *object) rndColor() {
	obj.color = randomHex(3)
}

func (obj object) Initialize() *object {
	if obj.size == 0 {
		obj.size = 16
	}
	if obj.mass == 0 {
		obj.size = 10
	}
	if obj.enabled == nil {
		obj.enabled = new(bool)
		*obj.enabled = true
	}
	if obj.color == "" {
		obj.color = "#ffff00"
	}
	if obj.color == "rand" {
		obj.rndColor()

	}
	return &obj
}

// https://en.wikipedia.org/wiki/Elastic_collision#One-dimensional_Newtonian
func colide(obj1 *object, obj2 *object) {
	if obj1.mass == obj2.mass {
		obj1.x_speed, obj2.x_speed = obj2.x_speed, obj1.x_speed
		obj1.y_speed, obj2.y_speed = obj2.y_speed, obj1.y_speed
		return
	}
	m_sum := obj1.mass + obj2.mass
	obj1.x_speed = ((obj1.mass-obj2.mass)/(m_sum))*obj1.x_speed + ((2*obj2.mass)/(m_sum))*obj2.x_speed
	obj2.x_speed = ((2*obj1.mass)/(m_sum))*obj1.x_speed + ((obj2.mass-obj1.mass)/(m_sum))*obj2.x_speed

	obj1.y_speed = ((obj1.mass-obj2.mass)/(m_sum))*obj1.y_speed + ((2*obj2.mass)/(m_sum))*obj2.y_speed
	obj2.y_speed = ((2*obj1.mass)/(m_sum))*obj1.y_speed + ((obj2.mass-obj1.mass)/(m_sum))*obj2.y_speed
}

func checkCollisions() bool {
	// var x_move int
	col := false
	for _, obj1 := range objects {
		if obj1.pending {
			continue
		}
		for _, obj2 := range objects {
			if obj1 == obj2 {
				continue
			}
			if obj2.pending {
				continue
			}
			var x_move, y_move float64
			if obj1.x+obj1.size >= obj2.x && obj1.x <= obj2.x && obj1.y+obj1.size >= obj2.y && obj1.y <= obj2.y {
				x_move = -(obj1.x - obj2.x)
				y_move = -(obj1.y - obj2.y)
			}
			if x_move != 0 || y_move != 0 {
				if math.Abs(x_move) < math.Abs(y_move) {
					obj1.x += x_move / 2
					obj2.x += x_move / 2
				} else if math.Abs(y_move) < math.Abs(x_move) {
					obj1.y += y_move / 2
					obj2.y += y_move / 2
				} else {
					obj1.x += x_move / 2
					obj2.x += x_move / 2
					obj1.y += y_move / 2
					obj2.y += y_move / 2

				}
				fmt.Println(obj1)
				colide(obj1, obj2)
				obj1.pending = true
				obj2.pending = true
				col = true
			}
		}
	}
	return col

}
func (obj *object) stop() {
	obj.x_speed = 0
	obj.y_speed = 0
	obj.x_acceleration = 0
	obj.y_acceleration = 0
}

func (obj *object) bounce_y() {
	obj.y_speed = -obj.y_speed
}
func (obj *object) bounce_x() {
	obj.x_speed = -obj.x_speed
}

func (obj *object) update() {

	obj.x = obj.x + obj.x_speed*time_step + (obj.x_acceleration*math.Pow(time_step, 2))/2
	obj.y = obj.y + obj.y_speed*time_step + (obj.y_acceleration*math.Pow(time_step, 2))/2
	obj.x_speed = obj.x_speed + obj.x_acceleration*time_step
	obj.y_speed = obj.y_speed + obj.y_acceleration*time_step

	if !inRange(int(obj.x)+int(obj.size), 0, fb_x) {
		obj.bounce_x()
	}
	if !inRange(int(obj.y)+int(obj.size), 0, fb_y) {
		obj.bounce_y()
	}
	obj.pending = false

}

func (obj object) draw(frame *gg.Context) { //draw an object
	fb.SetHexColor(obj.color)
	for c_x := int(obj.x); c_x <= int(obj.x)+int(obj.size); c_x++ {
		for c_y := int(obj.y); c_y <= int(obj.y)+int(obj.size); c_y++ {
			if !insideScreen(c_x, c_y) {
				return
			}
			frame.SetPixel(c_x, c_y)

		}
	}
}

func step() { //self explanatory
	fb.SetHexColor("#023e49")
	fb.Clear()
	for _, obj := range objects {
		obj.update()

		if !insideScreen(int(obj.x)+int(obj.size), int(obj.y)+int(obj.size)) {
			// fmt.Println(obj)
			*obj.enabled = false
		} else {
			*obj.enabled = true
		}
		if *obj.enabled {

			obj.draw(fb)
		}
	}
	for checkCollisions() {
	}
}

func randomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
