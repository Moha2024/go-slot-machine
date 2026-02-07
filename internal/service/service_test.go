package service

import (
	"slotmachine/internal/domain/models"
	"testing"
)

func TestIsBet(t *testing.T) {
	type data struct {
		name    string
		command string
		want    bool
	}

	testData := []data{
		{
			name:    "Valid positive number",
			command: "100",
			want:    true,
		},
		{
			name:    "Valid zero",
			command: "0",
			want:    true,
		},
		{
			name:    "Valid negative number",
			command: "-50",
			want:    true, // Atoi считает это числом, валидация на положительность будет позже
		},
		{
			name:    "Empty string",
			command: "",
			want:    false,
		},
		{
			name:    "String with letters",
			command: "10abc",
			want:    false,
		},
		{
			name:    "Just letters",
			command: "bet",
			want:    false,
		},
		{
			name:    "Float number",
			command: "10.5",
			want:    false, // Atoi работает только с целыми числами
		},
		{
			name:    "Spaces around number",
			command: " 50 ",
			want:    false, // Atoi не обрезает пробелы автоматически
		},
		{
			name:    "Special characters",
			command: "$#%!",
			want:    false,
		},
		{
			name:    "Very large number",
			command: "999999999999999999999999999",
			want:    false, // Ошибка переполнения (overflow) тоже возвращает err != nil
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			got := IsBet(td.command)
			if td.want != got {
				t.Errorf("Isbet(%q) = %v; want %v", td.command, got, td.want)
			}
		})
	}
}

func TestValidateBet(t *testing.T) {
	type data struct {
		command string
		balance uint
		name    string
		want    uint
	}

	testData := []data{
		{
			name:    "Valid bet: less than balance",
			command: "50",
			balance: 100,
			want:    50,
		},
		{
			name:    "Valid bet: equal to balance (All-in)",
			command: "100",
			balance: 100,
			want:    100,
		},
		{
			name:    "Invalid bet: greater than balance",
			command: "150",
			balance: 100,
			want:    0,
		},
		{
			name:    "Invalid bet: zero value",
			command: "0",
			balance: 100,
			want:    0,
		},
		{
			name:    "Invalid bet: negative value",
			command: "-10",
			balance: 100,
			want:    0,
		},
		{
			name:    "Invalid bet: non-numeric string",
			command: "abc",
			balance: 100,
			want:    0,
		},
		{
			name:    "Invalid bet: empty string",
			command: "",
			balance: 100,
			want:    0,
		},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			got := ValidateBet(td.command, td.balance)
			if td.want != got {
				t.Errorf("FAIL: %s\nInput: command=%q, balance=%d\nGot: %d\nWant: %d",
					td.name, td.command, td.balance, got, td.want)
			}
		})
	}
}

func TestQuitGame(t *testing.T) {
	type data struct {
		currentBalance uint
		startBalance   uint
		name           string
		wantProfit     uint
		wantFlag       bool
	}

	testData := []data{
		{
			name:           "Win: started with 100, ended with 150",
			startBalance:   100,
			currentBalance: 150,
			wantProfit:     50,
			wantFlag:       true,
		},
		{
			name:           "Loss: started with 500, ended with 400",
			startBalance:   500,
			currentBalance: 400,
			wantProfit:     100,
			wantFlag:       false,
		},
		{
			name:           "Even: no change",
			startBalance:   1000,
			currentBalance: 1000,
			wantProfit:     0,
			wantFlag:       true,
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			p := models.NewPlayer("Tester", td.currentBalance)
			gotProfit, gotFlag := QuitGame(p, td.startBalance)

			if gotProfit != td.wantProfit || gotFlag != td.wantFlag {
				t.Errorf("FAIL: %s\nGot: (%d, %v)\nWant: (%d, %v)",
					td.name, gotProfit, gotFlag, td.wantProfit, td.wantFlag)
			}
		})
	}
}
