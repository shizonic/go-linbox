package mount

import (
	"fmt"
	"os"
	"syscall"

	"github.com/shizonic/go-linbox/internal/utils"
)

func mountBind(source, target, fstype string) {
	sourcePath := utils.ExpandPath(source)
	targetPath := utils.ExpandPath(target)

	err := os.MkdirAll(targetPath, 0644)
	utils.CheckError("Creating mount point directory", err)

	err = syscall.Mount(sourcePath, targetPath, fstype, syscall.MS_BIND, "")
	utils.CheckError(fmt.Sprintf("Mounting '%s' to '%s'", sourcePath, targetPath), err)
}

func unmount(target string) {
	targetPath := utils.ExpandPath(target)
	err := syscall.Unmount(targetPath, syscall.MNT_DETACH)
	utils.CheckError(fmt.Sprintf("Unmounting '%s'", targetPath), err)
}
