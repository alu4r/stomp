package frame

import (
	"bufio"
	"io"
)

// slices used to write frames
var (
	crlfSlice    = []byte{CrByte, LfByte} // CR-LF
	newlineSlice = []byte{LfByte}         // newline (LF)
	nullSlice    = []byte{NulByte}        // null character
)

// Writer STOMP frames to an underlying io.Writer.
type Writer struct {
	writer *bufio.Writer
}

// NewWriter Creates a new Writer object, which writes to an underlying io.Writer.
func NewWriter(writer io.Writer) *Writer {
	return NewWriterSize(writer, 4096)
}

func NewWriterSize(writer io.Writer, bufferSize int) *Writer {
	return &Writer{writer: bufio.NewWriterSize(writer, bufferSize)}
}

// Write the contents of a frame to the underlying io.Writer.
func (w *Writer) Write(f *Frame) error {
	var err error

	if f == nil {
		_, err = w.writer.Write(newlineSlice)
		if err != nil {
			return err
		}
	} else {
		_, err = w.writer.Write([]byte(f.Command))
		if err != nil {
			return err
		}

		_, err = w.writer.Write(newlineSlice)
		if err != nil {
			return err
		}

		if f.Header != nil && f.Header.slice != nil {
			headerStr := f.Header.toString()
			_, err = replacerForEncodeValue.WriteString(w.writer, headerStr)
			if err != nil {
				return err
			}
		}

		_, err = w.writer.Write(newlineSlice)
		if err != nil {
			return err
		}

		if len(f.Body) > 0 {
			_, err = w.writer.Write(f.Body)
			if err != nil {
				return err
			}
		}

		_, err = w.writer.Write(nullSlice)
		if err != nil {
			return err
		}
	}

	err = w.writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
