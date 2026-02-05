package main

import (
	"slotmachine/internal/config"
	"slotmachine/internal/domain/models"
	"slotmachine/internal/service"
	"slotmachine/internal/ui"
)

func main() {
	player := models.NewPlayer(ui.GetName(), config.StartBalance)
	slot := models.NewSlotMachine(config.Rows, config.Cols, config.Symbols, config.Multipliers)
	svc := service.NewGameService(slot)
	ui.PlayGame(svc, player)
}
