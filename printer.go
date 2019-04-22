package pentimento

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type (
	// Printer adds functionality over an io.Writer which uses ANSI
	// codes to delete and overwrite previously written content.
	Printer struct {
		writer   io.Writer
		lines    int
		previous *Content
		ticker   *time.Ticker
		mutex    sync.Mutex
	}

	printerConfig struct {
		writer   io.Writer
		interval time.Duration
	}

	// ConfigFunc is a function used to initialize a printer.
	ConfigFunc func(*printerConfig)
)

var (
	sequenceMoveUp    = "\u001b[1A"
	sequenceClearLine = "\u001b[1000D\u001b[0K"
)

// PrintProgress creates a printer with the given configs and
// calls the function with a reference to the printer. The printer
// is configured to refresh during the function call then stop once
// it returns.
func PrintProgress(f func(*Printer), configs ...ConfigFunc) {
	config := getConfig(configs...)
	p := &Printer{writer: config.writer}
	p.Refresh(config.interval)
	defer p.Stop()
	f(p)
}

// NewStdoutPrinter creates a printer that writes to stdout.
func NewStdoutPrinter() *Printer {
	return NewPrinter(WithWriter(os.Stdout))
}

// NewStderrPrinter creates a printer that writes to stderr.
func NewStderrPrinter() *Printer {
	return NewPrinter(WithWriter(os.Stderr))
}

// WithWriter sets the writer backing a Printer.
func WithWriter(writer io.Writer) ConfigFunc {
	return func(c *printerConfig) { c.writer = writer }
}

// WithInterval sets the default refresh interval for a Printer.
func WithInterval(interval time.Duration) ConfigFunc {
	return func(c *printerConfig) { c.interval = interval }
}

// NewPrinter creates a new Printer with the given configs.
func NewPrinter(configs ...ConfigFunc) *Printer {
	return &Printer{writer: getConfig(configs...).writer}
}

// WriteString calls WriteContent described by given format string and arg list.
func (p *Printer) WriteString(format string, args ...interface{}) error {
	content := NewContent()
	content.AddLine(format, args...)
	return p.WriteContent(content)
}

// WriteContent clears the content written by the last call to WriteContent,
// then writes the new content to the backing writer.
func (p *Printer) WriteContent(content *Content) error {
	p.previous = content
	s := content.String()
	r := p.getResetSequence()
	p.lines = strings.Count(s, "\n")
	_, err := io.Copy(p.writer, bytes.NewReader([]byte(r+s)))
	return err
}

// Reset clears the content written by the last call to WriteContent.
func (p *Printer) Reset() error {
	s := p.getResetSequence()
	p.lines = 0
	_, err := io.Copy(p.writer, bytes.NewReader([]byte(s)))
	return err
}

// Refresh starts a goroutine that re-writes the last chunk of content
// to the screen. This will have the effect of animating strings in the
// content.
func (p *Printer) Refresh(interval time.Duration) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.stop()
	ticker := time.NewTicker(interval)
	p.ticker = ticker

	go func() {
		for range ticker.C {
			p.mutex.Lock()
			p.WriteContent(p.previous)
			p.mutex.Unlock()
		}
	}()
}

// Stop will cancel an active refresh ticker.
func (p *Printer) Stop() {
	p.mutex.Lock()
	p.stop()
	p.mutex.Unlock()
}

func (p *Printer) stop() {
	if p.ticker == nil {
		return
	}

	p.ticker.Stop()
	p.ticker = nil
}

func (p *Printer) getResetSequence() string {
	return strings.Repeat(sequenceClearLine+sequenceMoveUp, p.lines) + sequenceClearLine
}

//
// Helpers

func getConfig(configs ...ConfigFunc) *printerConfig {
	config := &printerConfig{
		writer:   os.Stdout,
		interval: time.Millisecond * 250,
	}

	for _, f := range configs {
		f(config)
	}

	return config
}
