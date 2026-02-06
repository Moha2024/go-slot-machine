package models

import (
	"slices"
	"testing"
)

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
