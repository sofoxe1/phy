package render

import (
	"encoding/json"
	"eng/calc"
	"eng/util"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"git.sr.ht/~sbinet/gg"
)

func (r *Renderer) Render(objects []*calc.Object) {
	var draw func()

	if r.Backend == "png" || r.Backend == "jpg" || r.Backend == "web" {
		draw = func() {
			fb := gg.NewContext(r.Width, r.Height)

			for _, obj := range objects {
				fb.SetHexColor(obj.Color)
				for c_x := int(obj.X) - int(obj.Size)/2; c_x <= int(obj.X)+int(obj.Size)/2; c_x++ {
					for c_y := int(obj.Y) - int(obj.Size)/2; c_y <= int(obj.Y)+int(obj.Size)/2; c_y++ {
						if util.InsideScreen(c_x, c_y, r.Width, r.Height) {
							continue
						}
						fb.SetPixel(c_x, c_y)

					}
				}
			}
			fb.SetHexColor(r.background_color)
			fb.SavePNG(r.Img_path)
		}
	}
	if r.Backend == "web" {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		http.HandleFunc("/img", func(w http.ResponseWriter, req *http.Request) {
			http.ServeFile(w, req, r.Img_path)
		})
		fs := http.FileServer(http.Dir(filepath.Join(exPath,"./static")))
		http.Handle("/static/", http.StripPrefix("/static/", fs))

		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			http.ServeFile(w, req, filepath.Join(exPath,"./static/main.html"))
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
		fmt.Println("http://127.0.0.1:8080")
		http.ListenAndServe("127.0.0.1:8080", nil)
	}

	for {
		time.Sleep(time.Second / time.Duration(r.Fps))
		draw()
	}

}
