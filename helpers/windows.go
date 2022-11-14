//go:build windows
// +build windows

package helpers

import "github.com/lxn/win"

func GetMaxX() int {
	var x int
	x = int(win.GetSystemMetrics(win.SM_CXSCREEN))
	return x
}

func GetMaxY() int {
	var y int
	y = int(win.GetSystemMetrics(win.SM_CYSCREEN))
	return y
}
