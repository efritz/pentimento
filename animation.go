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
	DefaultInterval = time.Millisecond * 250
	SpinnerStrings  = []string{`/`, `-`, `\`, `|`}
	SpinnerXStrings = []string{`+`, `x`}
	SpinnerVStrings = []string{`v`, `<`, `^`, `>`}
	ThrobberStrings = []string{`.`, `o`, `O`, `o`}
	BalloonStrings  = []string{`.`, `o`, `O`, `@`, `*`, ` `}
	PlatformStrings = []string{`_`, `-`}

	Spinner  = NewAnimatedString(SpinnerStrings, DefaultInterval)
	SpinnerX = NewAnimatedString(SpinnerXStrings, DefaultInterval)
	SpinnerV = NewAnimatedString(SpinnerVStrings, DefaultInterval)
	Throbber = NewAnimatedString(ThrobberStrings, DefaultInterval)
	Balloon  = NewAnimatedString(BalloonStrings, DefaultInterval)
	Platform = NewAnimatedString(PlatformStrings, DefaultInterval)
)

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
