package models

// Status to return level (RED/AMBER/GREEN)
type Status struct {
	Level       string `json:"level"`
	Description string `json:"description"`
}

//var status Status
