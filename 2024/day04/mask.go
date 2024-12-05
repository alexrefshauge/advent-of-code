package main

func mask(block []string, hSkip, vSkip, count int, cursorX, cursorY int) string {
	result := ""
	for n := range count {
		x := cursorX + n*hSkip
		y := cursorY + n*vSkip
		if y < 0 || y > len(block)-1 || x < 0 || x > len(block[cursorY])-1 {
			continue
		}
		result += string(rune(block[y][x]))
	}
	return result
}
