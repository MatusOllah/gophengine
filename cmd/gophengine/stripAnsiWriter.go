package main

import (
	"io"
	"regexp"
)

type StripANSIWriter struct {
	re *regexp.Regexp
	w  io.Writer
}

func NewStripANSIWriter(w io.Writer) *StripANSIWriter {
	return &StripANSIWriter{regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"), w}
}

func (saw *StripANSIWriter) Write(p []byte) (n int, err error) {
	return saw.w.Write(saw.re.ReplaceAll(p, nil))
}
