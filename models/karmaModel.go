package models

type Karma struct {
	UserID  string `json:"userid"`
	GuildID string `json:"guildid"`
	Value   uint   `json:"value"`
}

func (karma *Karma) SetKarma(userID, guildID string, value uint) {
	karma.UserID = userID
	karma.GuildID = guildID
	karma.Value = value
}

func (karma *Karma) AddKarma(amount uint) {
	karma.Value += amount
}
