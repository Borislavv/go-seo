package controller

import "io"

type WriterOrStringWriter interface {
	io.Writer
	io.StringWriter
}
