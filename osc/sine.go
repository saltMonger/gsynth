package osc

import (
	"gsynth/models"
	"math"
)

const FLOAT_TO_INT16_SCALE = 32768

type Sine struct {
}

func (s Sine) Generate(freq int, duration float64) []int16 {
	phaseInc := (TAU / models.CD_SAMPLE_RATE) * float64(freq)
	phase := float64(0)
	volume := float64(1)
	totalSamples := uint32(duration * models.CD_SAMPLE_RATE)

	//fmt.Println("tau: ", TAU)
	//fmt.Println("phase inc: ", phaseInc)

	samples := []int16{}

	for n := uint32(0); n < totalSamples; n++ {
		unmod := volume * math.Sin(phase)
		sample := Quantize(unmod)
		if n <= 200 {
			// fmt.Println("phase:", phase, "sample: ", sample)
			//fmt.Println(n, " sample:", sample, " unmod: ", unmod, "second: ", float64(n)/float64(models.CD_SAMPLE_RATE))
		}
		samples = append(samples, sample)
		phase += phaseInc
		if phase >= TAU {
			phase -= TAU
		}
	}
	return samples
}
