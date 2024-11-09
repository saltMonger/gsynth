package osc

import "math"

const TAU = math.Pi * 2

type Oscillator interface {
	Generate(freq int, duration float64) []int16
}

func Quantize(scale float64) int16 {
	return int16(FLOAT_TO_INT16_SCALE * scale)
}
