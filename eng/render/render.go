package render

import (
	"encoding/json"
	"eng/calc"
	"eng/util"
	"fmt"
	"net/http"
	"os"
	"render/static"
	"time"

	"git.sr.ht/~sbinet/gg"
)

func (r *Renderer) Render(objects []*calc.Object) {
	
	var draw func()
	f,_:=os.OpenFile(r.Img_path,os.O_CREATE,0777)
	f.Sync()
	
	if r.Backend == "png" || r.Backend == "jpg" || r.Backend == "web" {
		draw = func() {
			
			for {
				fb := gg.NewContext(r.Width, r.Height)
				time.Sleep(time.Second / time.Duration(r.Fps))
				fb.SetHexColor(r.background_color)
				fb.Clear()
			

			for _, obj := range objects {
				fb.SetHexColor(obj.Color)
				for c_x := int(obj.X) - int(obj.Size)/2+1; c_x <= int(obj.X)+int(obj.Size)/2; c_x++ {
					for c_y := int(obj.Y) - int(obj.Size)/2+1; c_y <= int(obj.Y)+int(obj.Size)/2; c_y++ {
						if !util.InsideScreen(c_x, c_y, r.Width, r.Height) {
							continue
						}
						fb.SetPixel(c_x, c_y)

					}
				}
			}
			
			// wasted over an hour to realize gg.context.SavePng doesn't call Sync which results in no or garbage data being written...
			f,err:=os.OpenFile(r.Img_path,os.O_WRONLY,0777)
			err= fb.EncodePNG(f)
			f.Sync()
			if err!= nil{
				panic(err)
			}
		}
	}
	}
	if r.Backend == "web" {
		http.HandleFunc("/img.png", func(w http.ResponseWriter, req *http.Request) {
			http.ServeFile(w, req, r.Img_path)
		})
		http.HandleFunc("/static/js.js", func(w http.ResponseWriter, req *http.Request) {
			w.Write(static.Js)
		})

		http.HandleFunc("/main.html", func(w http.ResponseWriter, req *http.Request) {
			w.Write(static.MainHtml)
		})
		http.HandleFunc("/config",
			func(w http.ResponseWriter, req *http.Request) {
				var data json_screen_size
				json.NewDecoder(req.Body).Decode(&data)
				w.WriteHeader(http.StatusOK)
				r.Width = data.Width - 26
				r.Height = data.Height - 26 //for some reason actual size is slightly smaller
				if r.Width < 100 || r.Height < 100 {
					panic("screen too small")
				}
			})
		fmt.Println("http://127.0.0.1:8080/main.html")
		go http.ListenAndServe("127.0.0.1:8080", nil)
	}

	
		draw()

}
