package output

import "os"

type ConsoleWriter struct{}

func (w *ConsoleWriter) Write(data []byte) error {
	_, err := os.Stdout.Write(data)
	os.Stdout.WriteString("\n")
	return err
}
