package main

import (
	"collection"
	"log"
)

func main() {
	var arr = []float64{1, 2, 2, 2, 2}
	avg := collection.NewVec(arr).Dedup().Avg()
	log.Println(avg)
}
