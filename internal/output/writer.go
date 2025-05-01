package output

type Writer interface {
	Write(data []byte) error
}
