package app

import "fmt"

type Application struct {
	cfg *Config
}

func New(cfg *Config) *Application {
	return &Application{cfg: cfg}
}

func (a *Application) Run() error {
	fmt.Println("Application started")
	return nil
}
