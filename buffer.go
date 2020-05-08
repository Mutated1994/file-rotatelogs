package rotatelogs

import (
	"bufio"
	"errors"
	"io"
	"time"
)

var defaultSyncInterval = time.Millisecond * 10

type Writer struct {
	*bufio.Writer

	ticker *time.Ticker
	closed bool
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Writer: bufio.NewWriter(w),
		ticker: time.NewTicker(defaultSyncInterval),
	}
}

func (b *Writer) Write(p []byte) (nn int, err error) {
	if b.Writer.Available() < len(p) {
		_ = b.Flush()
	}
	return b.Writer.Write(p)
}

func (b *Writer) close() {
	b.closed = true
	_ = b.Writer.Flush()
}

func (b *Writer) sync() error {
	if b.closed {
		b.ticker.Stop()
		return errors.New("Writer has been closed \n")
	}
	_ = b.Writer.Flush()
	return nil
}
