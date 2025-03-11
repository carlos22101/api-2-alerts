package main

import (
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "github.com/gorilla/mux"

    "alerts/src/core"
    "alerts/src/alerts/infrastructure/routes"
    "alerts/src/alerts/infrastructure/consumer"
)

func main() {
    // Cargar variables de entorno
    if err := godotenv.Load(); err != nil {
        log.Println("No se pudo cargar .env, se usarán variables de entorno del sistema")
    }

    // Conexión a la base de datos
    db, err := core.ConnectDB()
    if err != nil {
        log.Fatal("Error conectando a la BD:", err)
    }
    defer db.GetDB().Close()

    // Conexión a RabbitMQ
    rabbitChan, err := core.ConnectRabbit()
    if err != nil {
        log.Fatal("Error conectando a RabbitMQ:", err)
    }
    defer rabbitChan.Close()

    // Iniciar el consumidor en una goroutine
    go consumer.ConsumeMessages(rabbitChan, db)

    // Configurar rutas de la API (para el short polling del Frontend)
    router := mux.NewRouter()
    routes.SetupAlertRoutes(router, db.GetDB())

    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }
    log.Println("Notifier API corriendo en el puerto:", port)
    // Envolver el router con el middleware CORS
    log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(router)))
}

// corsMiddleware añade los encabezados necesarios para permitir CORS.
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // O especifica un origen en particular
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}