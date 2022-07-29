package readers

import (
	"compress/gzip"
	"io"
)

// gzipReader wraps a *gzip.Reader so it closes the underlying stream
// which the gzip library doesn't.
type gzipReader struct {
	*gzip.Reader
	in io.ReadCloser
}

func NewGzipReader(in io.ReadCloser) (io.ReadCloser, error) {
	zr, err := gzip.NewReader(in)
	if err != nil {
		return nil, err
	}
	return &gzipReader{
		Reader: zr,
		in:     in,
	}, nil
}

// Close the underlying stream and the gzip reader
func (gz *gzipReader) Close() error {
	zrErr := gz.Reader.Close()
	inErr := gz.in.Close()
	if inErr != nil {
		return inErr
	}
	if zrErr != nil {
		return zrErr
	}
	return nil
}
