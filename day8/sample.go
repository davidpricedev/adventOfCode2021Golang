package day8

import (
	"fmt"
	"strings"
)

type Sample struct {
	Patterns    []string
	Outputs     [4]string
	OrigOutputs [4]string
}

func NewSampleFromString(line string) *Sample {
	lineParts := strings.Split(line, " | ")
	patterns := strings.Fields(lineParts[0])
	var outputs [4]string
	copy(outputs[:], strings.Fields(lineParts[1]))
	return &Sample{patterns, outputs, outputs}
}

func (sample Sample) Inspect() {
	fmt.Println("Sample (", sample.Patterns, ") - ", sample.Outputs)
}

func (sample *Sample) ResolveSimple() {
	for i, output := range sample.OrigOutputs {
		if len(output) == 2 {
			sample.Outputs[i] = "1"
		} else if len(output) == 3 {
			sample.Outputs[i] = "7"
		} else if len(output) == 4 {
			sample.Outputs[i] = "4"
		} else if len(output) == 7 {
			sample.Outputs[i] = "8"
		}
	}
}

func PrintSamples(samples []*Sample) {
	for _, sample := range samples {
		sample.Inspect()
	}
}
