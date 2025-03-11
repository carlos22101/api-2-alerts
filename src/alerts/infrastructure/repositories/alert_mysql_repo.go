package repositories

import (
    "context"
    "database/sql"
    "time"
    "alerts/src/alerts/domain/entities"
)

type AlertMySQLRepo struct {
    DB *sql.DB
}

func NewAlertMySQLRepo(db *sql.DB) *AlertMySQLRepo {
    return &AlertMySQLRepo{DB: db}
}

func (repo *AlertMySQLRepo) GetAll(ctx context.Context) ([]entities.Alert, error) {
    query := "SELECT id, sensor_id, event_timestamp, description, status, created_at FROM alerts"
    rows, err := repo.DB.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var alerts []entities.Alert
    for rows.Next() {
        var a entities.Alert
        var createdAt string
        if err := rows.Scan(&a.ID, &a.SensorID, &a.EventTime, &a.Description, &a.Status, &createdAt); err != nil {
            return nil, err
        }
        // Convertir createdAt a time.Time
        a.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
        if err != nil {
            return nil, err
        }
        alerts = append(alerts, a)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return alerts, nil
}