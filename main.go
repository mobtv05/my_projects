package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	pointsMap := map[string]Point{
		"b": {
			Y: 13,
			X: 15,
		}
	}
	otherPointsMap := make(map[int]Point)
	// var oneMorePontsMap map[string]Point
	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	// fmt.Println(pointsMap)
	// fmt.Println(pointsMap["a"])
	otherPointsMap[1] = Point{Y: 1, X: 2}
	// fmt.Println(otherPointsMap)
	// fmt.Println(otherPointsMap[1])



	var oneMorePontsMap map[string]Point
	oneMorePontsMap
}
