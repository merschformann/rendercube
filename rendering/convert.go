package rendering

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//go:embed template.html
var template string

func Convert(input Input, outputFile string) error {
	// Collect some information about the solution
	pieces, containers := 0, 0
	for i := range input.Containers {
		containers++
		pieces += len(input.Containers[i].Assignments)
	}

	// Downsize lengths
	factor := 1.0
	for i := range input.Containers {
		input.Containers[i].Length /= factor
		input.Containers[i].Width /= factor
		input.Containers[i].Height /= factor
		for j := range input.Containers[i].Assignments {
			input.Containers[i].Assignments[j].Position.X /= factor
			input.Containers[i].Assignments[j].Position.Y /= factor
			input.Containers[i].Assignments[j].Position.Z /= factor
			for k := range input.Containers[i].Assignments[j].Cubes {
				input.Containers[i].Assignments[j].Cubes[k].X /= factor
				input.Containers[i].Assignments[j].Cubes[k].Y /= factor
				input.Containers[i].Assignments[j].Cubes[k].Z /= factor
				input.Containers[i].Assignments[j].Cubes[k].Length /= factor
				input.Containers[i].Assignments[j].Cubes[k].Width /= factor
				input.Containers[i].Assignments[j].Cubes[k].Height /= factor
			}
		}
	}

	// Output
	marshalled, err := json.Marshal(input)
	if err != nil {
		return err
	}
	output := strings.Replace(template, "{{solution}}", string(marshalled), 1)
	output = strings.Replace(output, "{{pieces}}", fmt.Sprintf("%d", pieces), 1)
	output = strings.Replace(output, "{{containers}}", fmt.Sprintf("%d", containers), 1)
	err = os.WriteFile(outputFile, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
