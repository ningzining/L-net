package encoder

// DelimiterBasedFrameEncoder 分隔符分帧编码器
type DelimiterBasedFrameEncoder struct {
	delimiter byte
}

func NewDelimiterBasedFrameDecoder(delimiter byte) Encoder {
	return &DelimiterBasedFrameEncoder{
		delimiter: delimiter,
	}
}

func (d DelimiterBasedFrameEncoder) Encode(bytes []byte) ([]byte, error) {
	return append(bytes, d.delimiter), nil
}
