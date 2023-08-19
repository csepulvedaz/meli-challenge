package models

type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Secret struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}
