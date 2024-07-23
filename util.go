package main


func inRange(n int, min int, max int) bool{
	if n<min{
		return false
	}
	if n>max{
		return false
	}
	return true
}
func insideScreen(x int, y int) bool{
	if !inRange(x,0,fb_x) || !inRange(y,0,fb_y){
		return false
	}
	return true
}