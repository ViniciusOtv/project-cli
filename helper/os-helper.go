package helper

import "runtime"

func IsWindowsSystem() bool {
	return runtime.GOOS == "windows"
}
