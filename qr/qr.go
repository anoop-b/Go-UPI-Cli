package qr

import (
	"os"

	"github.com/mdp/qrterminal/v3"
)

// RenderString as a QR code
func RenderString(s string) {
	qrConfig := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 2,
	}
	qrterminal.GenerateWithConfig(s, qrConfig)
}
