package decoder

import (
	"bytes"
	"errors"
	"strings"
)

const (
	lineBreaks = '\n'
)

// LineBasedFrameDecoder 行分帧解码器
type LineBasedFrameDecoder struct {
}

func NewLineBasedFrameDecoder() Decoder {
	return &LineBasedFrameDecoder{}
}

func (d LineBasedFrameDecoder) Decode(buffer *bytes.Buffer) [][]byte {
	frames := make([][]byte, 0)
	for {
		decode, err := d.decode(buffer)
		if err != nil {
			return frames
		}
		// 正常数据帧，添加到结果当中
		if len(decode) > 0 {
			frames = append(frames, decode)
		}
	}
}

func (d LineBasedFrameDecoder) decode(buffer *bytes.Buffer) ([]byte, error) {
	if index := strings.IndexByte(buffer.String(), lineBreaks); index == -1 {
		return nil, errors.New("delimiter not found")
	}

	line, err := buffer.ReadBytes(lineBreaks)
	if err != nil {
		return nil, err
	}

	str := strings.TrimSuffix(string(line), string(lineBreaks))

	return []byte(str), nil
}
