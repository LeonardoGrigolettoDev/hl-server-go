package models

type Device struct {
	ID         int64  `json:"id"`
	MAC        string `json:"mac"`
	Actionable bool   `json:"actionable"`
	Camera     bool   `json:"camera"`
	Drone      bool   `json:"drone"`
}
