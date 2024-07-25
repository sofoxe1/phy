module phy

go 1.21.12

// // require eng v0.0.0
// require eng/cakc v0.0.0
// require eng/render v0.0.0
// require eng/util v0.0.0

// replace eng v0.0.0 => ./eng
replace eng/calc v0.0.0 => ./eng/calc

replace eng/render v0.0.0 => ./eng/render

replace render/static v0.0.0 => ./eng/render/static

replace eng/util v0.0.0 => ./eng/util

require (
	eng/calc v0.0.0
	eng/render v0.0.0
	git.sr.ht/~sbinet/gg v0.5.0 // indirect
	render/static v0.0.0 // indirect
)

require (
	eng/util v0.0.0 // indirect
	github.com/campoy/embedmd v1.0.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/image v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)
