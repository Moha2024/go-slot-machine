package rng

import "math/rand"

type RealGenerator struct {}

func (r *RealGenerator) NumberGenerator(min int, max int) int {
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}
