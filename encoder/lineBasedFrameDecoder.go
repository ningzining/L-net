package encoder

const (
	lineBreaks = '\n'
)

type LineBasedFrameEncoder struct {
}

func NewLineBasedFrameDecoder() Encoder {
	return &LineBasedFrameEncoder{}
}

func (d LineBasedFrameEncoder) Encode(bytes []byte) ([]byte, error) {
	return append(bytes, lineBreaks), nil
}
