package optiongen

import (
	"bytes"
	"fmt"
)

const (
	optionDeclarationSuffix = "OptionDeclareWithDefault"
	OptionGen               = "optiongen"
)

type BufWrite struct {
	buf *bytes.Buffer
}

func (b BufWrite) wf(format string, vals ...interface{}) {
	_, _ = fmt.Fprintf(b.buf, format, vals...)
}

func (b BufWrite) wln(vals ...interface{}) {
	_, _ = fmt.Fprintln(b.buf, vals...)
}
