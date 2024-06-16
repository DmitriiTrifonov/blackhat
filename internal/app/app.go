package app

import (
	"github.com/DmitriiTrifonov/blackhat/internal/game"
	"go.uber.org/zap"
)

type Application struct {
	cfg *Config
}

func New(cfg *Config) *Application {
	return &Application{cfg: cfg}
}

func (a *Application) Run() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()
	round := game.NewRound()
	logger.Info("round data", zap.Any("round", round))
	return nil
}
