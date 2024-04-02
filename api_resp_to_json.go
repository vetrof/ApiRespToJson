package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	var url string
	fmt.Println("api url: ")
	fmt.Scanf("%s\n", &url)

	// URL API
	// url := "https://api.sampleapis.com/beers/ale"

	// Отправка GET запроса
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при выполнении GET запроса: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Ошибка: %s\n", resp.Status)
		os.Exit(1)
	}

	// Декодирование JSON ответа
	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при декодировании JSON: %v\n", err)
		os.Exit(1)
	}

	// Форматирование JSON с отступами
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при форматировании JSON: %v\n", err)
		os.Exit(1)
	}

	// Вывод JSON данных в терминал
	fmt.Println(string(prettyJSON))

	// Получаем текущую дату и время
	currentTime := time.Now()

	// Форматируем текущую дату и время в строку
	dateTimeStr := currentTime.Format("2006-01-02_15-04-05")

	// Формируем имя файла с текущей датой и временем
	fileName := fmt.Sprintf("json_out_%s.json", dateTimeStr)

	// Создаем файл с указанным именем
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при создании файла: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(prettyJSON)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при записи в файл: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Данные успешно записаны в файл output.json")
}
