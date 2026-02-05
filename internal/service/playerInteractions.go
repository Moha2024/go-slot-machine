package service

import (
	"slotmachine/internal/config"
	"slotmachine/internal/domain/models"
	"strconv"
)

func IsBet(command string) bool {
	_, err := strconv.Atoi(command)
	if err != nil {
		return false
	}
	return true
}

func RestartGame(player *models.Player) {
	player.SetBalance(config.StartBalance)
}

func QuitGame(player *models.Player) (profit uint, flag bool){
	balance := player.GetBalance()
	if balance >= config.StartBalance{
		return balance - config.StartBalance, true
	} else {
		return config.StartBalance - balance, false
	}
}

func ValidateBet(command string, balance uint) uint {
	bet, _ := strconv.Atoi(command)
	if bet <= 0 {
		return 0
	} else if uint(bet) > balance {
		return 0
	}
	return uint(bet)
}

