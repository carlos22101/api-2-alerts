package applications

import (
	"context"
	"alerts/src/alerts/domain/entities"
	"alerts/src/alerts/domain/repositories"
)

type AlertUseCase struct {
	Repo repositories.AlertRepository
}

func NewAlertUseCase(repo repositories.AlertRepository) *AlertUseCase {
	return &AlertUseCase{Repo: repo}
}

func (uc *AlertUseCase) GetAllAlerts(ctx context.Context) ([]entities.Alert, error) {
	return uc.Repo.GetAll(ctx)
}
