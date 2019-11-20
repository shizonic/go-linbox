package linbox

import (
	"fmt"

	"github.com/shizonic/go-linbox/internal/mount"
)

const mountPoint = "/mnt/linbox"

func Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		mount.UmountPseudoFs(mountPoint)
	}()

	// mount.MountPseudoFs(mountPoint)
}
