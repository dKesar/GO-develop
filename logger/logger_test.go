package logger

import "testing"

func TestMemoryLoggerErrors(t *testing.T) {
	cases := []struct {
		name     string
		messages []string
		wantErrs int
	}{
		{
			name:     "одна ошибка из трёх",
			messages: []string{"инфо", "ERROR: сломалось", "ещё инфо"},
			wantErrs: 1,
		},
		{
			name:     "нет ошибок",
			messages: []string{"инфо", "всё хорошо"},
			wantErrs: 0,
		},
		{
			name:     "две ошибки",
			messages: []string{"ERROR: первая", "ERROR: вторая", "норм"},
			wantErrs: 2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ml := &MemoryLogger{}
			for _, msg := range c.messages {
				ml.Log(msg)
			}
			got := ml.Errors()
			if len(got) != c.wantErrs {
				t.Errorf("хотели %d ошибок, получили %d: %v",
					c.wantErrs, len(got), got)
			}
		})
	}
}

func TestProcessAll(t *testing.T) {
	ml1 := &MemoryLogger{}
	ml2 := &MemoryLogger{}
	loggers := []Logger{ml1, ml2}

	ProcessAll(loggers, "ERROR: проверка")
	if len(ml1.Errors()) != 1 {
		t.Errorf("ml1: ожидали 1 ошибку, получили %d", len(ml1.Errors()))
	}
	if len(ml2.Errors()) != 1 {
		t.Errorf("ml2: ожидали 1 ошибку, получили %d", len(ml2.Errors()))
	}
}

// BrokenGetLogger — специально сломанная версия для демонстрации ловушки.
func BrokenGetLogger(useMemory bool) Logger {
	var m *MemoryLogger
	if useMemory {
		m = &MemoryLogger{}
	}
	return m // ловушка здесь
}

func TestNilTrap(t *testing.T) {
	// Сломанная версия: useMemory=false, но вернёт "коробку" а не nil
	broken := BrokenGetLogger(false)
	if broken == nil {
		t.Error("сломанная версия вдруг вернула nil — неожиданно")
	}
	// broken != nil, хотя логгера нет — это и есть ловушка

	// Твоя рабочая версия: должна вернуть реальный объект
	real := GetLogger(false)
	if real == nil {
		t.Error("GetLogger(false) вернул nil, а должен ConsoleLogger")
	}

	// Настоящий nil выглядит так:
	var trueNil Logger
	if trueNil != nil {
		t.Error("это никогда не должно сработать")
	}
}
