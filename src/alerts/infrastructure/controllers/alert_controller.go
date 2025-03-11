package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"alerts/src/alerts/application"
)

type AlertController struct {
	UseCase *applications.AlertUseCase
}

func NewAlertController(uc *applications.AlertUseCase) *AlertController {
	return &AlertController{UseCase: uc}
}

// GET /alerts
func (c *AlertController) GetAllAlertsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	alerts, err := c.UseCase.GetAllAlerts(ctx)
	if err != nil {
		http.Error(w, "Error al obtener alertas", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alerts)
}
