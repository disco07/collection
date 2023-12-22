package main

import (
	"collections"
	"log"
)

func main() {
	var arr = []float64{1, 20}
	avg := collections.Collect(arr).Sum()
	log.Println(avg)
}
