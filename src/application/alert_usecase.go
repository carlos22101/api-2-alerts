package applications

import (
	"context"
	"alerts/src/domain/entities"
	"alerts/src/domain/repositories"
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
