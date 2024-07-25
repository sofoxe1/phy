module ga

go 1.21.12

// require eng v0.0.0 // indirect

require eng/render v0.0.0

require eng/calc v0.0.0

replace eng/render v0.0.0 => ./eng

replace eng/calc v0.0.0 => ./eng/calc

// replace eng v0.0.0 => ./eng
