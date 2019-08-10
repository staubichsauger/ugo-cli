package game

type Join struct {
	Name string `json:"name"`
}

type Card struct {
	Color string `json:"color"`
	Value string `json:"value"`
}

func (c *Card) String() string {
	return c.Value + "-" + c.Color
}

type Player struct {
	Name string `json:"player_name"`
	CardCount int `json:"card_count"`
}

type Turn struct {
	PlayCard *Card `json:"play_card"`
}

type Id struct {
	PlayerId string `json:"player_id"`
	PlayerName string `json:"player_name"`
}