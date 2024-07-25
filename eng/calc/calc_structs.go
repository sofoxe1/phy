package calc

import (
	"eng/util"
	"math/rand"
)

type Object struct{
	id int
	Size           float64
	X              float64
	Y              float64
	Xs        float64  //x_speed
	Ys        float64
	Xa float64			//x_acceleration
	Ya float64
	render        *bool
	Mass           float64
	Color          string
}
func (obj Object) Initialize() *Object {
	obj.id=rand.Int()
	if obj.Size == 0 {
		obj.Size = 16
	}
	if obj.Mass == 0 {
		obj.Mass = 1
	}
	if obj.render == nil {
		obj.render = new(bool)
		*obj.render = true
	}
	if obj.Color == "" {
		obj.Color = "#ffff00"
	}
	if obj.Color == "rand" {
		obj.Color=util.RndColor()

	}
	return &obj
}