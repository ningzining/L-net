package encoder

type Encoder interface {
	Encode(bytes []byte) ([]byte, error)
}
