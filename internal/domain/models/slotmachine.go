package models

type SlotMachine struct {
	symbols     map[string]uint
	multipliers map[string]uint
	rows        int
	cols        int
	generator NumberGenerator
	reel []string
}

func NewSlotMachine(rows int, cols int, symbols map[string]uint, multipliers map[string]uint, gen NumberGenerator) SlotMachine {
	return SlotMachine{
		symbols:     symbols,
		multipliers: multipliers,
		rows:        rows,
		cols:        cols,
		generator: gen,
		reel: GenerateReel(symbols),
	}
}

func (s *SlotMachine) GetReel() []string{
	return s.reel
}

func (s *SlotMachine) GetSymbols() map[string]uint {
	return s.symbols
}

func GenerateReel(symbols map[string]uint) []string {
	reel := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			reel = append(reel, symbol)
		}
	}
	return reel
}

func (s *SlotMachine) GetSpinResult(reel []string) [][]string {
	result := [][]string{}
	for i := 0; i < s.rows; i++ {
		result = append(result, []string{})
	}
	for col := 0; col < s.cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < s.rows; row++ {
			for true {
				randomIndex := s.generator.NumberGenerator(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

func (s *SlotMachine) GetWinningLines(spinResult [][]string) []uint {
	lines := []uint{}
	for _, row := range spinResult {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, s.multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

func (s *SlotMachine) GetProfit(bet uint, winningLines []uint) uint {
	var profit uint
	for _, multi := range winningLines {
		win := multi * bet
		profit += win
	}
	return profit
}
