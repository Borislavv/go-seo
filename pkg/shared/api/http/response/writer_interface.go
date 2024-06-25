package response

import "io"

type WriterOrStringWriter interface {
	io.Writer
	io.StringWriter
}
