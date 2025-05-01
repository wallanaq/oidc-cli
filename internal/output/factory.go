package output

func NewWriter(output string) Writer {
	switch output {
	case "", "-":
		return &ConsoleWriter{}
	default:
		return &FileWriter{Path: output}
	}
}
