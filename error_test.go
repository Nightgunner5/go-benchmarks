package test

import (
	"testing"
)

func oldError(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

type newError string

func (e newError) Error() string {
	return string(e)
}

func BenchmarkErrorOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var e error = oldError("abcdefg")
		_ = e.Error()
	}
}

func BenchmarkErrorNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var e error = newError("abcdefg")
		_ = e.Error()
	}
}
