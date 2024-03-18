package encoder

const (
	lineBreaks = '\n'
)

// LineBasedFrameEncoder 行分帧编码器
type LineBasedFrameEncoder struct {
}

func NewLineBasedFrameDecoder() Encoder {
	return &LineBasedFrameEncoder{}
}

func (d LineBasedFrameEncoder) Encode(bytes []byte) ([]byte, error) {
	return append(bytes, lineBreaks), nil
}
