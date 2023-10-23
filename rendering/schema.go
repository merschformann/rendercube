package rendering

type Input struct {
	Containers []Container `json:"containers"`
}

type Container struct {
	ID          int          `json:"id"`
	Assignments []Assignment `json:"assignments"`
	Length      float64      `json:"length"`
	Width       float64      `json:"width"`
	Height      float64      `json:"height"`
}

type Assignment struct {
	Piece    int      `json:"piece"`
	Position Position `json:"position"`
	Cubes    []Cube   `json:"cubes"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
}

type Cube struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Z      float64 `json:"z"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}
