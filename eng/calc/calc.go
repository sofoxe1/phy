package calc

import (
	"eng/util"
	"fmt"
	"math"
)

// https://en.wikipedia.org/wiki/Elastic_collision#One-dimensional_Newtonian
func colide(obj1 *Object, obj2 *Object) {
	if obj1.Mass == obj2.Mass {
		obj1.Xs, obj2.Xs = obj2.Xs, obj1.Xs
		obj1.Ys, obj2.Ys = obj2.Ys, obj1.Ys
		return
	}
	m_sum := obj1.Mass + obj2.Mass
	obj1.Xs = ((obj1.Mass-obj2.Mass)/(m_sum))*obj1.Xs + ((2*obj2.Mass)/(m_sum))*obj2.Xs
	obj2.Xs = ((2*obj1.Mass)/(m_sum))*obj1.Xs + ((obj2.Mass-obj1.Mass)/(m_sum))*obj2.Xs

	obj1.Ys = ((obj1.Mass-obj2.Mass)/(m_sum))*obj1.Ys + ((2*obj2.Mass)/(m_sum))*obj2.Ys
	obj2.Ys = ((2*obj1.Mass)/(m_sum))*obj1.Ys + ((obj2.Mass-obj1.Mass)/(m_sum))*obj2.Ys
}

func checkCollisions(obj1 Object, objects []*Object) (*Object) {
	for _, obj2 := range objects {
		if obj1 == *obj2 {
			continue
		}
		if obj1.X+obj1.Size>obj2.X && obj1.X<obj2.X+obj2.Size && obj1.Y>obj2.Y-obj2.Size && obj1.Y<obj2.Y+obj2.Size {
			return obj2
		}

	}

	return nil

}
func Step(time_step float64,objects []*Object,fbx int, fby int){
	for _,obj:=range objects{
		obj.update(time_step,fbx,fby,objects)
	}
}
func (obj *Object) stop() {
	obj.Xs = 0
	obj.Ys = 0
	obj.Xa = 0
	obj.Ya = 0
}

func (obj *Object) bounce() {
	obj.Ys = -obj.Ys
	obj.Xs = -obj.Xs
}

//https://en.wikipedia.org/wiki/Equations_of_motion#Constant_translational_acceleration_in_a_straight_line
func (obj *Object) update(time_step float64,fbx int, fby int, objects []*Object) {
	obj_c:=*obj
	
	obj.Xs = obj.Xs + obj.Xa*time_step
	obj.Ys = obj.Ys + obj.Ya*time_step
	obj.X = obj.X+((obj_c.Xs+obj.Xs)*time_step)/2
	obj.Y = obj.Y+((obj_c.Ys+obj.Ys)*time_step)/2
	if obj.X != obj_c.X || obj.Y!= obj_c.Y{
		if obj2:= checkCollisions(*obj, objects); obj2!=nil{
			obj.Color=util.RndColor()
			fmt.Println(obj	)
		x_travel:=obj.X-obj_c.X
		y_travel:=obj.Y-obj_c.Y
		path:= math.Sqrt(math.Pow(x_travel,2)+math.Pow(y_travel,2))
		x_ratio:=path/x_travel
		y_ratio:=path/y_travel
		var dx, dy float64
		if x_travel<0{
			 dx=-obj.X+obj2.X+obj2.Size
		}else{
			dx=-obj.X+obj2.X-obj.Size
		} 
		if y_travel<0{
			dy=-obj.Y+obj2.Y+obj2.Size
		}else{
			dy=-obj.Y+obj2.Y-obj2.Size
		}
		if math.Abs(dx)<math.Abs(dy){
			fmt.Println("a")
			obj.X+=dx
			obj.Y+=(dx*x_ratio)/y_ratio
		} else{
			obj.X+=(dy*y_ratio)/x_ratio
			obj.Y+=dy
		}
		obj.stop()
		if x_travel>0{
			tx:= 2*x_travel/(obj_c.Xs+obj.Xs)
			ty:= 2*y_travel/(obj_c.Ys+obj.Ys)
			if ty!=tx{
				panic("math is not mathing")
			}

		}

	}
	}
	if !util.InsideScreen (int(obj.X)+int(obj.Size), int(obj.Y)+int(obj.Size),fbx,fby) || !util.InsideScreen(int(obj.X),int(obj.Y),fbx,fby){
		obj.bounce()
	}

}



