package testutil

import (
	"bytes"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

type BufferReader interface {
	io.Reader
	Reset(string)
}

type BufferWriter interface {
	io.Writer
	Reset()
	Bytes() []byte
	String() string
}

func ApplyMockIO(c *cobra.Command) (BufferReader, BufferWriter) {
	mockIn := strings.NewReader("")
	mockOut := bytes.NewBufferString("")

	c.SetIn(mockIn)
	c.SetOut(mockOut)
	c.SetErr(mockOut)

	return mockIn, mockOut
}
