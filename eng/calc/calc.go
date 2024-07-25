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
		if obj1.X+obj1.Size>=obj2.X && obj1.X<=obj2.X+obj2.Size && obj1.Y>=obj2.Y-obj2.Size && obj1.Y<=obj2.Y+obj2.Size {
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
			obj.stop()
		// obj=obj_c
		x_travel:=obj.X-obj_c.X
		y_travel:=obj.Y-obj_c.Y
		path:= math.Sqrt(math.Pow(x_travel,2)+math.Pow(y_travel,2))
		x_ratio:=path/x_travel
		y_ratio:=path/y_travel
		fmt.Println(x_ratio)
		fmt.Println(y_ratio)
		// obj.X=obj2.X-obj2.Size
		d:=((obj2.X-obj2.Size)-obj.X)/x_ratio
		obj.X+=d
		fmt.Println(d)
		obj.Y+=(d/y_ratio)
		// obj.X=obj2.X+(16*x_ratio)
		fmt.Println(obj.Y)
		// obj.X=obj2.X+(16/x_ratio)
		// panic("a")
		// obj.stop()
		// time.Sleep(time.Second*10000)

		//  ty kurawa org-c = przemeieszczenie, chieć by to równać obj2+half size (krawendz)
		fmt.Println(obj	)
		fmt.Println(obj2)
		if x_travel>0{
			tx:= 2*x_travel/(obj_c.Xs+obj.Xs)
			ty:= 2*y_travel/(obj_c.Ys+obj.Ys)
			fmt.Println(tx)
			fmt.Println(ty)
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



