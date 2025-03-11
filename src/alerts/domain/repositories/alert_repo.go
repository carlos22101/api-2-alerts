package repositories

import (
	"context"
	"alerts/src/alerts/domain/entities"
)

type AlertRepository interface {
	GetAll(ctx context.Context) ([]entities.Alert, error)
}
