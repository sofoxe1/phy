package util

import (
	"encoding/hex"
	"math/rand"
)

func inRange(n int, min int, max int) bool{
	if n<min{
		return false
	}
	if n>max{
		return false
	}
	return true
}
func InsideScreen(x int, y int,size_x int,size_y int) bool{
	if !inRange(x,0,size_x) || !inRange(y,0,size_y){
		return false
	}
	return true
}
func RndColor() string{
	bytes := make([]byte, 3)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}