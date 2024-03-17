package decoder

import (
	"bytes"
	"errors"
	"strings"
)

type DelimiterBasedFrameDecoder struct {
	delimiter byte
}

func NewDelimiterBasedFrameDecoder(delimiter byte) Decoder {
	return &DelimiterBasedFrameDecoder{
		delimiter: delimiter,
	}
}

func (d DelimiterBasedFrameDecoder) Decode(buffer *bytes.Buffer) ([]byte, error) {
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
