package main

import (
	"fmt"
	"gocourse11/pkg/database"
	"reflect"
)

func main() {
	// builder
	externalT, externalR, externalUF := database.SensorType{"Temperature"}, 100.0, 5

	sensorBuilder := &database.SpecificSensorBuilder{}
	director := database.NewDirector(sensorBuilder)
	temperatureSensor := director.Construct(externalT, externalR, externalUF)

	fmt.Printf("Sensor model: %s\n", temperatureSensor)

	// adapter
	var markdownWriter database.Writer = temperatureSensor
	if temperatureSensor.Format() == `html` {
		markdownWriter = database.MarkdownAdapter{temperatureSensor}
	}

	// di
	// Створення DI контейнера
	container := database.NewDIContainer()

	// Реєстрація залежностей
	db := database.NewDatabase()
	container.Provide(db) // Зареєструвати db

	//fmt.Printf("%+v\n", reflect.TypeOf((*database.DataStorage)(nil)))
	//fmt.Printf("%+v\n", container.Get(reflect.TypeOf((database.DataStorage)(nil))))

	// Створення SensorService із залежністю, отриманою через DI
	storage := container.Get(reflect.TypeOf((*database.Database)(nil))).(*database.Database)
	sensorService := database.NewSensorService(storage)
	// Використання SensorService
	err := sensorService.ProcessSensorData(markdownWriter.Data())
	if err != nil {
		fmt.Println("Помилка при обробці даних з датчика:", err)
	}
}
