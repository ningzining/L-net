package decoder

import "bytes"

type Decoder interface {
	Decode(buffer *bytes.Buffer) ([]byte, error)
}
