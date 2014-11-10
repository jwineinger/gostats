package main

import (
	"fmt"
	"time"
	"math/rand"
)


type dataset struct {
	length int
	samples []float64
}

func average(data *dataset) float64 {
	total := 0.0
	for _, val := range data.samples {
		total += val
	}
	return total / float64(data.length)
}


func new_sample(data *dataset) {
	data.length += 1
	data.samples = append(data.samples, rand.Float64() * float64(rand.Int()))
}


func main() {
	data := dataset{0, []float64{}}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10; i++ {
		new_sample(&data)
	}
	fmt.Println(data.samples)
	fmt.Println(average(&data))
}
