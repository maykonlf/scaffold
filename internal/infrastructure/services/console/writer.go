package console

type Writer interface {
	Write(buf []byte) (n int, err error)
	Flush() error
}
