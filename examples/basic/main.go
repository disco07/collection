package main

import (
	"collections"
	"log"
)

func main() {
	var arr = []float64{1, 2, 2, 2, 2}
	avg := collections.Collect(arr).Dedup().Avg()
	log.Println(avg)
}
