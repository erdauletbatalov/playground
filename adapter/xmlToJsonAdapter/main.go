package main

import (
	"encoding/json"
	"fmt"
)

// Интерфейс сторонней библиотеки, принимающий JSON данные
type AnalyticsLibrary interface {
	AnalyzeData(data string)
}

// Реализация библиотеки аналитики, работающей с JSON
type JSONAnalytics struct{}

func (a *JSONAnalytics) AnalyzeData(data string) {
	fmt.Println("Analyzing JSON data:", data)
}

// XML структура для котировок
type XMLData struct {
	Symbol string
	Price  float64
}

type XMLToJSONAdapter struct {
	analytics AnalyticsLibrary
}

// Конструктор адаптера, принимает объект AnalyticsLibrary
func NewXMLToJSONAdapter(analytics AnalyticsLibrary) *XMLToJSONAdapter {
	return &XMLToJSONAdapter{analytics: analytics}
}

// Метод адаптера для обработки XML данных и передачи их в формате JSON
func (adapter *XMLToJSONAdapter) AnalyzeXMLData(xmlData XMLData) {
	// Конвертация XML структуры в JSON
	jsonData, err := json.Marshal(xmlData)
	if err != nil {
		fmt.Println("Error converting XML to JSON:", err)
		return
	}

	// Передача данных в библиотеку аналитики в формате JSON
	adapter.analytics.AnalyzeData(string(jsonData))
}

func main() {
	// Создаем экземпляр JSON аналитики
	jsonAnalytics := &JSONAnalytics{}

	// Создаем адаптер, передавая в него JSON аналитику
	adapter := NewXMLToJSONAdapter(jsonAnalytics)

	// XML данные, которые получаем из биржевого источника
	xmlData := XMLData{
		Symbol: "AAPL",
		Price:  150.25,
	}

	// Анализируем данные через адаптер
	adapter.AnalyzeXMLData(xmlData)
}
