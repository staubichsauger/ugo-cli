package game

type Status struct {
	MyTurn        bool     `json:"my_turn"`
	Hand          []Card   `json:"hand"`
	OtherPlayers  []Player `json:"other_players"`
	DiscardedCard Card     `json:"discarded_card"`
	possibleCol []Card
	possibleVal []Card
	numOfColors map[string]int
}

func (gs *Status) GetCards() (cards []string) {
	for _, c := range gs.Hand {
		cards = append(cards, c.String())
	}
	cards = append(cards, "Draw a card")

	return cards
}

func (gs *Status) GetCard(s string) (card *Card) {
	for _, c := range gs.Hand {
		if c.String() == s {
			return &c
		}
	}
	return nil
}