package consumer

import (
    "context"
    "encoding/json"
    "log"
    "os"
    "time"

    "alerts/src/core"
    amqp "github.com/rabbitmq/amqp091-go"
)

// consumeMessages escucha la cola de RabbitMQ y procesa cada mensaje recibido
func ConsumeMessages(ch *amqp.Channel, dbConn core.DBWrapperInterface) {
    queueName := os.Getenv("RABBITMQ_QUEUE")
    msgs, err := ch.Consume(
        queueName, // nombre de la cola
        "",        // consumer
        true,      // auto-ack
        false,     // exclusive
        false,     // no-local
        false,     // noWait
        nil,       // args
    )
    if err != nil {
        log.Fatal("Error al consumir la cola:", err)
    }

    log.Println("Esperando mensajes en la cola:", queueName)
    for msg := range msgs {
        log.Println("Mensaje recibido:", string(msg.Body))
        //Procesar el mensaje: convertirlo en alerta y guardarlo en la BD
        if err := processMessage(msg.Body, dbConn); err != nil {
        log.Println("Error procesando mensaje:", err)
        }
    }
}


 //processMessage deserializa el mensaje y guarda la alerta en la BD
func processMessage(body []byte, dbConn core.DBWrapperInterface) error {
    var mov struct {
        ID        int    `json:"id"`
        SensorID  string `json:"sensorId"`
        Timestamp string `json:"timestamp"`
        Motion    bool   `json:"motion"`
        Status    string `json:"status"` // Añadir el campo status
    }
    if err := json.Unmarshal(body, &mov); err != nil {
        return err
    }

    description := "Movimiento detectado por sensor " + mov.SensorID + " en la Sala de estar "

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    query := `INSERT INTO alerts (sensor_id, event_timestamp, description, status) VALUES (?, ?, ?, ?)`
    _, err := dbConn.GetDB().ExecContext(ctx, query, mov.SensorID, mov.Timestamp, description, mov.Status) 
    return err
}