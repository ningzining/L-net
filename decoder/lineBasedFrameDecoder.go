package decoder

import (
	"bytes"
	"errors"
	"strings"
)

const (
	lineBreaks = '\n'
)

type LineBasedFrameDecoder struct {
}

func NewLineBasedFrameDecoder() Decoder {
	return &LineBasedFrameDecoder{}
}

func (d LineBasedFrameDecoder) Decode(buffer *bytes.Buffer) ([]byte, error) {
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
