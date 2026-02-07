package models

import (
	"slices"
	"sort"
	"testing"
)

type mockGenerator struct {
	numbers []int
	cursor  int
}

func (m *mockGenerator) NumberGenerator(min int, max int) int {
	num := m.numbers[m.cursor]
	m.cursor++
	return num
}

func TestGetWinningLines(t *testing.T) {
	s := SlotMachine{
		multipliers: map[string]uint{
			"A": 20,
			"B": 10,
			"C": 5,
			"D": 2,
		},
	}
	type data struct {
		lines [][]string
		want  []uint
		name  string
	}
	testData := []data{
		// --- СИМВОЛ A (Множитель 20) ---
		{
			name: "Symbol A: Top line win",
			lines: [][]string{
				{"A", "A", "A"},
				{"B", "C", "D"},
				{"C", "D", "B"},
			},
			want: []uint{20, 0, 0},
		},
		{
			name: "Symbol A: Top and Middle lines win",
			lines: [][]string{
				{"A", "A", "A"},
				{"A", "A", "A"},
				{"C", "D", "B"},
			},
			want: []uint{20, 20, 0},
		},
		{
			name: "Symbol A: All three lines win (Jackpot)",
			lines: [][]string{
				{"A", "A", "A"},
				{"A", "A", "A"},
				{"A", "A", "A"},
			},
			want: []uint{20, 20, 20},
		},

		// --- СИМВОЛ B (Множитель 10) ---
		{
			name: "Symbol B: Middle line win",
			lines: [][]string{
				{"A", "C", "D"},
				{"B", "B", "B"},
				{"D", "A", "C"},
			},
			want: []uint{0, 10, 0},
		},
		{
			name: "Symbol B: Middle and Bottom lines win",
			lines: [][]string{
				{"A", "C", "D"},
				{"B", "B", "B"},
				{"B", "B", "B"},
			},
			want: []uint{0, 10, 10},
		},
		{
			name: "Symbol B: All three lines win",
			lines: [][]string{
				{"B", "B", "B"},
				{"B", "B", "B"},
				{"B", "B", "B"},
			},
			want: []uint{10, 10, 10},
		},

		// --- СИМВОЛ C (Множитель 5) ---
		{
			name: "Symbol C: Bottom line win",
			lines: [][]string{
				{"A", "B", "D"},
				{"D", "A", "B"},
				{"C", "C", "C"},
			},
			want: []uint{0, 0, 5},
		},
		{
			name: "Symbol C: Top and Bottom lines win",
			lines: [][]string{
				{"C", "C", "C"},
				{"D", "A", "B"},
				{"C", "C", "C"},
			},
			want: []uint{5, 0, 5},
		},
		{
			name: "Symbol C: All three lines win",
			lines: [][]string{
				{"C", "C", "C"},
				{"C", "C", "C"},
				{"C", "C", "C"},
			},
			want: []uint{5, 5, 5},
		},

		// --- СИМВОЛ D (Множитель 2) ---
		{
			name: "Symbol D: Top line win",
			lines: [][]string{
				{"D", "D", "D"},
				{"A", "B", "C"},
				{"C", "A", "B"},
			},
			want: []uint{2, 0, 0},
		},
		{
			name: "Symbol D: Top and Middle lines win",
			lines: [][]string{
				{"D", "D", "D"},
				{"D", "D", "D"},
				{"C", "A", "B"},
			},
			want: []uint{2, 2, 0},
		},
		{
			name: "Symbol D: All three lines win",
			lines: [][]string{
				{"D", "D", "D"},
				{"D", "D", "D"},
				{"D", "D", "D"},
			},
			want: []uint{2, 2, 2},
		},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			got := s.GetWinningLines(td.lines)
			if !slices.Equal(td.want, got) {
				t.Errorf("FAIL: %s\nGot: %v\nWant: %v", td.name, got, td.want)
			}
		})
	}
}

func TestGetProfit(t *testing.T) {
	s := SlotMachine{}
	type data struct {
		bet      uint
		winLines []uint
		want     uint
		name     string
	}
	testData := []data{
		{
			name:     "Normal win: single line",
			bet:      10,
			winLines: []uint{5},
			want:     50,
		},
		{
			name:     "Normal win: multiple lines",
			bet:      10,
			winLines: []uint{5, 10, 2},
			want:     170, // (5*10) + (10*10) + (2*10)
		},
		{
			name:     "Zero bet: profit must be zero",
			bet:      0,
			winLines: []uint{2, 5, 10},
			want:     0,
		},
		{
			name:     "No winning lines: profit must be zero",
			bet:      100,
			winLines: []uint{},
			want:     0,
		},
		{
			name:     "Nil winning lines: safety check",
			bet:      100,
			winLines: nil,
			want:     0,
		},
		{
			name:     "Zero multipliers: profit must be zero",
			bet:      10,
			winLines: []uint{0, 0, 0},
			want:     0,
		},
		{
			name:     "Large values: checking multiplication",
			bet:      1000000,
			winLines: []uint{100, 200},
			want:     300000000,
		},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			got := s.GetProfit(td.bet, td.winLines)
			if td.want != got {
				t.Errorf("FAIL: %s\nGot: %v\nWant: %v", td.name, got, td.want)
			}
		})
	}
}

func TestGetSpinResult(t *testing.T) {
	mock := &mockGenerator{}

	s := SlotMachine{
		generator: mock,
	}

	type data struct {
		name     string
		rows     int
		cols     int
		reel     []string
		mockNums []int      // сценарий "случайных" чисел
		want     [][]string // ожидаемая матрица
	}

	testData := []data{
		{
			name: "Standard 3x2 Grid",
			rows: 3,
			cols: 2,
			reel: []string{"A", "B", "C", "D"},
			// Столбец 1 получит индексы 0,1,2 (A,B,C)
			// Столбец 2 получит индексы 3,0,1 (D,A,B)
			mockNums: []int{0, 1, 2, 3, 0, 1},
			want: [][]string{
				{"A", "D"},
				{"B", "A"},
				{"C", "B"},
			},
		},
		{
			name: "Collision Check: Generator repeats index",
			rows: 2,
			cols: 1,
			reel: []string{"A", "B", "C"},
			// Генератор выдает 0, потом опять 0 (дубликат!), потом 2.
			// Функция ДОЛЖНА пропустить второй 0 и взять 2.
			mockNums: []int{0, 0, 2},
			want: [][]string{
				{"A"},
				{"C"},
			},
		},
		{
			name:     "Minimal 1x1 Grid",
			rows:     1,
			cols:     1,
			reel:     []string{"Z"},
			mockNums: []int{0},
			want: [][]string{
				{"Z"},
			},
		},
		{
			name:     "Wide Grid: 1 row, 3 columns",
			rows:     1,
			cols:     3,
			reel:     []string{"A", "B", "C"},
			mockNums: []int{0, 1, 2},
			want: [][]string{
				{"A", "B", "C"},
			},
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			s.rows = td.rows
			s.cols = td.cols
			mock.numbers = td.mockNums
			mock.cursor = 0

			got := s.GetSpinResult(td.reel)

			if !slices.EqualFunc(got, td.want, slices.Equal) {
				t.Errorf("FAIL: %s\nGot:  %v\nWant: %v", td.name, got, td.want)
			}
		})
	}
}

func TestGenerateReel(t *testing.T) {
	type data struct {
		name    string
		symbols map[string]uint
		want    []string
	}
	testData := []data{
		{
			name: "Standard distribution",
			symbols: map[string]uint{
				"A": 1,
				"B": 2,
			},
			want: []string{"A", "B", "B"},
		},
		{
			name: "Single symbol",
			symbols: map[string]uint{
				"C": 3,
			},
			want: []string{"C", "C", "C"},
		},
		{
			name:    "Empty map",
			symbols: map[string]uint{},
			want:    []string{},
		},
		{
			name: "Symbol with zero count",
			symbols: map[string]uint{
				"A": 1,
				"D": 0,
			},
			want: []string{"A"},
		},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			got := GenerateReel(td.symbols)
			sort.Strings(got)
			sort.Strings(td.want)

			if !slices.Equal(got, td.want) {
				t.Errorf("FAIL: %s\nGot: %v\nWant: %v", td.name, got, td.want)
			}
		})
	}
}
