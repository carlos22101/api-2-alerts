package routes

import (
    "database/sql"
    "log"
    "net/http"
    "alerts/src/alerts/application"
    "alerts/src/alerts/infrastructure/controllers"
    repo "alerts/src/alerts/infrastructure/repositories"

    "github.com/gorilla/mux"
)

func SetupAlertRoutes(router *mux.Router, dbWrapper interface{}) {
    // Convertir dbWrapper a *sql.DB
    db, ok := dbWrapper.(*sql.DB)
    if !ok {
        log.Fatal("dbWrapper no es *sql.DB")
    }

    // Verificar la conexión a la base de datos
    if err := db.Ping(); err != nil {
        log.Fatal("No se puede conectar a la base de datos:", err)
    }

    alertRepo := repo.NewAlertMySQLRepo(db)
    alertUC := applications.NewAlertUseCase(alertRepo)
    alertCtrl := controllers.NewAlertController(alertUC)

    router.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Handling /alerts request")
        alertCtrl.GetAllAlertsHandler(w, r)
    }).Methods(http.MethodGet)
}