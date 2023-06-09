// The term package provides useful tooling for working with formatting and colors in a terminal.
//
// The following formatting options are available:
//
//   - Bold
//   - Faint
//   - Italic
//   - Underline
//   - Blink
//   - Negative
//   - Invisible
//   - StrikeThrough
//
// The following colors are available for both text and background:
//
//   - Black and BlackBackground
//   - Red and RedBackground
//   - Green and GreenBackground
//   - Yellow and YellowBackground
//   - Blue and BlueBackground
//   - Magenta and MagentaBackground
//   - Cyan and CyanBackground
//   - White and WhiteBackground
//
// The colors also exist in a bright variant:
//
//   - BrightBlack and BrightBlackBackground
//   - BrightRed and BrightRedBackground
//   - BrightGreen and BrightGreenBackground
//   - BrightYellow and BrightYellowBackground
//   - BrightBlue and BrightBlueBackground
//   - BrightMagenta and BrightMagentaBackground
//   - BrightCyan and BrightCyanBackground
//   - BrightWhite and BrightWhiteBackground
package term

import (
	"syscall"
	"unsafe"
)

// Size returns the dimensions of the terminal.
//
// To get notified when the terminal is resized, use the os/signal package.
//
//	resize := make(chan os.Signal, 1)
//	signal.Notify(resize, syscall.SIGWINCH)
//	for {
//		select {
//		case <-resize:
//			rows, cols := term.Size()
//			// ...
//		}
//	}
func Size() (rows, cols int) {
	var size struct {
		Rows    uint16
		Cols    uint16
		Xpixels uint16
		Ypixels uint16
	}

	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&size)),
	)

	if int(retCode) == -1 {
		panic(errno)
	}

	return int(size.Rows), int(size.Cols)
}
