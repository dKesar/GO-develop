package logger // Объявляем пакет logger: этот файл относится к пакету с кодом для логирования.

import ( // Подключаем внешние пакеты, которые нужны ниже в коде.
	"fmt"     // fmt нужен, чтобы печатать сообщения в консоль через fmt.Println.
	"strings" // strings нужен, чтобы проверять начало строки и менять регистр текста.
)

// Logger описывает общее поведение любого логгера в этой программе.
type Logger interface {
	Log(message string) // Log принимает текст сообщения и куда-то его записывает или выводит.
	Errors() []string   // Errors возвращает список сообщений, которые считаются ошибками.
}

// ConsoleLogger выводит сообщения в консоль и дополнительно хранит их в памяти.
type ConsoleLogger struct {}

// MemoryLogger хранит сообщения только в памяти и ничего не печатает в консоль.
type MemoryLogger struct {
	errors []string // errors хранит только сообщения, которые начинаются с "ERROR:".
	logs   []string // logs хранит все сообщения, которые прошли через этот логгер.
}

// isErrorMessage проверяет, является ли сообщение ошибкой.
func ErrorMessage(message string) bool {
	upperMessage := strings.ToUpper(message)         // Приводим сообщение к верхнему регистру, чтобы "error:" и "ERROR:" считались одинаково.
	return strings.HasPrefix(upperMessage, "ERROR:") // Возвращаем true, если сообщение начинается с префикса "ERROR:".
}

// Errors возвращает все ошибки, которые накопил MemoryLogger.
func (l *MemoryLogger) Errors() []string {
	return l.errors // Отдаем срез ошибок, сохраненный внутри структуры MemoryLogger.
}

// Log сохраняет сообщение внутри MemoryLogger.
func (l *MemoryLogger) Log(message string) {
	l.logs = append(l.logs, message) // Добавляем любое сообщение в общий список логов.

	if ErrorMessage(message) { // Проверяем, начинается ли сообщение с "ERROR:".
		l.errors = append(l.errors, message) // Если это ошибка, добавляем сообщение еще и в список ошибок.
	}
}

// Log выводит сообщение в консоль и сохраняет его внутри ConsoleLogger.
func (l *ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

// Errors возвращает все ошибки, которые накопил ConsoleLogger.
func (l *ConsoleLogger) Errors() []string {
	return l.errors // Отдаем срез ошибок, сохраненный внутри структуры ConsoleLogger.
}

// ProcessAll отправляет одно сообщение сразу во все переданные логгеры.
func ProcessAll(loggers []Logger, message string) {
	for _, logger := range loggers { // Проходим по каждому логгеру из среза loggers.
		logger.Log(message) // Вызываем Log у текущего логгера; конкретный тип может быть ConsoleLogger или MemoryLogger.
	}
}

func GetLogger(useMemory bool) Logger {
	if useMemory {
		return &MemoryLogger{}
	}
	return &ConsoleLogger{}
}
