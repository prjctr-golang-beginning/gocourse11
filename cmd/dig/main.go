package main

import (
	"fmt"
	"go.uber.org/dig"
	"gocourse11/pkg/database"
)

func main() {
	// Створення DI контейнера
	container := dig.New()

	// Реєстрація залежностей
	err := container.Provide(database.NewDatabase)
	if err != nil {
		panic(err)
	}
	err = container.Provide(database.NewSensorServiceNoInterface)
	if err != nil {
		panic(err)
	}

	// Вирішення залежностей та створення SensorService
	err = container.Invoke(func(service *database.SensorService) {
		// Використання SensorService
		err := service.ProcessSensorData(`some sensor data`)
		if err != nil {
			fmt.Println("Помилка при обробці даних з датчика:", err)
		}
	})

	if err != nil {
		panic(err)
	}
}
