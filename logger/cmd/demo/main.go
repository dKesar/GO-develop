package main // Пакет main нужен для файла, из которого запускается программа.

import "logger" // Подключаем наш пакет logger из корня модуля.

func main() { // main - точка входа: с этой функции начинается выполнение программы.
	loggers := []logger.Logger{ // Создаем срез логгеров; типы берем через имя пакета logger.
		&logger.ConsoleLogger{}, // Создаем ConsoleLogger из пакета logger.
		&logger.MemoryLogger{},  // Создаем MemoryLogger из пакета logger.
	}

	logger.ProcessAll(loggers, "This is a test message")      // Отправляем обычное сообщение во все логгеры.
	logger.ProcessAll(loggers, "ERROR: something went wrong") // Отправляем сообщение-ошибку во все логгеры.
}
