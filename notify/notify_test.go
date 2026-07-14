package notify

import (
	"testing"
)

func TestNotify(t *testing.T) {
	t.Run("одно сообщение", func(t *testing.T) {
		sender := &BufferSender{}
		Notify(sender, "Hello, world!")
		if sender.Written() != "Hello, world!" {
			t.Errorf("Expected 'Hello world!', got '%s'", sender.Written())
		}

	})
	t.Run("два сообщения подряд", func(t *testing.T) {
		sender := &BufferSender{}
		Notify(sender, "Hello, world!")
		Notify(sender, "Bazar!")
		if sender.Written() != "Hello, world!Bazar!" {
			t.Errorf("Expected 'Hello world!Bazar!', got '%s'", sender.Written())
		}
	})

	t.Run("Пустое сообщение", func(t *testing.T) {
		sender := &BufferSender{}
		err := Notify(sender, "")
		if err != nil {
			t.Errorf("не ожидали ошибку, но получили %v", err)
		}
		if sender.Written() != "" {
			t.Errorf("Expected '', got '%s'", sender.Written())
		}
	})


}

