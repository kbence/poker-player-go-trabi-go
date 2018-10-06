package main


func GetPokerLevel(cards []Card) (int, int) {
	pokerTypes := make(map[string]int)
	pokerCards := make(map[string]int)
	pokerColors := make(map[string]int)
	cardPosition := 0
	level := 1
	for _, c := range cards {
		if(pokerCards[c.Rank] == 2 && pokerTypes["drill"] == 0) {
			pokerTypes["drill"]++
		}
		if(pokerCards[c.Rank] == 1) {
			pokerTypes["pairs"]++
		}
		pokerColors[c.Suit]++
		if(pokerColors[c.Suit] == 5) {
			pokerTypes["flush"]++
		}
		pokerCards[c.Rank]++
		if(pokerCards[c.Rank] == 4) {
			return 8, 1
		}
		if(cardPosition < c.Position()) {
			cardPosition = c.Position()
		}
	}
	if(pokerTypes["pairs"] == 1) {
		level = 2
	}
	if(pokerTypes["pairs"] == 2 || pokerTypes["pairs"] == 3) {
		level = 3
	}
	if(pokerTypes["drill"] == 1) {
		level = 4
	}
	if(pokerTypes["flush"] != 0) {
		level = 6
	}
	if(level == 1) {
		return 1, cardPosition
	}


	return level, 1
}
