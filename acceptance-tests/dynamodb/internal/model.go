package dynamo

// GameScore ...
type GameScore struct {
	UserID           string `json:"UserId"`
	GameTitle        string `json:"GameTitle"`
	TopScore         int    `json:"TopScore"`
	TopScoreDateTime string `json:"TopScoreDateTime"`
	Wins             int    `json:"Wins"`
	Losses           int    `json:"Losses"`
}
