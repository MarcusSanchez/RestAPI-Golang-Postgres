package models

type Team struct {
	ID       uint     `gorm:"primaryKey" json:"id,omitempty"`
	TeamName string   `json:"team_name"`
	Location string   `json:"location"`
	Nickname string   `json:"nickname"`
	Players  []Player `gorm:"foreignKey:TeamID" json:"players,omitempty"`
}

type Player struct {
	ID         uint   `gorm:"primaryKey" json:"id,omitempty"`
	TeamID     uint   `json:"team_id,omitempty"`
	PlayerName string `json:"player_name"`
	Position   string `json:"position"`
}
