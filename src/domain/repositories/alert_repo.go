package repositories

import (
	"context"
	"alerts/src/domain/entities"
)

type AlertRepository interface {
	GetAll(ctx context.Context) ([]entities.Alert, error)
}
