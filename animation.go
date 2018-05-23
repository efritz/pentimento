package pentimento

import "time"

// AnimatedString is an implementation of the Stringer interface
// that returns the next part of the string in a cycle based on
// the time since the string was created. Used to create spinners
// and throbbers in the terminal.
type AnimatedString struct {
	parts    []string
	interval time.Duration
	start    time.Time
}

var (
	// SpinnerStrings are the strings used to make a spinner.
	SpinnerStrings = []string{`/`, `-`, `\`, `|`}

	// ThrobberStrings are the strings used to make a throbber.
	ThrobberStrings = []string{`.`, `o`, `O`, `o`}

	// Spinner is a global instance of a spinner.
	Spinner = NewDefaultSpinner()

	// Throbber is a global instance of a throbber.
	Throbber = NewDefaultThrobber()
)

// NewDefaultSpinner creates a spinner that animates four times a second.
func NewDefaultSpinner() *AnimatedString {
	return NewSpinner(time.Millisecond * 250)
}

// NewSpinner creates a spinner that animates at the given interval.
func NewSpinner(interval time.Duration) *AnimatedString {
	return NewAnimatedString(SpinnerStrings, interval)
}

// NewDefaultThrobber creates a throbber that animates four times a second.
func NewDefaultThrobber() *AnimatedString {
	return NewThrobber(time.Millisecond * 250)
}

// NewThrobber creates a throbber that animates at the given interval.
func NewThrobber(interval time.Duration) *AnimatedString {
	return NewAnimatedString(ThrobberStrings, interval)
}

// NewAnimatedString makes an animated string with the given string list
// (which animates left to right) and animation interval.
func NewAnimatedString(parts []string, interval time.Duration) *AnimatedString {
	return &AnimatedString{
		parts:    parts,
		interval: interval,
		start:    time.Now(),
	}
}

// String returns the current string in the animation cycle.
func (ac *AnimatedString) String() string {
	return ac.parts[int(time.Now().Sub(ac.start)/ac.interval)%len(ac.parts)]
}
