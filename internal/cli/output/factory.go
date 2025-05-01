package output

func NewOutputWriter(path string) (OutputWriter, error) {
	switch path {
	case "", "-":
		return &ConsoleWriter{}, nil
	default:
		return &FileWriter{Path: path}, nil
	}
}
