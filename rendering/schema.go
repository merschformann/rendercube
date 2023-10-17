package rendering

type Input struct {
}

type Instance struct {
	Containers []Container `json:"containers"`
}

type Container struct {
	ID          string       `json:"id"`
	Assignments []Assignment `json:"assignments"`
	Length      float64      `json:"length"`
	Width       float64      `json:"width"`
	Height      float64      `json:"height"`
}

type Assignment struct {
	Piece string `json:"piece"`
	Cubes []Cube `json:"cubes"`
}

type Cube struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Z      float64 `json:"z"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}
