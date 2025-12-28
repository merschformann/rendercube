package rendering

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

//go:embed template.html
var template string

// TargetLength is the length to which the solution is resized.
const TargetLength = 1000

// Convert creates a HTML file from the given input.
func Convert(input Input) (string, error) {
	// Collect some information about the solution
	pieces, containers, offload := 0, 0, len(input.Offload)
	for i := range input.Containers {
		containers++
		pieces += len(input.Containers[i].Assignments)
	}

	// Sort containers by utilization
	sort.Slice(input.Containers, func(i, j int) bool {
		capacityI := input.Containers[i].Length * input.Containers[i].Width * input.Containers[i].Height
		capacityJ := input.Containers[j].Length * input.Containers[j].Width * input.Containers[j].Height
		usedI := 0.0
		usedJ := 0.0
		for _, assignment := range input.Containers[i].Assignments {
			for _, cube := range assignment.Cubes {
				usedI += cube.Length * cube.Width * cube.Height
			}
		}
		for _, assignment := range input.Containers[j].Assignments {
			for _, cube := range assignment.Cubes {
				usedJ += cube.Length * cube.Width * cube.Height
			}
		}
		return usedI/capacityI > usedJ/capacityJ
	})

	// Sort pieces inside each container by position (z, y, x)
	for i := range input.Containers {
		sort.Slice(input.Containers[i].Assignments, func(a, b int) bool {
			posA := input.Containers[i].Assignments[a].Position
			posB := input.Containers[i].Assignments[b].Position
			if posA.Z != posB.Z {
				return posA.Z < posB.Z
			}
			if posA.Y != posB.Y {
				return posA.Y < posB.Y
			}
			return posA.X < posB.X
		})
	}

	// Write output
	marshalled, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	output := strings.Replace(template, "{{solution}}", string(marshalled), 1)
	output = strings.Replace(output, "{{pieces}}", fmt.Sprintf("%d", pieces), 1)
	output = strings.Replace(output, "{{containers}}", fmt.Sprintf("%d", containers), 1)
	output = strings.Replace(output, "{{offload}}", fmt.Sprintf("%d", offload), 1)
	return output, nil
}
