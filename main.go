package main

import (
	"slotmachine/internal/config"
	"slotmachine/internal/domain/models"
	"slotmachine/internal/pkg/rng"
	"slotmachine/internal/service"
	"slotmachine/internal/ui"
)

func main() {
	name := ui.GetName()
	player := models.NewPlayer(name, config.StartBalance)
	realGen := &rng.RealGenerator{}
	slot := models.NewSlotMachine(config.Rows, config.Cols, config.Symbols, config.Multipliers, realGen)
	svc := service.NewGameService(slot)
	ui.PlayGame(svc, player)
}
