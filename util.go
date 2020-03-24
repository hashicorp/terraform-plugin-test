package tftest

import (
	"os"
)

func symlinkFile(src string, dest string) (err error) {
	err = os.Symlink(src, dest)
	if err == nil {
		srcInfo, err := os.Stat(src)
		if err != nil {
			err = os.Chmod(dest, srcInfo.Mode())
		}
	}

	return
}
