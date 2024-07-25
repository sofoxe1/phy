package main

import (
	"eng/calc"
	"eng/render"
	"log"
	"time"
)

 func main(){
	err,r:=render.Renderer{Backend:"web",Img_path:"/tmp/im.png",Width: 1920, Height: 1080}.Initialize()
	if err!=nil{
		log.Fatal(err)
	}
		var objects []*calc.Object
		objects = append(objects, calc.Object{X: 100, Y: 80,Xs:2.2,Ys:2.2,  Color: "rand"}.Initialize())
		objects = append(objects, calc.Object{X: 142, Y: 124,Xs:-2.2,Ys:-2.2,  Color: "rand"}.Initialize())
		objects = append(objects, calc.Object{X: 123, Y: 110,Color: "rand"}.Initialize())
		go r.Render(objects)
		var tick_rate float64 =60*8
		var speed float64 = 2
		for {
			time.Sleep(time.Second/time.Duration(tick_rate))
			calc.Step(speed/tick_rate,objects,r.Width,r.Height)
		} 
		
	
 }
