package models

import "errors"

type Player struct {
	name    string
	balance uint
}

func NewPlayer(newName string, startBalance uint) *Player {
	return &Player{
		name:    newName,
		balance: startBalance,
	}
}

func (p *Player) GetBalance() uint {
	return p.balance
}

func (p *Player) SetBalance(balance uint) {
	p.balance = balance
}

func (p *Player) UpdateBalance(profit uint, bet uint) error {
	if bet > p.balance {
		return errors.New("Invalid bet")
	}
	p.balance -= uint(bet)
	p.balance += uint(profit)
	return nil
}
