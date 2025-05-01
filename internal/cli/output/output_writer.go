package output

type OutputWriter interface {
	Write(data []byte) error
}
