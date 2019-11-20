package mount

import (
	"path"
)

var paths []string = []string{
	"/proc",
	"/sys",
	"/dev",
	"/run",
}

func MountPseudoFs(p string) {
	for _, pa := range paths {
		mountBind(pa, path.Join(p, pa), "none")
	}
}

func UmountPseudoFs(p string) {
	for _, pa := range paths {
		unmount(path.Join(p, pa))
	}
}
