package service

import (
	"github.com/koblas/go-angular-es6/app"
	"github.com/koblas/go-angular-es6/conf"
)

type LikeMarkService struct {
}

func (s *LikeMarkService) Migrate(cfg *conf.ConfigData) error {
	a := app.Application{Config: cfg}

	return a.Migrate()
}

func (s *LikeMarkService) Run(cfg *conf.ConfigData) error {
	a := app.Application{Config: cfg}

	a.Init()

	a.Run()

	return nil
}
