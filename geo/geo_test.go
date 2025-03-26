package geo_test

import (
	"testing"

	"github.com/riuzaki24/app-5/geo"
)

func TestGetMyLocation(t *testing.T) {
	//Arange - подготовка, expected результат, данные для функции
	city := "Moskva"
	expected := geo.GeoData {
		City: "Moskva",
	}

	//Act - выполняем функцию
	got, err := geo.GetMyLocation(city)

	//Assert - проверка результата с expected
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
} 
