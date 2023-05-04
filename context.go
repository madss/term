package term

import (
	"fmt"
	"io"
	"strings"
)

// A Context keeps track of the current state of the terminal.
type Context struct {
	parent *Context
	bypass bool

	// Styling
	bold          bool
	faint         bool
	italic        bool
	underline     bool
	blink         bool
	negative      bool
	invisible     bool
	strikeThrough bool

	// Colors
	foreground color
	background color
}

// With creates a new child context.
//
// The new context overrides the attributes of the parent context.
func (ctx *Context) With(opts ...Option) *Context {
	childCtx := Context{parent: ctx, bypass: ctx.bypass}
	childCtx.Apply(opts...)
	return &childCtx
}

// Reset resets the context to the default state.
func (ctx *Context) Reset() {
	ctx.bold = false
	ctx.faint = false
	ctx.italic = false
	ctx.underline = false
	ctx.blink = false
	ctx.negative = false
	ctx.invisible = false
	ctx.strikeThrough = false
	ctx.foreground = original
	ctx.background = original
}

// Set sets the state of the context to the provided options.
func (ctx *Context) Set(opts ...Option) {
	ctx.Reset()
	ctx.Apply(opts...)
}

// Apply updates the state of the context with the additional options.
func (ctx *Context) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(ctx)
	}
}

// Format formats a string as with fmt.Sprintf but with formatting applied.
func (ctx *Context) Format(format string, a ...interface{}) string {
	s := fmt.Sprintf(format, a...)
	return ctx.Render(s)
}

// Render returns s with with formatting applied.
func (ctx *Context) Render(s string) string {
	var b strings.Builder
	ctx.Write(&b, []byte(s))
	return b.String()
}

// Write writes data to w with formatting applied.
//
// Note: The number of bytes written excludes the escape codes, to comply with
// constumers that expect the number of bytes written to be less than or equal
// to then length of data.
func (ctx *Context) Write(w io.Writer, data []byte) (int, error) {
	if _, err := w.Write(diff(ctx.parent, ctx).Bytes()); err != nil {
		return 0, err
	}
	n, err := w.Write(data)
	if err != nil {
		return n, err
	}
	_, err = w.Write(diff(ctx, ctx.parent).Bytes())
	return n, err
}

func diff(oldCtx, newCtx *Context) *escapeSequence {
	if oldCtx == nil {
		oldCtx = &Context{}
	}
	if newCtx == nil {
		newCtx = &Context{}
	}
	seq := escapeSequence{kind: 'm'}
	seq.Adjust(oldCtx.bold, newCtx.bold, boldOn, boldOff)
	seq.Adjust(oldCtx.faint, newCtx.faint, faintOn, faintOff)
	seq.Adjust(oldCtx.italic, newCtx.italic, italicOn, italicOff)
	seq.Adjust(oldCtx.underline, newCtx.underline, underlineOn, underlineOff)
	seq.Adjust(oldCtx.blink, newCtx.blink, blinkOn, blinkOff)
	seq.Adjust(oldCtx.negative, newCtx.negative, negativeOn, negativeOff)
	seq.Adjust(oldCtx.invisible, newCtx.invisible, invisibleOn, invisibleOff)
	seq.Adjust(oldCtx.strikeThrough, newCtx.strikeThrough, strikeThroughOn, strikeThroughOff)
	if oldCtx.foreground != newCtx.foreground {
		seq.Add(fgColors[newCtx.foreground])
	}
	if oldCtx.background != newCtx.background {
		seq.Add(bgColors[newCtx.background])
	}
	return &seq
}

// An Option is a function that modifies a Context.
type Option func(*Context)

func Bold(ctx *Context) {
	ctx.bold = true
}

func NoBold(ctx *Context) {
	ctx.bold = false
}

func Faint(ctx *Context) {
	ctx.faint = true
}

func NoFaint(ctx *Context) {
	ctx.faint = false
}

func Italic(ctx *Context) {
	ctx.italic = true
}

func NoItalic(ctx *Context) {
	ctx.italic = false
}

func Underline(ctx *Context) {
	ctx.underline = true
}

func NoUnderline(ctx *Context) {
	ctx.underline = false
}

func Blink(ctx *Context) {
	ctx.blink = true
}

func NoBlink(ctx *Context) {
	ctx.blink = false
}

func Negative(ctx *Context) {
	ctx.negative = true
}

func NoNegative(ctx *Context) {
	ctx.negative = false
}

func Invisible(ctx *Context) {
	ctx.invisible = true
}

func NoInvisible(ctx *Context) {
	ctx.invisible = false
}

func StrikeThrough(ctx *Context) {
	ctx.strikeThrough = true
}

func NoStrikeThrough(ctx *Context) {
	ctx.strikeThrough = false
}

func Default(ctx *Context) {
	ctx.foreground = original
}

func Black(ctx *Context) {
	ctx.foreground = black
}

func Red(ctx *Context) {
	ctx.foreground = red
}

func Green(ctx *Context) {
	ctx.foreground = green
}

func Yellow(ctx *Context) {
	ctx.foreground = yellow
}

func Blue(ctx *Context) {
	ctx.foreground = blue
}

func Magenta(ctx *Context) {
	ctx.foreground = magenta
}

func Cyan(ctx *Context) {
	ctx.foreground = cyan
}

func White(ctx *Context) {
	ctx.foreground = white
}

func BrightBlack(ctx *Context) {
	ctx.foreground = brightBlack
}

func BrightRed(ctx *Context) {
	ctx.foreground = brightRed
}

func BrightGreen(ctx *Context) {
	ctx.foreground = brightGreen
}

func BrightYellow(ctx *Context) {
	ctx.foreground = brightYellow
}

func BrightBlue(ctx *Context) {
	ctx.foreground = brightBlue
}

func BrightMagenta(ctx *Context) {
	ctx.foreground = brightMagenta
}

func BrightCyan(ctx *Context) {
	ctx.foreground = brightCyan
}

func BrightWhite(ctx *Context) {
	ctx.foreground = brightWhite
}

func DefaultBackground(ctx *Context) {
	ctx.background = original
}

func BlackBackground(ctx *Context) {
	ctx.background = black
}

func RedBackground(ctx *Context) {
	ctx.background = red
}

func GreenBackground(ctx *Context) {
	ctx.background = green
}

func YellowBackground(ctx *Context) {
	ctx.background = yellow
}

func BlueBackground(ctx *Context) {
	ctx.background = blue
}

func MagentaBackground(ctx *Context) {
	ctx.background = magenta
}

func CyanBackground(ctx *Context) {
	ctx.background = cyan
}

func WhiteBackground(ctx *Context) {
	ctx.background = white
}

func BrightBlackBackground(ctx *Context) {
	ctx.background = brightBlack
}

func BrightRedBackground(ctx *Context) {
	ctx.background = brightRed
}

func BrightGreenBackground(ctx *Context) {
	ctx.background = brightGreen
}

func BrightYellowBackground(ctx *Context) {
	ctx.background = brightYellow
}

func BrightBlueBackground(ctx *Context) {
	ctx.background = brightBlue
}

func BrightMagentaBackground(ctx *Context) {
	ctx.background = brightMagenta
}

func BrightCyanBackground(ctx *Context) {
	ctx.background = brightCyan
}

func BrightWhiteBackground(ctx *Context) {
	ctx.background = brightWhite
}

type color int

const (
	original color = iota
	black
	red
	green
	yellow
	blue
	magenta
	cyan
	white
	brightBlack
	brightRed
	brightGreen
	brightYellow
	brightBlue
	brightMagenta
	brightCyan
	brightWhite
)

var fgColors = map[color]attribute{
	original:      defaultFg,
	black:         blackFg,
	red:           redFg,
	green:         greenFg,
	yellow:        yellowFg,
	blue:          blueFg,
	magenta:       magentaFg,
	cyan:          cyanFg,
	white:         whiteFg,
	brightBlack:   brightBlackFg,
	brightRed:     brightRedFg,
	brightGreen:   brightGreenFg,
	brightYellow:  brightYellowFg,
	brightBlue:    brightBlueFg,
	brightMagenta: brightMagentaFg,
	brightCyan:    brightCyanFg,
	brightWhite:   brightWhiteFg,
}

var bgColors = map[color]attribute{
	original:      defaultBg,
	black:         blackBg,
	red:           redBg,
	green:         greenBg,
	yellow:        yellowBg,
	blue:          blueBg,
	magenta:       magentaBg,
	cyan:          cyanBg,
	white:         whiteBg,
	brightBlack:   brightBlackBg,
	brightRed:     brightRedBg,
	brightGreen:   brightGreenBg,
	brightYellow:  brightYellowBg,
	brightBlue:    brightBlueBg,
	brightMagenta: brightMagentaBg,
	brightCyan:    brightCyanBg,
	brightWhite:   brightWhiteBg,
}
