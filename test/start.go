package main

import (
	"fmt"
	"sort"
)

func main() {
	var kolvo, numValute, answer int
	valute := map[string]int {
		"NOK": 1,
		"RUB": 2,
		"KZT": 3,
		"AUD": 4,
		"CHF": 5,
		"SEK": 6,
		"INR": 7,
		"BRL": 8,
		"GBP": 9,
		"CAD": 10,
		"USD": 11,
		"JPY": 12,
		"TRY": 13,
		"EUR": 14,
		"CNY": 15,
	}
	var rates = map[string]float64{
		"USD": 1.0,   // Доллар США
		"EUR": 0.92,  // Евро
		"RUB": 90.0,  // Российский рубль
		"JPY": 157.0, // Японская иена
		"CNY": 7.25,  // Китайский юань
		"GBP": 0.78,  // Британский фунт
		"KZT": 460.0, // Казахстанский тенге
		"TRY": 32.5,  // Турецкая лира
		"INR": 83.0,  // Индийская рупия
		"BRL": 5.12,  // Бразильский реал
		"AUD": 1.50,  // Австралийский доллар
		"CAD": 1.36,  // Канадский доллар
		"CHF": 0.89,  // Швейцарский франк
		"SEK": 10.8,  // Шведская крона
		"NOK": 10.5,  // Норвежская крона
	}
	// m := make(map[string]int)
	Welcome()
	Seen(valute)
	Input(&kolvo, &numValute)
	Convert(&kolvo, &numValute, &answer, rates, valute)
}

func Welcome(){
	fmt.Println("Добро пожаловать в конвертер валют!")
	fmt.Println("Выберите валюту из списка: ")
}

func Seen(m map[string]int){
	keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }

    // Шаг 2: сортируем ключи по значению из map
    sort.Slice(keys, func(i, j int) bool {
        return m[keys[i]] < m[keys[j]] // по возрастанию
    })

    // Шаг 3: выводим
    for _, k := range keys {
        fmt.Printf("%d: %s\n", m[k], k)
    }
}

func Input(kolvo, numValute *int) {
	fmt.Println("Введите количество валюты в USD: ")
	fmt.Scan(kolvo)
	fmt.Println("Введите номер валюты из предложенного списка: ")
	fmt.Scan(numValute)

}
func Convert(kolvo, numValute, answer *int, m map[string]float64, k map[string]int){
	if kolvo == nil || numValute == nil || answer == nil {
		return
	}
	// Поиск строки валюты по её номеру
	var valuteName string
	for key, val := range k {
		if val == *numValute {
			valuteName = key
			break
		}
	}
	if valuteName == "" {
		return // валюта не найдена
	}
	rate, ok := m[valuteName]
	if !ok {
		return // курс не найден
	}
	*answer = int(float64(*kolvo) * rate)
	fmt.Println(*answer)
}