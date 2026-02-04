package main

import (
	"context"
	"log"
	"time"

	postgres "Sybersports/pgk/postgresql"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := postgres.CreateConnection(ctx)
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer conn.Close(ctx)

	log.Println("Подключение успешно!")
}