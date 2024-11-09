package main

import (
	"gsynth/models"
	"gsynth/osc"
)

func main() {
	gen := osc.Sine{}
	output := gen.Generate(880, 1.5)
	models.WriteWaveFile(output, uint32(len(output)), 1, "output.wav")
}
