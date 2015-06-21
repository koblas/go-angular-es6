package service

import (
	"github.com/koblas/likemark/app"
	"github.com/koblas/likemark/conf"
)

type LikeMarkService struct {
}

func (s *LikeMarkService) Migrate(cfg conf.Config) error {
	a := app.Application{Config: cfg}

	return a.Migrate()
}

func (s *LikeMarkService) Run(cfg conf.Config) error {
	a := app.Application{Config: cfg}

	a.Init()

	a.Run()

	return nil
}
