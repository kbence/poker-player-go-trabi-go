package main


func GetPokerLevel(cards []Card) (int, int) {
	pokerTypes := make(map[string]int)
	pokerCards := make(map[string]int)
	level := 1
	for _, c := range cards {
		if(pokerCards[c.Rank] == 2 && pokerTypes["drill"] == 0) {
			pokerTypes["drill"]++
		}
		if(pokerCards[c.Rank] == 1) {
			pokerTypes["pairs"]++
		}
		pokerCards[c.Rank]++
		if(len(cards) == 1) {
			return 1, c.Position()
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


	return level, 1
}
