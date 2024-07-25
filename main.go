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
		objects = append(objects, calc.Object{X: 100, Y: 133,Color: "rand",Xa:2.2,Ya:8.4}.Initialize())
		objects = append(objects, calc.Object{X: 120, Y: 133,Color: "rand"}.Initialize())
		go r.Render(objects)
		var tick_rate float64 =60*8
		var speed float64 = 2
		for {
			time.Sleep(time.Second/time.Duration(tick_rate))
			calc.Step(speed/tick_rate,objects,r.Width,r.Height)
		} 
		
	
 }
