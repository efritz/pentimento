package main

import (
	"fmt"
	"time"

	"github.com/efritz/pentimento"
)

func main() {
	fmt.Printf("Before\n")

	pentimento.PrintProgress(func(p *pentimento.Printer) {
		for i := 1; i <= 25; i++ {
			content := pentimento.NewContent()
			content.AddLine("(%2d) Spinner: %s", i, pentimento.Spinner)
			content.AddLine("(%2d) SpinnerX: %s", i, pentimento.SpinnerX)
			content.AddLine("(%2d) SpinnerV: %s", i, pentimento.SpinnerV)
			content.AddLine("(%2d) Throbber: %s", i, pentimento.Throbber)
			content.AddLine("(%2d) Balloon: %s", i, pentimento.Balloon)
			content.AddLine("(%2d) Platform: %s", i, pentimento.Platform)
			content.AddLine("(%2d) Dot: %s", i, pentimento.Dots)
			content.AddLine("(%2d) ScrollingDot: %s", i, pentimento.ScrollingDots)
			content.AddLine("(%2d) Point: %s", i, pentimento.StarDots)
			content.AddLine("(%2d) Flip: %s", i, pentimento.Flip)
			content.AddLine("(%2d) Toggle: %s", i, pentimento.Toggle)
			content.AddLine("(%2d) DQPB: %s", i, pentimento.DQPB)
			content.AddLine("(%2d) BoundingBar: %s", i, pentimento.BoundingBar)

			p.WriteContent(content)
			<-time.After(time.Second / 2)
		}

		p.Reset()
		p.WriteString("Done")
	})

	fmt.Printf("After\n")
}
