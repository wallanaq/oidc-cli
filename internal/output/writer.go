package output

type Writer interface {
	Write(data []byte) error
}

func NewWriter(output string) Writer {
	switch output {
	case "", "-":
		return &ConsoleWriter{}
	default:
		return &FileWriter{Path: output}
	}
}
