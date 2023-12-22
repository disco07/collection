package main

import (
	"collections"
	"log"
)

func main() {
	var arr = []float64{1564, 20, 8, 31, 44, 5008, 6, 7, 8, 9, 1000}
	min := collections.Collect(arr).Max()
	log.Println(min)
}
