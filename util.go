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