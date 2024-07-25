package main

import (
	"eng/calc"
	"eng/render"
	"log"
	"math/rand"
	"time"
)

 func main(){
	err,r:=render.Renderer{Backend:"web",Img_path:"/tmp/im.png",Width: 1920, Height: 1080}.Initialize()
	if err!=nil{
		log.Fatal(err)
	}
		rnd:= func() float64{
			return rand.Float64()*2
		}
		var objects []*calc.Object
		objects = append(objects, calc.Object{X: 100, Y: 80,Xa:2.2,Ya:2.2,  Color: "rand"}.Initialize())
		// objects = append(objects, calc.Object{X: 142, Y: 124,Xa:2.2,Ya:-2.2,  Color: "rand"}.Initialize())
		// objects = append(objects, calc.Object{X: 123, Y: 110,Color: "rand"}.Initialize())
		for i:=0; i<40; i++{
			objects = append(objects, calc.Object{X: float64(i*20), Y: float64(i*20),Xa:rnd(),Ya:rnd(), Color: "rand"}.Initialize())
		}
		go r.Render(objects)
		var tick_rate float64 =60*400
		var speed float64 = 500
		for {
			time.Sleep(time.Second/time.Duration(tick_rate))
			calc.Step(speed/tick_rate,objects,r.Width,r.Height)
		} 
		
	
 }
