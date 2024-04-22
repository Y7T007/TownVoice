package models

type Rating struct {
	EntityID string             `json:"entity_id"`
	UserID   string             `json:"user_id"`
	Scores   map[string]float64 `json:"scores"`
}
