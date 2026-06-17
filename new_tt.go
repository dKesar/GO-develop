package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Самый простой случай: запуск горутины
	go say("Привет из горутины!")

	// main тоже горутина — без паузы программа завершится раньше, чем успеет say
	time.Sleep(100 * time.Millisecond)

	// 2. Несколько горутин работают параллельно
	go say("Один")
	go say("Два")
	go say("Три")

	time.Sleep(100 * time.Millisecond)

	// 3. Анонимная горутина
	go func(msg string) {
		fmt.Println(msg)
	}("Анонимная горутина")

	time.Sleep(100 * time.Millisecond)
	fmt.Println("main завершился")
}

func say(text string) {
	fmt.Println(text)
}
