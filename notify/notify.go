package notify

import (
	"bytes"
	"fmt"
	"os"
)

type Sender interface {
	Send(message string) error
}

type ConsoleSender struct{}

type BufferSender struct {
	buf bytes.Buffer
}

func (w *BufferSender) Written() string {
	return w.buf.String()
}

func (console *ConsoleSender) Send(message string) error {
	_, err := fmt.Fprintf(os.Stdout, "%s", message)
	return err
}

func (s *BufferSender) Send(message string) error {
	_, err := fmt.Fprintf(&s.buf, "%s", message)
	return err
}
func Notify(sender Sender, message string)error{
	return sender.Send(message)
}
