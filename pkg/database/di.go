package database

import (
	"fmt"
	"reflect"
)

// SensorData - структура для даних з датчиків
type SensorData struct {
	Value float64
}

// DataStorage - інтерфейс для зберігання даних
type DataStorage interface {
	Store(data string) error
}

// Database - реалізація DataStorage для роботи з базою даних
type Database struct{}

// NewDatabase - конструктор для Database
func NewDatabase() *Database {
	return &Database{}
}

// Store - зберігає дані в базу даних
func (db *Database) Store(data string) error {
	fmt.Println("Зберігаємо дані в базу даних:", data)
	return nil
}

// SensorService - сервіс для роботи з датчиками
type SensorService struct {
	Storage DataStorage
}

// NewSensorService - конструктор для SensorService
func NewSensorService(storage DataStorage) *SensorService {
	return &SensorService{Storage: storage}
}

func NewSensorServiceNoInterface(storage *Database) *SensorService {
	return &SensorService{Storage: storage}
}

// ProcessSensorData - обробляє дані з датчика та зберігає їх
func (s *SensorService) ProcessSensorData(data string) error {
	return s.Storage.Store(data)
}

// DIContainer - контейнер для Dependency Injection
type DIContainer struct {
	services map[reflect.Type]any
}

// NewDIContainer - створює новий DI контейнер
func NewDIContainer() *DIContainer {
	return &DIContainer{
		services: make(map[reflect.Type]any),
	}
}

// Provide - реєструє сервіс у контейнері
func (c *DIContainer) Provide(service any) {
	c.services[reflect.TypeOf(service)] = service
}

// Get - отримує сервіс із контейнера
func (c *DIContainer) Get(serviceType reflect.Type) any {
	return c.services[serviceType]
}

// Get - отримує сервіс із контейнера
func (c *DIContainer) Invoke() {
	// magic
}
