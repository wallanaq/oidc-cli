package output

func NewWriter(path string) (Writer, error) {
	switch path {
	case "", "-":
		return &ConsoleWriter{}, nil
	default:
		return &FileWriter{Path: path}, nil
	}
}
