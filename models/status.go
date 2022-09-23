package models

type StatusAirAngin struct {
	Status struct {
		Air   int `json:"air"`
		Angin int `json:"anging"`
	} `json:"status"`
}
