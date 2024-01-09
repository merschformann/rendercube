package rendering

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed template.html
var template string

// TargetLength is the length to which the solution is resized.
const TargetLength = 1000

// Convert creates a HTML file from the given input.
func Convert(input Input) (string, error) {
	// Collect some information about the solution
	pieces, containers := 0, 0
	for i := range input.Containers {
		containers++
		pieces += len(input.Containers[i].Assignments)
	}

	// Write output
	marshalled, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	output := strings.Replace(template, "{{solution}}", string(marshalled), 1)
	output = strings.Replace(output, "{{pieces}}", fmt.Sprintf("%d", pieces), 1)
	output = strings.Replace(output, "{{containers}}", fmt.Sprintf("%d", containers), 1)
	return output, nil
}
