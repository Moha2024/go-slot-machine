package ui

import (
	"slotmachine/internal/config"
	"slotmachine/internal/domain/models"
	"slotmachine/internal/service"
	"strings"
)

func PlayGame(svc *service.GameService, player *models.Player) {
Loop:
	for {
		command := getCommand()
		command = strings.ToUpper(strings.TrimSpace(command))
		balance := player.GetBalance()
		ok := service.IsBet(command)
		if ok {
			bet := service.ValidateBet(command, balance)

			if bet == 0 {
				invalidBetMessage()
				continue
			}

			roundResult, err := svc.ExecuteRound(player, bet)
			if err != nil {
				invalidBetMessage()
				continue
			}
			balance = roundResult.NewBalance

			printSpinResult(roundResult.SpinResult)
			printWinLose(roundResult.Profit, bet)
			printBalance(balance)

			if balance == 0 {
				printGameOver(service.QuitGame(player, config.StartBalance))
				break
			}
		} else {
			switch command {
			case "QUIT":
				printGameOver(service.QuitGame(player, config.StartBalance))
				break Loop
			case "RESTART":
				service.RestartGame(player)
				restartMessage()
			default:
				wrongCommandMessage()
			}
		}
	}
}
