package qr

import (
	"os"

	"github.com/mdp/qrterminal/v3"
)

// RenderString as a QR code
func RenderString(s string) {
	qrConfig := qrterminal.Config{
		HalfBlocks:     true,
		Level:          qrterminal.L,
		Writer:         os.Stdout,
		BlackWhiteChar: "\u001b[37m\u001b[40m\u2584\u001b[0m",
		BlackChar:      "\u001b[30m\u001b[40m\u2588\u001b[0m",
		WhiteBlackChar: "\u001b[30m\u001b[47m\u2585\u001b[0m",
		WhiteChar:      "\u001b[37m\u001b[47m\u2588\u001b[0m",
		QuietZone:      1,
	}
	qrterminal.GenerateWithConfig(s, qrConfig)
}
