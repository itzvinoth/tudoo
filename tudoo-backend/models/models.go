package models

// Board scheme of the board table
type Board struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
}