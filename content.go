package pentimento

import (
	"fmt"
	"strings"
)

type (
	// Content is a wrapper around formatted strings which can
	// be applied to a Printer atomically.
	Content struct {
		parts []*contentPart
	}

	contentPart struct {
		format string
		args   []interface{}
	}
)

// NewContent creates an empty Content.
func NewContent() *Content {
	return &Content{}
}

// AddLine adds a format string and arg list to the end of the
// content list.
func (c *Content) AddLine(format string, args ...interface{}) {
	c.parts = append(c.parts, &contentPart{format, args})
}

// String serializes the lines of content. If some part contains
// an animated string, it will be re-evaluated on each invocation
// of this method.
func (c *Content) String() string {
	lines := []string{}
	for _, part := range c.parts {
		lines = append(lines, fmt.Sprintf(part.format+"\n", part.args...))
	}

	return strings.Join(lines, "")
}
