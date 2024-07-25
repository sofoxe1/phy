package calc

import (
	"eng/util"
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

// func checkCollisions() bool {
// 	// var x_move int
// 	col := false
// 	for _, obj1 := range objects {

// 		for _, obj2 := range objects {
// 			if obj1 == obj2 {
// 				continue
// 			}
// 			if slices.Contains(pending_collisions,obj1.id+obj2.id){continue}
// 			if obj1.x+obj1.size>=obj2.x && obj1.x<=obj2.x+obj2.size && obj1.y+obj1.size>=obj2.y && obj1.y<=obj2.y+obj2.size {
// 				colide(obj1, obj2)
// 				pending_collisions=append(pending_collisions,obj1.id+obj2.id)
// 				col = true
// 			}

// 		}
// 	}
// 	return col

// }
func Step(time_step float64,objects []*Object,fbx int, fby int){
	for _,obj:=range objects{
		obj.update(time_step,fbx,fby)
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


func (obj *Object) update(time_step float64,fbx int, fby int) {

	obj.X = obj.X + obj.Xs*time_step + (obj.Xa*math.Pow(time_step, 2))/2
	obj.Y = obj.Y + obj.Ys*time_step + (obj.Ya*math.Pow(time_step, 2))/2
	obj.Xs = obj.Xs + obj.Xa*time_step
	obj.Ys = obj.Ys + obj.Ya*time_step

	if !util.InsideScreen (int(obj.X)+int(obj.Size), int(obj.Y)+int(obj.Size),fbx,fby) || !util.InsideScreen(int(obj.X),int(obj.Y),fbx,fby){
		obj.bounce()
	}

}



