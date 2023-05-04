package term

import (
	"strconv"
	"strings"
)

type attribute int

const (
	reset attribute = 0

	// Formatting
	boldOn          attribute = 1
	faintOn         attribute = 2
	italicOn        attribute = 3
	underlineOn     attribute = 4
	blinkOn         attribute = 5
	negativeOn      attribute = 7
	invisibleOn     attribute = 8
	strikeThroughOn attribute = 9

	boldOff          attribute = 21
	faintOff         attribute = 22
	italicOff        attribute = 23
	underlineOff     attribute = 24
	blinkOff         attribute = 25
	negativeOff      attribute = 27
	invisibleOff     attribute = 28
	strikeThroughOff attribute = 29

	// Foreground colors
	defaultFg       attribute = 39
	blackFg         attribute = 30
	redFg           attribute = 31
	greenFg         attribute = 32
	yellowFg        attribute = 33
	blueFg          attribute = 34
	magentaFg       attribute = 35
	cyanFg          attribute = 36
	whiteFg         attribute = 37
	brightBlackFg   attribute = 90
	brightRedFg     attribute = 91
	brightGreenFg   attribute = 92
	brightYellowFg  attribute = 93
	brightBlueFg    attribute = 94
	brightMagentaFg attribute = 95
	brightCyanFg    attribute = 96
	brightWhiteFg   attribute = 97

	// Background colors
	defaultBg       attribute = 49
	blackBg         attribute = 40
	redBg           attribute = 41
	greenBg         attribute = 42
	yellowBg        attribute = 43
	blueBg          attribute = 44
	magentaBg       attribute = 45
	cyanBg          attribute = 46
	whiteBg         attribute = 47
	brightBlackBg   attribute = 100
	brightRedBg     attribute = 101
	brightGreenBg   attribute = 102
	brightYellowBg  attribute = 103
	brightBlueBg    attribute = 104
	brightMagentaBg attribute = 105
	brightCyanBg    attribute = 106
	brightWhiteBg   attribute = 107
)

type escapeSequence []attribute

func (attrs *escapeSequence) Add(attr attribute) {
	*attrs = append(*attrs, attr)
}

func (attrs *escapeSequence) Adjust(was, is bool, on, off attribute) {
	if !was && is {
		attrs.Add(on)
	} else if was && !is {
		attrs.Add(off)
	}
}

func (attrs escapeSequence) Bytes() []byte {
	if len(attrs) == 0 {
		return nil
	}

	var b strings.Builder
	b.Grow(16) // should be enough in must situations

	b.WriteString("\033[")
	for i := range attrs {
		if i > 0 {
			b.WriteByte(';')
		}
		b.WriteString(strconv.Itoa(int(attrs[i])))
	}
	b.WriteString("m")
	return []byte(b.String())
}
