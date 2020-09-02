package docs

import (
	"fmt"
	"os/exec"
	"runtime"
)

const (
	docsURL = "https://github.com/anoop-b/Go-UPI-Cli"
)

//Launch url for project wiki/documentation
func Launch() error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", docsURL).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", docsURL).Start()
	case "darwin":
		err = exec.Command("open", docsURL).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		return fmt.Errorf("open docs: %w", err)
	}

	return err
}
