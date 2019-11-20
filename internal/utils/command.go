package utils

import (
	"fmt"
	"os/exec"
)

// ShellExec Executes a command string using the systems shell
func ShellExec(c string) string {
	cmd := exec.Command("sh", "-c", c)
	stdoutStderr, err := cmd.CombinedOutput()
	CheckError(fmt.Sprintf("Executing '%s'", c), err)
	return string(stdoutStderr)
}
