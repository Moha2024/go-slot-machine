package models

import "testing"

func TestPlayer_UpdateBalance(t *testing.T) { // complex test of main logic, get/set functions and constructor
	type data struct {
		profit       uint
		bet          uint
		name         string
		startBalance uint
		wantBalance  uint
		wantErr      bool
	}

	testData := []data{
		{
			name:         "Successful spin: win more than bet",
			startBalance: 100,
			bet:          10,
			profit:       50,
			wantBalance:  140, // 100 - 10 + 50
			wantErr:      false,
		},
		{
			name:         "All-in bet: balance becomes profit",
			startBalance: 50,
			bet:          50,
			profit:       100,
			wantBalance:  100, // 50 - 50 + 100
			wantErr:      false,
		},
		{
			name:         "Insufficient funds: bet too high",
			startBalance: 20,
			bet:          30,
			profit:       0,
			wantBalance:  20, // Баланс не должен измениться
			wantErr:      true,
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			p := NewPlayer("Test Player", td.startBalance)
			err := p.UpdateBalance(td.profit, td.bet)

			if (err != nil) != td.wantErr {
				t.Errorf("UpdateBalance() error = %v, wantErr %v", err, td.wantErr)
			}

			if p.GetBalance() != td.wantBalance {
				t.Errorf("Final balance = %v, want %v", p.GetBalance(), td.wantBalance)
			}
		})
	}
}
