package markdown

import (
	"bytes"
	"fmt"

	pkgio "github.com/cardil/ip-tax-contributions/pkg/io"
)

type Builder struct {
	buf bytes.Buffer
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WriteTo(out pkgio.Output) {
	out.Print(b.buf.String())
}

func (b *Builder) Title(text ...any) {
	_, _ = fmt.Fprint(&b.buf, "# ")
	_, _ = fmt.Fprint(&b.buf, text...)
	_, _ = fmt.Fprintln(&b.buf)
	_, _ = fmt.Fprintln(&b.buf)
}

func (b *Builder) Header(text ...any) {
	_, _ = fmt.Fprint(&b.buf, "## ")
	_, _ = fmt.Fprint(&b.buf, text...)
	_, _ = fmt.Fprintln(&b.buf)
	_, _ = fmt.Fprintln(&b.buf)
}

func (b *Builder) Text(text ...any) {
	_, _ = fmt.Fprintln(&b.buf, text...)
	_, _ = fmt.Fprintln(&b.buf)
}

func (b *Builder) Link(text, address string) {
	_, _ = fmt.Fprintln(&b.buf, "- [", text, "](", address, ")")
}
