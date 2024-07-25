package render

import (
	"errors"
	"slices"
	"strings"
)

type json_screen_size struct{
	Width  int
	Height int
}

type Renderer struct{
	Backend string 
	Img_path string 
	Width int
	Height int
	background_color string
	Fps int
}
func (r *Renderer) Initialize() error{
	backends:= []string{"png","jpg","web"}
	if r.background_color==""{
		r.background_color="#023e49"
	}
	if r.Fps==0{
		r.Fps=60
	}
	if r.Width == 0 || r.Height == 0{
		return errors.New("screen size not set")
	}

	if !slices.Contains(backends,r.Backend){
		return errors.New("possible backends are:"+":"+strings.Join(backends, ","))
	}
	if (r.Backend=="png"||r.Backend =="jpg") && r.Img_path == "" {
		return errors.New("file path not set")
	}

	return nil
}