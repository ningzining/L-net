package decoder

import (
	"bytes"
	"errors"
	"strings"
)

// DelimiterBasedFrameDecoder 分隔符分帧解码器
type DelimiterBasedFrameDecoder struct {
	delimiter byte
}

func NewDelimiterBasedFrameDecoder(delimiter byte) Decoder {
	return &DelimiterBasedFrameDecoder{
		delimiter: delimiter,
	}
}

func (d DelimiterBasedFrameDecoder) Decode(buffer *bytes.Buffer) [][]byte {
	frames := make([][]byte, 0)
	for {
		decode, err := d.decode(buffer)
		if err != nil {
			// 解码出现错误说明，说明是半包
			return frames
		}
		// 正常数据帧，添加到结果当中
		if len(decode) > 0 {
			frames = append(frames, decode)
		}
	}
}

func (d DelimiterBasedFrameDecoder) decode(buffer *bytes.Buffer) ([]byte, error) {
	if index := strings.IndexByte(buffer.String(), d.delimiter); index == -1 {
		return nil, errors.New("delimiter not found")
	}

	line, err := buffer.ReadBytes(d.delimiter)
	if err != nil {
		return nil, err
	}

	str := strings.TrimSuffix(string(line), string(d.delimiter))

	return []byte(str), nil
}
