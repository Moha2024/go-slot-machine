package service

import "slotmachine/internal/domain/models"

type RoundResult struct {
	SpinResult [][]string
	WinningLines []uint
	Profit uint
	NewBalance uint
}

type GameService struct{
	slot models.SlotMachine
}

func NewGameService(m models.SlotMachine) *GameService{
	return &GameService{slot: m}
}

func (g *GameService) ExecuteRound(player *models.Player, bet uint) (RoundResult, error){
	spinRes := g.slot.GetSpinResult(g.slot.GetReel())
	winLines := g.slot.GetWinningLines(spinRes)
	profit := g.slot.GetProfit(bet, winLines)

	err := player.UpdateBalance(profit, bet)
	if err != nil{
		return RoundResult{}, err
	}

	return RoundResult{
		SpinResult: spinRes,
		WinningLines: winLines,
		Profit: profit,
		NewBalance: player.GetBalance(),
	}, nil
}





