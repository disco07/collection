package main

import (
	"collection"
	"collection/arithmetic"
	"collection/cmp"
	"log"
)

func main() {
	var arr = []float64{1, 2, 2, 2, 2}
	avg := collection.NewVec(arr).Dedup(cmp.Eq[float64]).Avg(arithmetic.Avg[float64])
	log.Println(avg)
}
